package deepcopy

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type deepCopyGen struct {
	*generator.Generator
	generator.PluginImports
	copyPkg generator.Single
}

func init() {
	generator.RegisterPlugin(new(deepCopyGen))
}

func (d *deepCopyGen) Name() string { _ = "STUB: not implemented"; return "" }

func (d *deepCopyGen) Init(g *generator.Generator) { _ = "STUB: not implemented"; return }

func (d *deepCopyGen) genCopyFunc(dst, src string) { _ = "STUB: not implemented"; return }

func (d *deepCopyGen) genCopyBytes(dst, src string) { _ = "STUB: not implemented"; return }

// allocate dst object

// copy bytes from src to dst

func (d *deepCopyGen) genMsgDeepCopy(m *generator.Descriptor) { _ = "STUB: not implemented"; return }

// Generate backwards compatible, type-safe Copy() function.

// shallow copy handles all scalars

// Handle oneof type, we defer them to a loop below

// Handle all kinds of message type

// Handle map type

// Handle any message which is not repeated or part of oneof

// allocate dst object

// copy into the allocated struct

// Handle repeated field

// Handle bytes

// skip: field was a scalar handled by shallow copy!

func (d *deepCopyGen) genMap(_ *generator.Descriptor, f *descriptor.FieldDescriptorProto) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *deepCopyGen) genRepeated(m *generator.Descriptor, f *descriptor.FieldDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

// TODO(stevvooe): Handle custom type here?
// elides [] or *

func (d *deepCopyGen) genOneOf(m *generator.Descriptor, oneof *descriptor.OneofDescriptorProto, fields []*descriptor.FieldDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

// elides [] or *

func (d *deepCopyGen) Generate(file *generator.FileDescriptor) { _ = "STUB: not implemented"; return }

// TODO(stevvooe): Ideally, this could be taken as a parameter to the
// deepcopy plugin to control the package import, but this is good enough,
// for now.
