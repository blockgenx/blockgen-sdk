package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/blockgenx/blockgen-sdk/orm/internal/codegen"
)

func main() {
	protogen.Options{}.Run(codegen.PluginRunner)
}
