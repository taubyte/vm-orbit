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

# Creating A Plugin 

## Exported Functions
* Plugins can be created from structures with method names starting with **W_**
* The signature parameters can include the following: 
    * context, module, uint8, uint16, uint32, uint64, int8, int16, int32, int64 
        * module is used use for reading and writing to/from module memory
        * though context and module are optional, if used they must be in the first and second parameters, respectfully 
#### Example
```go 
type tester struct{}

import "github.com/taubyte/vm-orbit/satellite"

func (t *tester) W_hello(ctx context.Context, module satellite.Module, num uint32) uint32 {
	f, _ := os.Create("/tmp/hello.txt")
	defer f.Close()

	fmt.Println(module)
	fmt.Fprintln(f, "The answer is:", num)

	return 0
}

func (t *tester) W_sum(a, b int64) int64 {
	return a + b
}
```

## Building The Plugin
* Plugins are meant to be built as a binary to be injected for use the Taubyte VM
* To do so a main() function must be declared in a main package 
* Utilizing the plugin.Serve function the structured plugin methods can be exported 
#### Example
```go
package main

import 	"github.com/taubyte/vm-orbit/plugin"

func main() {
	plugin.Serve("testing", &tester{})
}
```
Then Built with: 
``` bash
go build 
```

* This binary can be referenced with shape deployment via spore-drive

# License
Please see the LICENSE file for details.


# Help
Find us on our [Discord](https://discord.gg/taubyte)


# Maintainers
 - Samy Fodil @samyfodil
 - Tafseer Khan @tafseer-khan

