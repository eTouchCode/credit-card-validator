package main

func luhnAlgorithm(cardNumber string) bool {
	// this function implements the luhn algorithm
	// it takes as argument a cardnumber of type string
	// and it returns a boolean (true or false) if the
	// card number is valid or not

	// initialise a variable to keep track of the total sum of digits
	total := 0
	// Initialize a flag to track whether the current digit is the second digit from the right.
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0
}
