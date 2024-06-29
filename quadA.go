package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				// Print the top and bottom edges
				z01.PrintRune('o')
				for j := 1; j < x-1; j++ {
					z01.PrintRune('-')
				}
				if x > 1 {
					z01.PrintRune('o')
				}
			} else {
				// Print the sides with spaces in between
				z01.PrintRune('|')
				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('|')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
