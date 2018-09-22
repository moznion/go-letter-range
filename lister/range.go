// Package lister provides some functions to list letters according to specified letters range.
package lister

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// ListRunesWithinRange returns the rune slice that contains letters within the specified range.
func ListRunesWithinRange(from, to rune) ([]rune, error) {
	if from > to {
		return nil, errors.New("`from` is bigger than `to`: `from` must be smaller than `to`")
	}

	cur := from
	length := int(to - from + 1)
	rs := make([]rune, length)
	for i := 0; i < length; i, cur = i+1, cur+1 {
		rs[i] = cur
	}

	return rs, nil
}

// ListStringsWithinRange returns the string slice that contains letters within the specified range.
func ListStringsWithinRange(from, to rune) ([]string, error) {
	if from > to {
		return nil, errors.New("`from` is bigger than `to`: `from` must be smaller than `to`")
	}

	cur := from
	length := int(to - from + 1)
	ss := make([]string, length)
	for i := 0; i < length; i, cur = i+1, cur+1 {
		ss[i] = string(cur)
	}

	return ss, nil
}

// ListRunesWithinRangeString returns the rune slice that contains letters within the specified range-string.
// Range-string must have following format: `u{hex}-u{hex}`. The first hex is handled as `from` and the second one is handles as `to`.
func ListRunesWithinRangeString(rangeString string) ([]rune, error) {
	from, to, err := extractRangeFromRangeString(rangeString)
	if err != nil {
		return nil, err
	}

	return ListRunesWithinRange(from, to)
}

// ListStringsWithinRangeString returns the string slice that contains letters within the specified range-string.
// Range-string must have following format: `u{hex}-u{hex}`. The first hex is handled as `from` and the second one is handles as `to`.
func ListStringsWithinRangeString(rangeString string) ([]string, error) {
	from, to, err := extractRangeFromRangeString(rangeString)
	if err != nil {
		return nil, err
	}

	return ListStringsWithinRange(from, to)
}

var rangeStringRe = regexp.MustCompile("^u{([0-9a-fA-F]+)}-u{([0-9a-fA-F]+)}$")

func extractRangeFromRangeString(rangeString string) (rune, rune, error) {
	matched := rangeStringRe.FindStringSubmatch(rangeString)
	if matched == nil {
		return 0, 0, fmt.Errorf("given rangeString is invalid format (example: it should like be `%s`)", `u{3040}-u{309F}`)
	}

	from, _ := strconv.ParseUint(matched[1], 16, 64)
	to, _ := strconv.ParseUint(matched[2], 16, 64)

	return rune(from), rune(to), nil
}
