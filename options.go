package validator

// Option represents a configurations option to be applied to validator during initialization.
type Option func(*Validate)

// WithRequiredStructEnabled enables required tag on non-pointer structs to be applied instead of ignored.
//
// This was made opt-in behaviour in order to maintain backward compatibility with the behaviour previous
// to being able to apply struct level validations on struct fields directly.
//
// It is recommended you enabled this as it will be the default behaviour in v11+
func WithRequiredStructEnabled() Option {
	return func(v *Validate) {
		v.requiredStructEnabled = true
	}
}

// WithRequiredStructEnabled enables omitting the name of embedded anonymous structs from the namespace.
func WithOmitAnonymousName() Option {
	return func(v *Validate) {
		v.omitAnonymousName = true
	}
}
