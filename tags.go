package mole

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

const (
	StructTAG = "mole"
)

// written so that future operators are easier to add
const (
	invalidCtlRunes = "!#$%&()+./:*=?@[]^_{|}~ " // including space rune
	validCtlRunes   = "<>,-"                     // comma dash(future use)
)

var (
	// ErrInvaidCharInStructTag is returned if struct tag contains any invalid runes
	// character in struct tag must be a number, letter, "," or "-"
	ErrInvaidCharInStructTag = errors.New("invalid character in struct tags")
	// ErrInvaidCharInStructTag is returned if struct tag is empty
	ErrEmptyStructTag = errors.New("empty struct tags")
)

func getTag(sf reflect.StructField) (string, optionTags, error) {
	return readTag(sf.Tag.Get(StructTAG))
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
				return tagValue, optionTags{}, nil
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
	return tag, optionTags([]string{})
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
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !strings.ContainsRune(validCtlRunes, c) {
				return ErrInvaidCharInStructTag
			}
		}
	}
	return nil
}
