# vm-orbit
A tool for creating plugin binaries to be deployed on a Taubyte Cloud 

# Building Proto 

## Install go protoc gen
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"
```

## Build 
```bash 
cd <path/to/vm-orbit>
``` 
```bash
go generate ./...
```

# License
Please see the LICENSE file for details.


# Help
Find us on our [Discord](https://discord.gg/taubyte)


# Maintainers
 - Samy Fodil @samyfodil
 - Tafseer Khan @tafseer-khan

