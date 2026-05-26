package validate

import (
	"knative.dev/pkg/apis"
)

func FieldNotEmpty(value, field string) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func ListNotEmpty(value []string, field string) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func ImmutableField(original, current interface{}, field string) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func Tag(value string) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func Tags(tags []string, fieldName string) *apis.FieldError { _ = "STUB: not implemented"; return nil }

//noinspection GoNilness

func Image(value string) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func StripComponents(value int64) *apis.FieldError { _ = "STUB: not implemented"; return nil }
