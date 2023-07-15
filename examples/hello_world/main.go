package main

import "github.com/taubyte/vm-orbit/plugin"

func main() {
	// methods of helloWorlder will be exported to the module "helloWorld"
	plugin.Export("helloWorld", &helloWorlder{})
}
