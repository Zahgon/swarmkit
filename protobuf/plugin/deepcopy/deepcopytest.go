package deepcopy

import (
	"github.com/gogo/protobuf/plugin/testgen"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type test struct {
	*generator.Generator
}

// NewTest creates a new deepcopy testgen plugin
func NewTest(g *generator.Generator) testgen.TestPlugin {
	_ = "STUB: not implemented"
	return *new(testgen.TestPlugin)
}

func (p *test) Generate(imports generator.PluginImports, file *generator.FileDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

// copying from nil should result in nil

func init() {
	testgen.RegisterTestPlugin(NewTest)
}
