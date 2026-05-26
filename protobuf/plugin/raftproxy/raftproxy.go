package raftproxy

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type raftProxyGen struct {
	gen *generator.Generator
}

func init() {
	generator.RegisterPlugin(new(raftProxyGen))
}

func (g *raftProxyGen) Init(gen *generator.Generator) { _ = "STUB: not implemented"; return }

func (g *raftProxyGen) Name() string { _ = "STUB: not implemented"; return "" }

func (g *raftProxyGen) genProxyStruct(s *descriptor.ServiceDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genProxyConstructor(s *descriptor.ServiceDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genRunCtxMods(s *descriptor.ServiceDescriptorProto) {
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

func (g *raftProxyGen) genStreamWrapper(streamType string) {
	_ = "STUB: not implemented"
	// Generate stream wrapper that returns a modified context
	return
}

func (g *raftProxyGen) genClientStreamingMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

// Generate stream wrapper that returns a modified context

func (g *raftProxyGen) genServerStreamingMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genClientServerStreamingMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genSimpleMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genProxyMethod(s *descriptor.ServiceDescriptorProto, m *descriptor.MethodDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) genPollNewLeaderConn(s *descriptor.ServiceDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (g *raftProxyGen) Generate(file *generator.FileDescriptor) { _ = "STUB: not implemented"; return }

func (g *raftProxyGen) GenerateImports(file *generator.FileDescriptor) {
	_ = "STUB: not implemented"
	return
}

// don't conflict with import added by ptypes
