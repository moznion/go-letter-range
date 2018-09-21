package letterrange

import "fmt"

func ExampleGetRunesWithinRange() {
	got, err := GetRunesWithinRange('a', 'z')
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}

func ExampleGetRunesWithinRangeString() {
	got, err := GetRunesWithinRangeString("u{0061}-u{007A}")
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a rune slice that contains 'a' to 'z' runes
}

func ExampleGetStringsWithinRange() {
	got, err := GetStringsWithinRange('a', 'z')
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a string slice that contains "a" to "z" single character strings
}

func ExampleGetStringsWithinRangeString() {
	got, err := GetStringsWithinRangeString("u{0061}-u{007A}")
	if err != nil {
		// do error handling
		return
	}
	fmt.Println(got) // out: a string slice that contains "a" to "z" single character strings
}
