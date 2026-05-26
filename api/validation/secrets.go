package validation

// MaxSecretSize is the maximum byte length of the `Secret.Spec.Data` field.
const MaxSecretSize = 500 * 1024 // 500KB

// ValidateSecretPayload validates the secret payload size
func ValidateSecretPayload(data []byte) error { _ = "STUB: not implemented"; return nil }
