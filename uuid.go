package yab

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrEmptyUUID     = errors.New("empty UUID")
	ErrBadUUID       = errors.New("bad UUID")
	ErrDuplicateUUID = errors.New("duplicate UUID")
)

type UUIDValues []string

func (vs *UUIDValues) Set(v string) error {
	switch {
	case len(v) == 0:
		return ErrEmptyUUID
	case !IsUUID(v):
		return ErrBadUUID
	case slices.Contains(*vs, v):
		return ErrDuplicateUUID
	}

	*vs = append(*vs, v)
	return nil
}

func (vs *UUIDValues) String() string {
	return strings.Join(*vs, ",")
}

func IsUUID(v string) bool {
	return !strings.ContainsFunc(v, badRune)
}

func badRune(r rune) bool {
	if '0' <= r && r <= '9' || 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' {
		return false
	}
	return true
}
