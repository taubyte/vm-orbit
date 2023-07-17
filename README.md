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

### Create a Dffunc on Console 
You can access a local instance of webconsole by running 
```bash 
dream status console 
```

Create a [DFUNC](#creating-a-dfunc-with-reference-to-a-plugin) 
It is recommended to make a http GET function to quickly test, using a generated domain

To call this http DFunc you will need to add the generated fqdn to /etc/host 127.0.0.1 
You will also need to append the port that Substrate(node) protocol is running on to the url
You can get this port by running
```bash
dream status node
```

after these two steps hitting `<generated-fqdn>:<node-port>/<function-path>` will execute your function

## Go Test 
`vm-orbit/tests/suite` can be used to create local plugin tests quickly 
For the testing the following required: 
* Plugin Binary
* Wasm File with an exported method calling your plugin method

Suite has helpers to generate these
### Using the Builder For Plugins and Wasm
You will need to create a builder for the language you will be using to generate your Wasm file and plugin

```go 
	import goBuilder "github.com/taubyte/vm-orbit/tests/suite/builders/go"

	builder := goBuilder.New()
```

#### Build Your Plugin
```go 
	pluginPath, err := builder.Plugin("path/to/plugin")
```

#### Build Your Wasm File 
You will need to write code files that will properly reference your plugins refer to: [DFUNC](#creating-a-dfunc-with-reference-to-a-plugin)

Then used the builder to generate the wasm file
```go 
	wasmPath, err := builder.Wasm(context.Background(), ...list-of-Files-To-Include)
```

### Using the Suite to Test Plugins

#### Create a Testing Suite 
```go
import "github.com/taubyte/vm-orbit/tests/suite"

testingSuite, err := suite.New(context.Background)
```

#### Attach a Plugin onto the Suite
```go
err := testingSuite.AttachPluginFromPath("path/to/plugin")
```

#### Attach the Wasm File to the Suite, and Get the Module
```go
module, err := testingSuite.WasmModule("path/to/wasmFile")
```

#### Call the Dfunc
```go
ret, err := module.Call(context.Background(),"functionName")
```

## Creating a Dfunc with reference to a Plugin

### Go

Example: 
```go 
//go:wasm-module moduleName
//export writeSize
func writeSize(*uint32) 

//go:wasm-module moduleName
//export writeName
func writeName(*byte)

//export dFunc
func dFunc() {
	var size uint32 
	writeSize(&size)

	nameData := make([]byte, size)
	writeName(&nameData[0])

	name := string(nameData)
}
```
* The comments before the function declarations are required
* the //go:wasm-module comment gives a reference to the name of the wasm module of the plugin 
	* Example:
	```go
	package main

	import "github.com/taubyte/vm-orbit/plugin"

	func main() {
		// methods of helloWorlder will be exported to the module "helloWorld"
		plugin.Export("helloWorld", &helloWorlder{})
	}
	```
	* Here the plugin module name is `helloWorld`
* The name of the function is the name of your structure's method minus `W_`
	* Example: 
	```go 
	func (t *helloWorlder) W_helloSize(ctx context.Context, module satellite.Module, sizePtr uint32) uint32 {
		if _, err := module.WriteStringSize(sizePtr, helloWorld); err != nil {
			return 1
		}

		return 0
	}
	```
	* Here the name of this method would be `helloSize`

#### Understanding The Signature in the DFunc 
A good rule of thumb is if a value needs to be read from memory or written to the signature value needs to be a pointer. If the value is interpreted a raw value should be passed.

Uints, Floats, and Ints and their pointers are supported types for the signature

Any data more complex than these types are interpreted as []byte which must be read or written to in memory. 

Any data more complex than []byte is encoded to []byte through the use of helper methods, already available in an orbit module 
```go 
import 	"github.com/taubyte/vm-orbit/satellite"

func (h *helloWorlder) W_readWrite(
	ctx context.Context,
	module satellite.Module,
	stringSlicePtr,
	stringSliceSize,
){
	stringSlice, err := module.ReadStringSlice(stringSlicePtr,stringSliceSize)
}


```

More encoding methods can be found at `github.com/taubyte/go-sdk/utils/codec`
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

