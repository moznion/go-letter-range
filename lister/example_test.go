package lister

import "fmt"

func ExampleListRunesWithinRange() {
	got, err := ListRunesWithinRange('a', 'z')
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}

func ExampleListRunesWithinRangeString() {
	got, err := ListRunesWithinRangeString("u{0061}-u{007A}")
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}

func ExampleListStringsWithinRange() {
	got, err := ListStringsWithinRange('a', 'z')
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a string slice that contains "a" to "z" single character strings
}

func ExampleListStringsWithinRangeString() {
	got, err := ListStringsWithinRangeString("u{0061}-u{007A}")
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a string slice that contains "a" to "z" single character strings
}
