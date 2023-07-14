package fixtures

import _ "embed"

//go:generate tar -czvf go.tar -C go/ .

//go:embed go.tar
var GoFixture []byte
