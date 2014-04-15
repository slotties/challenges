package main

import (
	"math"
)

/*
	Challenge:
	I had the other day in my possession a label bearing the number 3 0 2 5 in large figures.
	This got accidentally torn in half, so that 3 0 was on one piece and 2 5 on the other.
	On looking at these pieces I began to make a calculation, when I discovered this little
	peculiarity. If we add the 3 0 and the 2 5 together and square the sum we get as the result,
	the complete original number on the label! Thus, 30 added to 25 is 55, and 55 multiplied by 55
	is 3025. Curious, is it not?
	Now, the challenge is to find another number, composed of four figures, all different, which
	may be divided in the middle and produce the same result.
*/

func IsTornNumber(number int) bool {
	if number < 1000 || number > 9999 {
		return false
	}

	a := int(math.Floor(float64(number) / 100.0))
	b := number % 100
	ab := a + b

	if (ab * ab) == number {
		/* Okay, the number *seems* to be a torn number, but the challenge has a special described:
		 * "Now, the challenge is to find another number, composed of four figures, all different, 
		 *  which may be divided in the middle and produce the same result."
		 * The point is: "all different"
		 * Therefore we have to check if the number is composed of unique digits (aka figures).
		 */
		return hasUniqueFigures(number)
	} else {
		return false
	}
}

/*
 * Run through all digits and store found digits in a bit mask. As soon as a digit was found twice
 * the loop ends.
 * I wonder if there's a faster way to check for duplicate (or unique) digits but I haven't found
 * one so far.
 */
func hasUniqueFigures(numberInt int) bool {
	number := float64(numberInt)
	var foundFigures int

	for i := 0; i < 4; i++ {
		digit := uint(int(number) % 10)
		bitIdxMask := 1 << digit

		if (foundFigures & bitIdxMask) == bitIdxMask {
			return false
		} else {
			foundFigures |= bitIdxMask
		}

		number = math.Floor(number / 10.0)
	}
	
	return true
}