package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func dotNotationReplace(yml map[string]interface{}, path []string, value string) map[string]interface{} {
	if len(path) > 1 {
		if v, ok := yml[path[0]]; ok {
			yml[path[0]] = dotNotationReplace(v.(map[string]interface{}), path[1:], value)
		} else {
			log.Fatal("Error finding path in output yaml. ", path)
		}
	} else {
		yml[path[0]] = fmt.Sprintf("%v", value)
	}
	return yml
}

func main() {
	// set logging to stdout
	log.SetOutput(os.Stdout)
	// Create new parser object
	parser := argparse.NewParser("ks", `
        Converts secrets into sealedsecrets using kubeseal.
        Secrets can be entered as strings with -s, or as an input .yaml file with -i
        When specifying files -o will be merged with the encrypted values from -i. 
        For example: 
        $.values.child1.yoursecret: "value to encrypt" 
        will be merged into
        values:
          child1:
            yoursecret: AgDXO9g9vDAhJbNxlIBbYDTkGE3gqEilK6DMxy4aJD12FGAclg2Sxa4q4VA90VcCPdDzojezD8vsh7X/Ef/1FhVATgd4+62jb9EVpqj5fFpdagZMm4Fx4FNSamrtbKEnzAhH//YaB+2Ak3fE0HjgtIZpRUomCgOsuRPIXfJlQ4n0l5wlc1wohZvoSkhaHm2SWgGjU9JY6GxYHK811gfw7xe+DSm1qoJvAs8XvXZw3jTF839sVGlSA6I4VguWH1Bx7Ev7H8+mKnnrZcSEicKedzYOXWdG58JaIjCPVNzsDhvmBlXQHlOBvSF07zQp3hvVS/aqx4MHiOrKC9IGRS42A9WtNB+pQXRyUrKGY8xCnjVFgaIDGiLWWF+gMNJPTb//fHOldiS+nQH6s/EQX3dv1rjLC+F8qY50emts/VPVUwyU7i0qIo99sdArRabaITbJ6fBdPA8QL4gvbqtH4Fhrfb7EtMm0MeyaaqL58qq5N/r9LVbYtKbbdNHmR4MLIN7FjIoEpmWxDYb4MFs+M+smuVmmWZmGR02XEP/W2d0ag7TiNFo+N6tIuMDPzYsNhTBWGen5b1CjkAVwcP9lgunUGHqloXpTeVdPODa8eY/MLOnlB8cosVWoWyd5G5Sp5z1ZeCLvYuOmiHSwoRq09qBRShd1lhv2Xhw5vlGquw3ODFZnL5Qpwg7IUUu5GsvMc7l1UBnHcXct
            `)
	// Create string flag
	s := parser.StringList("s", "secret", &argparse.Options{Required: false, Default: nil, Help: "Secrets."})
	i := parser.String("i", "input", &argparse.Options{Required: false, Default: "", Help: "Input secrets file."})
	o := parser.String("o", "output", &argparse.Options{Required: false, Default: "", Help: "Output file to put the secrest."})
	c := parser.String("c", "controller", &argparse.Options{Required: false, Default: "sealed-secrets", Help: "Sealed secrets controller name."})
	n := parser.String("n", "namespace", &argparse.Options{Required: false, Default: "sealed-secrets", Help: "Sealed secrets controller namespace."})
	scope := parser.String("", "scope", &argparse.Options{Required: false, Default: "cluster-wide", Help: "Sealed secret scope."})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	input_controller := *c
	input_namespace := *n
	input_scope := *scope
	if len(*s) > 0 && (*o == "") {
		input_secrets := *s
		for _, secret := range input_secrets {
			fmt.Println(secret)
			cmd := exec.Command("bash", "-c",
				"echo -n \""+secret+
					"\" | kubeseal --controller-namespace "+input_namespace+
					" --raw --scope "+input_scope+
					" --from-file=/dev/stdin --controller-name "+input_controller,
			)
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// Print the output
			fmt.Println(string(stdout))
		}
	} else if (*i != "" || len(*s) > 0) && (*o != "") {
		inputValues := make(map[string]interface{})
		var inputYamlFile []byte
		if len(*s) > 0 {
			inputYamlFile = []byte((*s)[0])
		} else {
			inputYamlFile, err = ioutil.ReadFile(*i)
		}

		if err != nil {
			log.Printf("inputYamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(inputYamlFile, &inputValues)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
                
		outputValues := make(map[string]interface{})
		outputYamlFile, err := ioutil.ReadFile(*o)
		if err != nil {
			log.Printf("outputYamlFile.Get err   #%v ", err)
		}
                cm := yaml.CommentMap{}
		err = yaml.UnmarshalWithOptions(outputYamlFile, &outputValues, yaml.CommentToMap(cm))
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		for path, value := range inputValues {
			cmd := exec.Command("bash", "-c",
				"echo -n \""+fmt.Sprintf("%v", value)+
					"\" | kubeseal --controller-namespace "+input_namespace+
					" --raw --scope "+input_scope+
					" --from-file=/dev/stdin --controller-name "+input_controller,
			)
			stdout, err := cmd.Output()
			if err != nil {
				log.Fatal("Error running kubectl command. ", err)
			}
			outputValues = dotNotationReplace(outputValues, strings.Split(path, "."), string(stdout))
		}
                output, err := yaml.MarshalWithOptions(outputValues, yaml.WithComment(cm))
		//output, err := yaml.Marshal(outputValues)
		if err != nil {
			log.Fatalf("Marshal: %v", err)
		}
		fmt.Printf("%v", string(output))

	} else {
		fmt.Print("Flags were incorrect. -s or -i are required. -o is required if input is yaml.")
	}

}
