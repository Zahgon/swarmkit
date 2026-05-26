package store

type orCombinator struct {
	bys []By
}

func (b orCombinator) isBy() {
	_ = "STUB: not implemented"

	// Or returns a combinator that applies OR logic on all the supplied By
	// arguments.
	return
}

func Or(bys ...By) By { _ = "STUB: not implemented"; return *new(By) }
