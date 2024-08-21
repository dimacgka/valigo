package num

import (
	"context"
	"reflect"

	"github.com/insei/fmap/v3"

	"github.com/insei/valigo/shared"
)

type intSliceBuilder[T []int | *[]int | []int8 | *[]int8 | []int16 | *[]int16 | []int32 | *[]int32 | []int64 | *[]int64] struct {
	h        shared.Helper
	field    fmap.Field
	appendFn func(field fmap.Field, fn shared.FieldValidationFn)
}

// Min checks if each integer in the slice has a minimum number.
func (s *intSliceBuilder[T]) Min(minNum int) IntSliceBuilder[T] {
	s.appendFn(s.field, func(ctx context.Context, h shared.Helper, value any) []shared.Error {
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if int(v.Index(i).Int()) < minNum {
					return []shared.Error{h.ErrorT(ctx, s.field, v.Index(i).Int(), minLocaleKey, minNum)}
				}
			}
		default:
			return []shared.Error{h.ErrorT(ctx, s.field, "", invalidLocaleKey)}
		}
		return nil
	})
	return s
}

// Max checks if each integer in the slice has a maximum number.
func (s *intSliceBuilder[T]) Max(maxNum int) IntSliceBuilder[T] {
	s.appendFn(s.field, func(ctx context.Context, h shared.Helper, value any) []shared.Error {
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if int(v.Index(i).Int()) > maxNum {
					return []shared.Error{h.ErrorT(ctx, s.field, v.Index(i).Int(), maxLocaleKey, maxNum)}
				}
			}
		default:
			return []shared.Error{h.ErrorT(ctx, s.field, "", invalidLocaleKey)}
		}
		return nil
	})
	return s
}

// Required checks if the integer slice is not empty.
func (s *intSliceBuilder[T]) Required() IntSliceBuilder[T] {
	s.appendFn(s.field, func(ctx context.Context, h shared.Helper, value any) []shared.Error {
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			if v.Len() < 1 {
				return []shared.Error{h.ErrorT(ctx, s.field, v, requiredLocaleKey)}
			}
		default:
			return []shared.Error{h.ErrorT(ctx, s.field, "", invalidLocaleKey)}
		}
		return nil
	})
	return s
}

// Custom allows for custom validation logic.
func (s *intSliceBuilder[T]) Custom(f func(ctx context.Context, h *shared.FieldCustomHelper, value *T) []shared.Error) IntSliceBuilder[T] {
	customHelper := shared.NewFieldCustomHelper(s.field, s.h)
	s.appendFn(s.field, func(ctx context.Context, h shared.Helper, value any) []shared.Error {
		return f(ctx, customHelper, value.(*T))
	})
	return s
}

// When allows for conditional validation based on a given condition.
func (s *intSliceBuilder[T]) When(whenFn func(ctx context.Context, value *T) bool) IntSliceBuilder[T] {
	if whenFn == nil {
		return s
	}
	s.appendFn = func(field fmap.Field, fn shared.FieldValidationFn) {
		fnWithEnabler := func(ctx context.Context, h shared.Helper, v any) []shared.Error {
			if !whenFn(ctx, v.(*T)) {
				return nil
			}
			return fn(ctx, h, v)
		}
		s.appendFn(field, fnWithEnabler)
	}
	return s
}