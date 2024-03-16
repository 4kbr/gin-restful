package helper

func ResultValidation(result error) error {
	if result != nil {
		return result
	}
	return nil

}
