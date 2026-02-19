package entities

type Visibility uint8

const (
	VISIBILITY_PUBLIC Visibility = iota
	VISIBILITY_PROTECTED
	VISIBILITY_PRIVATE
)
