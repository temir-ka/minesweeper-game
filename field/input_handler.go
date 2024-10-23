package field

// import "fmt"

// Function to convert string to an integer
func Atoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i == 0 || s[i] == '+' && i == 0 {
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
	}
	minus := 1
	if s[0] == '-' {
		minus = -1
		s = s[1:]
	}
	if s[0] == '+' {
		s = s[1:]
	}
	num := 0
	for i := 0; i < len(s); i++ {
		num *= 10
		num += int(s[i] - '0')
	}
	num *= minus
	return num
}

// Error handling for Height, Width
func HWError(s string) int {
	if len(s) < 1 || len(s) > 3 || s[0] == '0' {
		return -1 // Invalid length or leading zero
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
	}
	num := Atoi(s) // Convert to integer
	if num >= 100 || num < 3 {
		return -1
	}
	return num
}

// Check the custom map strings
func FieldError(s string, width int) bool {
	if len(s) == 0 { // Empty string is invalid
		return false
	}
	// Trim leading and trailing spaces
	left := 0
	for s[left] == ' ' {
		left++
	}
	right := len(s)
	for s[right-1] == ' ' {
		right--
	}
	s = s[left:right]
	if len(s) != width { // Length doesn't match the width
		return false
	}
	for i := 0; i < width; i++ {
		if s[i] != '.' && s[i] != '*' {
			return false
		}
	}
	return true
}

// Error handling for input coordinates
func RCError(s string) int {
	if len(s) < 1 || len(s) > 3 || s[0] == '0' {
		return -1 // Invalid length or leading zero
	}
	for i := 0; i < len(s); i++ { // Check if the input is integer
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
	}
	num := Atoi(s)
	if num >= 100 || num < 1 {
		return -1
	}
	return num
}

// Check for user choice of 1 and 2
func ChoiceError(s string) int {
	if len(s) != 1 { // Only one number
		return -1
	}
	if s[0] < '1' || s[0] > '2' {
		return -1
	}
	num := Atoi(s)
	if num < 1 || num > 2 { // It can be only 1 and 2
		return -1
	}
	return num
}
