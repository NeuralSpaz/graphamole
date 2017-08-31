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

var (
	// ErrInvaidCharInStructTag is returned if struct tag contains any invalid runes
	// character in struct tag must be a number, letter, "," or "-"
	ErrInvaidCharInStructTag = errors.New("invalid character in struct tags")
)

func getTag(sf reflect.StructField) (string, optionTags, error) {
	return readTag(sf.Tag.Get(StructTAG))
}

// readTag() wrapped for ease of testing
func readTag(st string) (string, optionTags, error) {
	tagValue, opts := parseTag(st)
	if !isValidTag(tagValue) {
		return "", optionTags{}, ErrInvaidCharInStructTag
	}
	for _, v := range opts {
		if !isValidTag(v) {
			return "", optionTags{}, ErrInvaidCharInStructTag
		}
	}
	return tagValue, opts, nil
}

type optionTags []string

// split tag into tag and []options
func parseTag(tag string) (string, optionTags) {
	if tags := strings.Split(tag, ","); len(tags) > 1 {
		return tags[0], optionTags(tags[1:])
	}
	return tag, optionTags([]string{})
}

// written so that future operators are easier to add
var invalidCtlRunes = "!#$%&()+./:*=?@[]^_{|}~ " // including space rune
var validCtlRunes = "<>,-"                       // comma dash(future use)

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	// return early on any invalid charactors
	for _, c := range s {
		switch {
		case strings.ContainsRune(invalidCtlRunes, c):
			return false
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !strings.ContainsRune(validCtlRunes, c) {
				return false
			}
		}
	}
	return true
}
