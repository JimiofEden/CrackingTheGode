package ch01

/**
 * Will determine if one string is a permutation of another
 */

func urlify(inputString string, trueLength int) string {
	if inputString == "" || trueLength == 0 {
		return ""
	}

	inputString = inputString[:trueLength]

	for i := trueLength-1; i >= 0; i-- {
		if inputString[i] == ' ' {

			if i == 0 {
				inputString = "%20" + inputString[i+1:]
				continue
			}
			if i == trueLength-1 {
				inputString = inputString[:i] + "%20"
				continue
			}
			inputString = inputString[:i] + "%20" + inputString[i+1:]
		}
	}

	return inputString
}

