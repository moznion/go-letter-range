package letterrange

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRunesWithinRangeShouldBeSuccess(t *testing.T) {
	got, err := GetRunesWithinRange('a', 'z')
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 26 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetRunesWithinRangeShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := GetRunesWithinRange('z', 'a')
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestGetStringsWithinRangeShouldBeSuccess(t *testing.T) {
	got, err := GetStringsWithinRange('a', 'z')
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 26 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetStringsWithinRangeShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := GetStringsWithinRange('z', 'a')
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestGetRunesWithinRangeStringShouldBeSuccess(t *testing.T) {
	got, err := GetRunesWithinRangeString(`u{3041}-u{304A}`)
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 10 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []rune{'ぁ', 'あ', 'ぃ', 'い', 'ぅ', 'う', 'ぇ', 'え', 'ぉ', 'お'}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetStringsWithinRangeStringShouldBeSuccess(t *testing.T) {
	got, err := GetStringsWithinRangeString(`u{3041}-u{304A}`)
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 10 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []string{"ぁ", "あ", "ぃ", "い", "ぅ", "う", "ぇ", "え", "ぉ", "お"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetStringsWithinRangeStringShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := GetStringsWithinRangeString(`u{304A}-u{3041}`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestGetRunesWithinRangeStringWithEmojiShouldBeSuccess(t *testing.T) {
	got, err := GetRunesWithinRangeString(`u{1F98A}-u{1F98D}`)
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 4 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []rune{'🦊', '🦋', '🦌', '🦍'}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetStringsWithinRangeStringWithEmojiShouldBeSuccess(t *testing.T) {
	got, err := GetStringsWithinRangeString(`u{1F98A}-u{1F98D}`)
	if err != nil {
		t.Fatal("err is not nil")
	}

	if len(got) != 4 {
		t.Fatal("length of got slice is invalid")
	}

	expected := []string{"🦊", "🦋", "🦌", "🦍"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatal("got result doesn't equal to expected")
	}
}

func TestGetRunesWithinRangeStringShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := GetRunesWithinRangeString(`u{304A}-u{3041}`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestGetRunesWithinRangeStringShouldBeFailWhenRangeStringIsInvalid(t *testing.T) {
	got, err := GetRunesWithinRangeString(`INVA-LID`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestGetStringsWithinRangeStringShouldBeFailWhenRangeStringIsInvalid(t *testing.T) {
	got, err := GetStringsWithinRangeString(`INVA-LID`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}
