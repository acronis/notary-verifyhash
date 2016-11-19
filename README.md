# Acronis Notary verify hash CLI utility
## Overview
CLI for Acronis Notary to verify hash in Merkle Patricia Tree root/proof.

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)

## Installation

1. Clone the repository `mkdir verifyhash && git clone git@github.com:acronis/notary-verifyhash.git && cd verify`
2. Build and install dependencies `go get -d ./... && go install`

## How to use

- Print the `verifyhash` help

```
$ verifyhash --help

USAGE:
   verifyhash [global options]

   Required flags: -c|-o|-fo, -p|-fp, -r

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cert value, -c value       ID of certificate
   --object value, -o value     "Object" from certificate
   --fobject value, --fo value  Path to the file with the "object" from certificate
   --root value, -r value       Merkle root
   --proof value, -p value      Merkle proof
   --fproof value, --fp value   Path to the file with merkle proof
   --help, -h                   show help
   --version, -v                print the version
```


#### Example usage:

Standart input:

  ```
  $ verifyhash -c 7f420817b35d23a9e6f52fa715bffee975c009340419d82d596b700ef03222e9 -r b829f2d554623b22b6ee431dff294e71f24eab79ddc1f83435a2cc606ca53623 -p [\"f843a1207f420817b35d23a9e6f52fa715bffee975c009340419d82d596b700ef03222e9a03131656533626565393065666433383237323465376135643631666531376430\"]
  ```

  The object instead of a certificate ID:

  ```
  $ verifyhash -o {\"eTag\": \"11ee3bee90efd382724e7a5d61fe17d0\", \"key\":\"init.sql\",\"sequencer\": \"24EA5D176F5F21959E\",\"size\": 5391} -r b829f2d554623b22b6ee431dff294e71f24eab79ddc1f83435a2cc606ca53623 -p [\"f843a1207f420817b35d23a9e6f52fa715bffee975c009340419d82d596b700ef03222e9a03131656533626565393065666433383237323465376135643631666531376430\"]
  ```  

  Object or Proof from file:

  ```
  $ verifyhash -fo "path/to/objectfile" -r b829f2d554623b22b6ee431dff294e71f24eab79ddc1f83435a2cc606ca53623 -fp "path/to/prooffile"
  ```

## License

MIT
