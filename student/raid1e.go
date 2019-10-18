package student

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	x--
	y--
	s1 := 'A'
	s2 := 'C'
	s3 := 'B'
	if y == 0 {
		for j := 0; j <= x; j++ {
			if j == 0 {
				z01.PrintRune(s1)
			} else if j == x {
				z01.PrintRune(s2)
			} else {
				z01.PrintRune(s3)
			}
		}
		z01.PrintRune(10)
	} else if x == 0 {
		for i := 0; i <= y; i++ {
			if i == 0 {
				z01.PrintRune(s1)
			} else if i == y {
				z01.PrintRune(s2)
			} else {
				z01.PrintRune(s3)
			}
			z01.PrintRune(10)
		}
	} else {

		for i := 0; i <= y; i++ {
			for j := 0; j <= x; j++ {
				if i == 0 && j == x ||
					i == y && j == 0 {
					z01.PrintRune(s2)
				} else if i == 0 && j == 0 ||
					i == y && j == x {
					z01.PrintRune(s1)
				} else if i == 0 || i == y {
					z01.PrintRune(s3)
				} else if j == 0 || j == x {
					z01.PrintRune(s3)
				} else {
					z01.PrintRune(32)
				}
			}
			z01.PrintRune(10)
		}
	}
}
