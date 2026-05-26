package authenticatedwrapper

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/moby/swarmkit/v2/protobuf/plugin"
)

type authenticatedWrapperGen struct {
	gen *generator.Generator
}

func init() {
	generator.RegisterPlugin(new(authenticatedWrapperGen))
}

func (g *authenticatedWrapperGen) Init(gen *generator.Generator) { _ = "STUB: not implemented"; return }

func (g *authenticatedWrapperGen) Name() string { _ = "STUB: not implemented"; return "" }

func (g *authenticatedWrapperGen) genAuthenticatedStruct(s *descriptor.ServiceDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) genAuthenticatedConstructor(s *descriptor.ServiceDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func getInputTypeName(m *descriptor.MethodDescriptorProto) string {
	_ = "STUB: not implemented"
	return ""
}

func getOutputTypeName(m *descriptor.MethodDescriptorProto) string {
	_ = "STUB: not implemented"
	return ""
}

func serviceTypeName(s *descriptor.ServiceDescriptorProto) string {
	_ = "STUB: not implemented"
	return ""
}

func sigPrefix(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) string {
	_ = "STUB: not implemented"
	return ""
}

func genRoles(auth *plugin.TLSAuthorization) string { _ = "STUB: not implemented"; return "" }

func (g *authenticatedWrapperGen) genServerStreamingMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) genClientServerStreamingMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) genSimpleMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) genAuthenticatedMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) Generate(file *generator.FileDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (g *authenticatedWrapperGen) GenerateImports(_ *generator.FileDescriptor) {
	_ = "STUB: not implemented"
	return
}
