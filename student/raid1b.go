package student

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	x--
	y--
	if y == 0 {
		for j := 0; j <= x; j++ {
			if j == 0 {
				z01.PrintRune('/')
			} else if j == x {
				z01.PrintRune('\\')
			} else {
				z01.PrintRune('*')
			}
		}
		z01.PrintRune(10)
	} else if x == 0 {
		for i := 0; i <= y; i++ {
			if i == 0 {
				z01.PrintRune('/')
			} else if i == y {
				z01.PrintRune('\\')
			} else {
				z01.PrintRune('*')
			}
			z01.PrintRune(10)
		}
	} else {

		for i := 0; i <= y; i++ {
			for j := 0; j <= x; j++ {
				if i == y && j == 0 ||
					i == 0 && j == x {
					z01.PrintRune('\\')
				} else if i == 0 && j == 0 ||
					i == y && j == x {
					z01.PrintRune('/')
				} else if i == 0 || i == y {
					z01.PrintRune('*')
				} else if j == 0 || j == x {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(32)
				}
			}
			z01.PrintRune(10)
		}
	}

}
