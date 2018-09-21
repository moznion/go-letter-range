package letterrange

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// GetRunesWithinRange returns the rune slice that contains letters within the specified range.
func GetRunesWithinRange(from, to rune) ([]rune, error) {
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

// GetStringsWithinRange returns the string slice that contains letters within the specified range.
func GetStringsWithinRange(from, to rune) ([]string, error) {
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

// GetRunesWithinRangeString returns the rune slice that contains letters within the specified range-string.
// Range-string must have following format: `u{hex}-u{hex}`. The first hex is handled as `from` and the second one is handles as `to`.
func GetRunesWithinRangeString(rangeString string) ([]rune, error) {
	from, to, err := extractRangeFromRangeString(rangeString)
	if err != nil {
		return nil, err
	}

	return GetRunesWithinRange(from, to)
}

// GetStringsWithinRangeString returns the string slice that contains letters within the specified range-string.
// Range-string must have following format: `u{hex}-u{hex}`. The first hex is handled as `from` and the second one is handles as `to`.
func GetStringsWithinRangeString(rangeString string) ([]string, error) {
	from, to, err := extractRangeFromRangeString(rangeString)
	if err != nil {
		return nil, err
	}

	return GetStringsWithinRange(from, to)
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
