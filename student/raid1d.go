package student

import "github.com/01-edu/z01"

// Raid1d function prints a square
// of width x and of height y
func Raid1d(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')

		if y >= 1 {
			for i := 0; i < (x - 2); i++ {
				z01.PrintRune('B')
			}
			if x >= 2 {
				z01.PrintRune('C')
			}
			z01.PrintRune('\n')

			for j := 0; j < (y - 2); j++ {
				z01.PrintRune('B')
				for i := 0; i < (x - 2); i++ {
					z01.PrintRune(' ')
				}
				if x >= 2 {
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
			}

			if y >= 2 {
				z01.PrintRune('A')
				for i := 0; i < (x - 2); i++ {
					z01.PrintRune('B')
				}
				if x >= 2 {
					z01.PrintRune('C')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
