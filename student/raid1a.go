package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	x--
	y--
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			if i == 0 && j == 0 ||
				i == y && j == x ||
				i == y && j == 0 ||
				i == 0 && j == x {
				z01.PrintRune('o')
			} else if i == 0 || i == y {
				z01.PrintRune('-')
			} else if j == 0 || j == x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(32)
			}
		}
		z01.PrintRune(10)
	}
}
