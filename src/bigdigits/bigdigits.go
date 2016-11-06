package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var digits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		"    1  ",
		"   11  ",
		"  1 1  ",
		"    1  ",
		"    1  ",
		"    1  ",
		"   111 ",
	},
	{
		"  222  ",
		" 2   2 ",
		"    2  ",
		"   2   ",
		"  2    ",
		" 2     ",
		" 22222 ",
	},
	{
		"  333  ",
		" 3   3 ",
		"     3 ",
		"   33  ",
		"     3 ",
		" 3   3 ",
		"  333  ",
	},
	{
		"    4  ",
		"   44  ",
		"  4 4  ",
		" 4  4  ",
		" 444444",
		"    4  ",
		"    4  ",
	},
	{
		" 55555 ",
		" 5     ",
		" 5     ",
		"  555  ",
		"     5 ",
		" 5   5 ",
		"  555  ",
	},
	{
		"  666  ",
		" 6     ",
		" 6     ",
		" 6666  ",
		" 6   6 ",
		" 6   6 ",
		" 6666  ",
	},
	{
		" 77777 ",
		"     7 ",
		"    7  ",
		"   7   ",
		"  7    ",
		" 7     ",
		" 7     ",
	},
	{
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
	},
	{
		"  9999 ",
		" 9   9 ",
		" 9   9 ",
		"  9999 ",
		"     9 ",
		"     9 ",
		"  999  ",
	},
}

/*
Playground app to magnify entered digits. Exercise for slices and runes.
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <n>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	//traverse scanlines from top to bottom
	for scanline := 0; scanline < len(digits[0]); scanline++ {
		line := ""
		//traverse digits left to right
		for _, v := range os.Args[1] {
			//get int value from rune
			var digit = v - '0'
			if digit < 0 || digit > 9 {
				fmt.Println("Expected positive integer for n")
				os.Exit(2)
			} else {
				//add the string part of this digit in this scanline
				line = line + digits[digit][scanline]
			}
		}
		fmt.Println(line)
	}
}
