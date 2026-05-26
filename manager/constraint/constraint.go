package constraint

import (
	"regexp"

	"github.com/moby/swarmkit/v2/api"
)

const (
	eq = iota
	noteq

	// NodeLabelPrefix is the constraint key prefix for node labels.
	NodeLabelPrefix = "node.labels."
	// EngineLabelPrefix is the constraint key prefix for engine labels.
	EngineLabelPrefix = "engine.labels."
)

var (
	alphaNumeric = regexp.MustCompile(`^(?i)[a-z_][a-z0-9\-_.]+$`)
	// value can be alphanumeric and some special characters. it shouldn't container
	// current or future operators like '>, <, ~', etc.
	valuePattern = regexp.MustCompile(`^(?i)[a-z0-9:\-_\s\.\*\(\)\?\+\[\]\\\^\$\|\/]+$`)

	// operators defines list of accepted operators
	operators = []string{"==", "!="}
)

// Constraint defines a constraint.
type Constraint struct {
	key      string
	operator int
	exp      string
}

// Parse parses list of constraints.
func Parse(env []string) ([]Constraint, error) { _ = "STUB: not implemented"; return nil, nil }

// each expr is in the form of "key op value"

// split with the op

// validate key

// validate Value

// TODO(dongluochen): revisit requirements to see if globing or regex are useful

// found an op, move to next entry

// Match checks if the Constraint matches the target strings.
func (c *Constraint) Match(whats ...string) bool {
	_ = "STUB: not implemented"

	// full string match
	return false
}

// case insensitive compare

// NodeMatches returns true if the node satisfies the given constraints.
func NodeMatches(constraints []Constraint, n *api.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// if this node doesn't have hostname
// it's equivalent to match an empty hostname
// where '==' would fail, '!=' matches

// single IP address, node.ip == 2001:db8::2

// CIDR subnet, node.ip != 210.8.4.0/24

// reject constraint with malformed address/network

// node labels constraint in form like 'node.labels.key==value'

// label itself is case sensitive

// engine labels constraint in form like 'engine.labels.key!=value'

// key doesn't match predefined syntax
