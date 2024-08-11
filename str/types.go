package str

import (
	"context"
	"regexp"

	"github.com/insei/valigo/helper"
)

type StringBuilder[T string | *string] interface {
	Trim() StringBuilder[T]
	Required() StringBuilder[T]
	AnyOf(allowed ...string) StringBuilder[T]
	Custom(f func(ctx context.Context, h *helper.Helper, value *T) []error) StringBuilder[T]
	Regexp(regexp *regexp.Regexp, opts ...RegexpOption) StringBuilder[T]
	MaxLen(int) StringBuilder[T]
	MinLen(int) StringBuilder[T]
	When(f func(ctx context.Context, value *T) bool) StringBuilder[T]
}

type StringsBundleBuilder interface {
	String(field *string) StringBuilder[string]
	StringPtr(field **string) StringBuilder[*string]
}