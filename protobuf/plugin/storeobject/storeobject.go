package storeobject

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/moby/swarmkit/v2/protobuf/plugin"
)

// FIXME(aaronl): Look at fields inside the descriptor instead of
// special-casing based on name.
var typesWithNoSpec = map[string]struct{}{
	"Task":      {},
	"Resource":  {},
	"Extension": {},
}

type storeObjectGen struct {
	*generator.Generator
	generator.PluginImports
	eventsPkg  generator.Single
	stringsPkg generator.Single
}

func init() {
	generator.RegisterPlugin(new(storeObjectGen))
}

func (d *storeObjectGen) Name() string { _ = "STUB: not implemented"; return "" }

func (d *storeObjectGen) Init(g *generator.Generator) { _ = "STUB: not implemented"; return }

func (d *storeObjectGen) genMsgStoreObject(m *generator.Descriptor, storeObject *plugin.StoreObject) {
	_ = "STUB: not implemented"
	return
}

// Generate event types

// generate the event object type interface for this type
// event types implement some empty interfaces, for ease of use, like such:
//
//   type EventCreate interface {
//     IsEventCreatet() bool
//   }
//
//   type EventNode interface {
//     IsEventNode() bool
//   }
//
// then, each event has the corresponding interfaces implemented for its
// type. for example:
//
//   func (e EventCreateNode) IsEventCreate() bool {
//     return true
//   }
//
//   func (e EventCreateNode) IsEventNode() bool {
//     return true
//   }
//
// this lets the user filter events based on their interface type.
// note that the event type for each object type needs to be generated for
// each object. the event change type (Create/Update/Delete) is
// hand-written in the storeobject.go file because they are only needed
// once.

// implement event change type interface (IsEventCreate)

// implement event object type interface (IsEventNode)

// Generate methods for this type

// Generate event check functions

// Node is a special case

// Node is a special case

// Node is a special case

// Node is a special case

// Generate Convert*Watch function, for watch API.

/*                switch v := filter.By.(type) {
default:
        return nil, status.Errorf(codes.InvalidArgument, "selector type %T is unsupported for tasks", filter.By)
}
*/

// Generate indexer by ID

// Add the null character as a terminator

// Generate indexer by name

// Add the null character as a terminator

// Generate custom indexer

func (d *storeObjectGen) genFromArgs(indexerName string) { _ = "STUB: not implemented"; return }

func (d *storeObjectGen) genPrefixFromArgs(indexerName string) { _ = "STUB: not implemented"; return }

func (d *storeObjectGen) genNewStoreAction(topLevelObjs []string) {
	_ = "STUB: not implemented"
	// Generate NewStoreAction
	return
}

func (d *storeObjectGen) genWatchMessageEvent(topLevelObjs []string) {
	_ = "STUB: not implemented"
	// Generate WatchMessageEvent
	return
}

func (d *storeObjectGen) genEventFromStoreAction(topLevelObjs []string) {
	_ = "STUB: not implemented"
	// Generate EventFromStoreAction
	return
}

func (d *storeObjectGen) genConvertWatchArgs(topLevelObjs []string) {
	_ = "STUB: not implemented"
	// Generate ConvertWatchArgs
	return
}

func (d *storeObjectGen) Generate(file *generator.FileDescriptor) {
	_ = "STUB: not implemented"
	return
}

// no StoreObject extension

// for watch API
