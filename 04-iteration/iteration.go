package iteration

// Repeat ...
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		//repeated = repeated + character
		repeated += character
	}
	return repeated
}
