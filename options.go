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

// WithPrivateFieldValidation activates validation for unexported fields via the use of the `unsafe` package.
//
// By opting into this feature you are acknowledging that you are aware of the risks and accept any current or future
// consequences of using this feature.
func WithPrivateFieldValidation() Option {
	return func(v *Validate) {
		v.privateFieldValidation = true
	}
}

// WithEarlyExit configures the validator to immediately stop validation as soon as the first error is encountered
//
// This feature could be an opt-in behavior, allowing to opt into "early exit" validation, without breaking current workflows
// Early exit on the first failure would save time by avoiding unnecessary checks on the remaining fields
func WithEarlyExit() Option {
	return func(v *Validate) {
		v.earlyExit = true
	}
}
