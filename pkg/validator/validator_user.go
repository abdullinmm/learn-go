package validator

func ValidatorUser(v Validator) error {
	return v.Validate()
}
