package piscine

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb <= 21 {
		result := 1
		for i := 1; i <= nb; i++ {
			result = i * result
		}
		return result
	} else {
		return 0
	}
}
