package lister

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListRunesWithinRangeShouldBeSuccess(t *testing.T) {
	got, err := ListRunesWithinRange('a', 'z')
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

func TestListRunesWithinRangeShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := ListRunesWithinRange('z', 'a')
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestListStringsWithinRangeShouldBeSuccess(t *testing.T) {
	got, err := ListStringsWithinRange('a', 'z')
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

func TestListStringsWithinRangeShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := ListStringsWithinRange('z', 'a')
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestListRunesWithinRangeStringShouldBeSuccess(t *testing.T) {
	got, err := ListRunesWithinRangeString(`u{3041}-u{304A}`)
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

func TestListStringsWithinRangeStringShouldBeSuccess(t *testing.T) {
	got, err := ListStringsWithinRangeString(`u{3041}-u{304A}`)
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

func TestListStringsWithinRangeStringShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := ListStringsWithinRangeString(`u{304A}-u{3041}`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestListRunesWithinRangeStringWithEmojiShouldBeSuccess(t *testing.T) {
	got, err := ListRunesWithinRangeString(`u{1F98A}-u{1F98D}`)
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

func TestListStringsWithinRangeStringWithEmojiShouldBeSuccess(t *testing.T) {
	got, err := ListStringsWithinRangeString(`u{1F98A}-u{1F98D}`)
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

func TestListRunesWithinRangeStringShouldBeFailWhenFromAndToAreInverted(t *testing.T) {
	got, err := ListRunesWithinRangeString(`u{304A}-u{3041}`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestListRunesWithinRangeStringShouldBeFailWhenRangeStringIsInvalid(t *testing.T) {
	got, err := ListRunesWithinRangeString(`INVA-LID`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}

func TestListStringsWithinRangeStringShouldBeFailWhenRangeStringIsInvalid(t *testing.T) {
	got, err := ListStringsWithinRangeString(`INVA-LID`)
	if err == nil {
		t.Fatal("err is nil")
	}
	fmt.Printf("%s\n", err)

	if got != nil {
		t.Fatal("got result is not nil")
	}
}
