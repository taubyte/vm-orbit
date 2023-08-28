package main

import "github.com/taubyte/vm-orbit/satellite"

func main() {
	// methods of helloWorlder will be exported to the module "helloWorld"
	satellite.Export("helloWorld", &helloWorlder{})
}
