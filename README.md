# ks
 Utility to make kubeseal --raw a bit easier.

 This allows saving your unencrypted secrets in collapsed / dot-notation yaml. (In a secrets/password manager...)

## Building

```bash
GOOS=windows GOARCH=amd64 go build -o ks-windows-amd64.exe ks.go
GOOS=windows GOARCH=386 go build -o ks-windows-x86.exe ks.go
GOOS=darwin GOARCH=amd64 go build -o ks-darwin-amd64 ks.go
GOOS=darwin GOARCH=arm64 go build -o ks-darwin-arm64 ks.go
GOOS=linux GOARCH=amd64 go build -o ks-linux-amd64 ks.go
GOOS=linux GOARCH=386 go build -o ks-linux-x86 ks.go
GOOS=linux GOARCH=arm go build -o ks-linux-arm ks.go
GOOS=linux GOARCH=arm64 go build -o ks-linux-arm64 ks.go
GOOS=linux GOARCH=riscv64 go build -o ks-linux-riscv64 ks.go
```

## Installation (MacOS, amd64)

```bash
curl https://github.com/alphabet5/ks/releases/download/0.0.3/ks-darwin-amd64
chmod 
mv ./ks-darwin-amd64 /usr/local/bin/ks
```

## Installation (Linux, <arch>)

```bash
curl https://github.com/alphabet5/ks/releases/download/0.0.3/ks-linux-amd64
chmod +x ./ks-linux-amd64
mv ./ks-linux-amd64 /usr/local/bin/ks
```

## Installation from Source

```bash
cd ~/
git clone https://github.com/alphabet5/ks.git
cd ks
go build
sudo mv ks /usr/local/bin/ks

cd ~/
git clone https://github.com/bitnami/sealed-secrets.git
cd sealed-secrets/cmd/kubeseal
go build
sudo mv kubeseal /usr/local/bin/kubeseal
```

## Usage

```bash
% ks -h
usage: ks [-h|--help] [-s|--secret "<value>" [-s|--secret "<value>" ...]]
          [-i|--input "<value>"] [-o|--output "<value>"] [-c|--controller
          "<value>"] [-n|--namespace "<value>"] [--scope "<value>"] [--cert
          "<value>"]


        Converts secrets into sealedsecrets using kubeseal.
        Secrets can be entered as strings with -s, or as an input .yaml file
          with -i
        When specifying files -o will be merged with the encrypted values from
          -i.
        For example:
        $.values.child1.yoursecret: "value to encrypt"
        will be merged into
        values:
          child1:
            yoursecret:
          AgDXO9g9vDAhJbNxlIBbYDTkGE3gqEilK6DMxy4aJD12FGAclg2Sxa4q4VA90VcCPdDzojezD8vsh7X/Ef/1FhVATgd4+62jb9EVpqj5fFpdagZMm4Fx4FNSamrtbKEnzAhH//YaB+2Ak3fE0HjgtIZpRUomCgOsuRPIXfJlQ4n0l5wlc1wohZvoSkhaHm2SWgGjU9JY6GxYHK811gfw7xe+DSm1qoJvAs8XvXZw3jTF839sVGlSA6I4VguWH1Bx7Ev7H8+mKnnrZcSEicKedzYOXWdG58JaIjCPVNzsDhvmBlXQHlOBvSF07zQp3hvVS/aqx4MHiOrKC9IGRS42A9WtNB+pQXRyUrKGY8xCnjVFgaIDGiLWWF+gMNJPTb//fHOldiS+nQH6s/EQX3dv1rjLC+F8qY50emts/VPVUwyU7i0qIo99sdArRabaITbJ6fBdPA8QL4gvbqtH4Fhrfb7EtMm0MeyaaqL58qq5N/r9LVbYtKbbdNHmR4MLIN7FjIoEpmWxDYb4MFs+M+smuVmmWZmGR02XEP/W2d0ag7TiNFo+N6tIuMDPzYsNhTBWGen5b1CjkAVwcP9lgunUGHqloXpTeVdPODa8eY/MLOnlB8cosVWoWyd5G5Sp5z1ZeCLvYuOmiHSwoRq09qBRShd1lhv2Xhw5vlGquw3ODFZnL5Qpwg7IUUu5GsvMc7l1UBnHcXct


Arguments:

  -h  --help        Print help information
  -s  --secret      Secrets.
  -i  --input       Input secrets file. Default:
  -o  --output      Output file to put the secrets. Default:
  -c  --controller  Sealed secrets controller name. Default: sealed-secrets
  -n  --namespace   Sealed secrets controller namespace. Default:
                    sealed-secrets
      --scope       Sealed secret scope. Default: cluster-wide
      --cert        Certificate file. Default:
```

```bash
ks -s secret1 -s secret2
```
```
ks -s "parent.child1.grandchild1: please encrypt me
parent.child2.grandchild2: please also encrypt me" -o output.yaml
```

```
ks -i input.yaml -o output.yaml
```

```
ks -i input.yaml -o output.yaml --cert=downloaded-certificate.pem
```

## Changelog

### 0.0.4
- Add support for providing a certificate file.
- Bumped go to 

### 0.0.3

- Changed the sending to kubeseal to be quoted with ' instead of "