package main

import (
        "fmt"
        "github.com/akamensky/argparse"
        "os"
        "os/exec"
)

func main() {
        // Create new parser object
        parser := argparse.NewParser("print", "Prints provided string to stdout")
        // Create string flag
        s := parser.StringList("s", "secret", &argparse.Options{Required: true, Help: "Secrets."})
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

        input_secrets := *s
        input_controller := *c
        input_namespace := *n
        input_scope := *scope

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

}