# vm-orbit
A tool for creating plugin binaries to be deployed on a Taubyte Cloud 

# Creating A Plugin 

## Exported Functions
* Plugins can be created from structures with method names starting with **W_**
* The signature parameters can include the following: 
    * context, module,uint, uint8, uint16, uint32, uint64,int, int8, int16, int32, int64 
        * module is used use for reading and writing to/from module memory
        * though context and module are optional, if used they must be in the first and second parameters, respectfully 
#### Example
```go 
type tester struct{}

import "github.com/taubyte/vm-orbit/satellite"

func (t *tester) W_atoiAdd(ctx context.Context, module satellite.Module, stringPtr uint32, stringLen uint32, addVal int,resPtr uint32) int {
	data, err := module.MemoryRead(stringPtr, stringLen)
	if err != nil {
		return 1
	}

	parsedVal, err := strconv.Atoi(string(data))
	if err != nil {
		return 1
	}

	sum := parsedVal + addVal
	sumString := fmt.Sprintf("added sum: %d",sum)

	n, err := module.MemoryWrite(resPtr, []byte(sumString))
	if err != nil{
		return 1
	}
	
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

# Testing 
Dreamland is a tool used to create a local taubyte network.
Plugins can be tested using dreamland before deploying to a production network. 

## Dreamland 

### Inject Plugin 
Start Dreamland 
```bash
dream new multiverse 
```

Once Network has been started (after  `SUCCESS  Universe <name-of-universe> started!`), inject plugin from a new terminal
```bash
dream inject attachPlugin -p <path/to/plugin/binary>
```

### Using in a DFunc
Create reference to the plugin module and function.
```go 
//go:wasm-module moduleName
//export functionName
func functionName(*byte, uint32) uint32
```
* The moduleName is the name given in the plugin.Export() method
* The name of the function is the name of your structure's method minus `W_`
* The signature follows 
### Example 






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

