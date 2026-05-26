package controlapi

import (
	"regexp"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

var isValidDNSName = regexp.MustCompile(`^[a-zA-Z0-9](?:[-_]*[A-Za-z0-9]+)*$`)

// configs and secrets have different naming requirements from tasks and services
var isValidConfigOrSecretName = regexp.MustCompile(`^[a-zA-Z0-9]+(?:[a-zA-Z0-9-_.]*[a-zA-Z0-9])?$`)

func buildFilters(by func(string) store.By, values []string) store.By {
	_ = "STUB: not implemented"
	return *new(store.By)
}

func filterContains(match string, candidates []string) bool {
	_ = "STUB: not implemented"
	return false
}

func filterContainsPrefix(match string, candidates []string) bool {
	_ = "STUB: not implemented"
	return false
}

func filterMatchLabels(match map[string]string, candidates map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateAnnotations(m api.Annotations) error { _ = "STUB: not implemented"; return nil }

// if the name doesn't match the regex

// DNS labels are limited to 63 characters

func validateConfigOrSecretAnnotations(m api.Annotations) error {
	_ = "STUB: not implemented"
	return nil
}

// if the name doesn't match the regex
