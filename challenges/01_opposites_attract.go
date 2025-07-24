// adapted from: https://www.codewars.com/kata/555086d53eac039a2a000083/train/go

// Timmy & Sarah think they are in love, but around where they live, they will
// only know once they pick a flower each. If one of the flowers has an even
// number of petals and the other has an odd number of petals it means they are in love.

// Write a function that will take the number of petals of each flower and
// return true if they are in love and false if they aren't.

package challenges

import (
	"math"
)

func LoveFunc(flower1, flower2 int) bool {
	flower1Mod := math.Mod(float64(flower1), 2)
	flower2Mod := math.Mod(float64(flower2), 2)
	if flower1Mod == 0 && flower2Mod == 1 {
		return true
	} else if flower1Mod == 1 && flower2Mod == 0 {
		return true
	} else {
		return false
	}
}

func OppositesAttract() {
	LoveFunc(1, 3)
}
