package field

import "github.com/alem-platform/ap"

// number output
func PutNumber(n int) {
	// case for number 0
	if n == 0 {
		ap.PutRune('0')
		return
	}
	// case for negative number
	sign := 1
	if n < 0 {
		sign = -1
		n *= -1
	}
	// dividing integer into digits
	var reverse int = 0
	var zero int = 0
	for n > 0 {
		if n%10 == 0 {
			zero++
		} else {
			break
		}
		n /= 10
	}
	// reversing them
	for n > 0 {
		reverse *= 10
		reverse += (n % 10)
		n /= 10
	}
	// creating a whole number including negative signs also
	if sign == -1 {
		ap.PutRune('-')
	}
	for reverse > 0 {
		ap.PutRune(rune((reverse % 10) + '0'))
		reverse /= 10
	}
	for i := 0; i < zero; i++ {
		ap.PutRune('0')
	}
}

func LenNumber(n int) int {
	len := 0
	for n > 0 {
		n /= 10
		len++
	}
	return len
}
