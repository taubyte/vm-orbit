package goBuilder

import _ "embed"

//go:generate tar -czvf fixtures.tar -C fixtures .

//go:embed fixtures.tar
var fixture []byte
