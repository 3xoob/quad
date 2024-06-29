package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			if i == 0 {
				// Print the top edges
				z01.PrintRune('A')
				for j := 1; j < x-1; j++ {
					z01.PrintRune('B')
				}
				if x > 1 {
					z01.PrintRune('C')
				}
			} else if i == y-1 {
				// Print  bottom edges
				z01.PrintRune('A')
				for j := 1; j < x-1; j++ {
					z01.PrintRune('B')
				}
				if x > 1 {
					z01.PrintRune('C')
				}

			} else {
				// Print the sides with spaces in between
				z01.PrintRune('B')
				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
