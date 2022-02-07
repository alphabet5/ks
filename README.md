# ks
 Utility to make kubeseal --raw a bit easier.

## Installation



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