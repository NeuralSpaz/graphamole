package mole

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

// StructTagString modify to suit your struct tag scheme
const StructTagString = "mole"

// written so that future operators are easier to add or remove
const (
	invalidCtlRunes = "!#$%&()+./*=?@[]^_{|}~ " // including space rune
	validCtlRunes   = "<>,-:"                   // comma dash(future use)
)

// ErrInvaidCharInStructTag is returned if struct tag contains any invalid runes
// character in struct tag must be a number, letter, ",", "<",">" or "-"
var ErrInvaidCharInStructTag = errors.New("error: invalid character in struct tag")

// ErrEmptyStructTag is returned if struct tag is empty
var ErrEmptyStructTag = errors.New("error: empty struct tag")

// ErrEmptyStructOptions is returned if struct tag options separator is present
// without an value
var ErrEmptyStructOptions = errors.New("error: empty struct options tag")

func getTag(sf reflect.StructField) (string, optionTags, error) {
	return readTag(sf.Tag.Get(StructTagString))
}

// readTag() wrapped for ease of testing
func readTag(st string) (string, optionTags, error) {
	tagValue, opts := splitTag(st)
	if err := validateTag(tagValue); err != nil {
		return "", optionTags{}, err
	}
	for _, v := range opts {
		if err := validateTag(v); err != nil {
			if err == ErrEmptyStructTag {
				return "", optionTags{}, ErrEmptyStructOptions
			}
			return "", optionTags{}, err
		}
	}
	return tagValue, opts, nil
}

type optionTags []string

// split tag into tag and []options
func splitTag(tag string) (string, optionTags) {
	if tags := strings.Split(tag, ","); len(tags) > 1 {
		return tags[0], optionTags(tags[1:])
	}
	return tag, optionTags{}
}

func validateTag(s string) error {
	if s == "" {
		return ErrEmptyStructTag
	}
	// return early on any invalid charactors
	for _, c := range s {
		switch {
		case strings.ContainsRune(invalidCtlRunes, c):
			return ErrInvaidCharInStructTag
		default:
			if !unicode.IsLetter(c) &&
				!unicode.IsDigit(c) &&
				!strings.ContainsRune(validCtlRunes, c) {
				return ErrInvaidCharInStructTag
			}
		}
	}
	return nil
}
