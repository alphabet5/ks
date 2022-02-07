# ks
 Utility to make kubeseal --raw a bit easier.

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
curl https://github.com/alphabet5/ks/releases/download/0.0.1/ks-darwin-amd64
mv ./ks /usr/local/bin
```

## Usage

```bash
% ks 
[-s|--secret] is required
usage: print [-h|--help] -s|--secret "<value>" [-s|--secret "<value>" ...]
             [-c|--controller "<value>"] [-n|--namespace "<value>"] [--scope
             "<value>"]

             Prints provided string to stdout

Arguments:

  -h  --help        Print help information
  -s  --secret      Secrets.
  -c  --controller  Sealed secrets controller name.. Default: sealed-secrets
  -n  --namespace   Sealed secrets controller namespace.. Default:
                    sealed-secrets
      --scope       Sealed secret scope.. Default: cluster-wide

```

```bash
ks -s secret1 -s secret2
```