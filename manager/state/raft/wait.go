package raft

import (
	"sync"
)

type waitItem struct {
	// channel to wait up the waiter
	ch chan interface{}
	// callback which is called synchronously when the wait is triggered
	cb func()
	// callback which is called to cancel a waiter
	cancel func()
}

type wait struct {
	l sync.Mutex
	m map[uint64]waitItem
}

func newWait() *wait { _ = "STUB: not implemented"; return nil }

func (w *wait) register(id uint64, cb func(), cancel func()) <-chan interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (w *wait) trigger(id uint64, x interface{}) bool { _ = "STUB: not implemented"; return false }

func (w *wait) cancel(id uint64) { _ = "STUB: not implemented"; return }

func (w *wait) cancelAll() { _ = "STUB: not implemented"; return }
