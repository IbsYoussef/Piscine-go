package main

import "github.com/01-edu/z01"

func main() {
	PrintComb2()
}

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '1'; j <= '9'; j++ {
			z01.PrintRune('0')
			z01.PrintRune('0')
			z01.PrintRune(' ')
			z01.PrintRune(i)
			z01.PrintRune(j)
			z01.PrintRune(',')
			z01.PrintRune(' ')
			if i == '9' && j == '9' {
				break
			}
		}
	}
}
