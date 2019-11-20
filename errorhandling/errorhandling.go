package errorhandling

import "fmt"

// KindError error with kind
type KindError struct {
	Kind   string
	Reason string
}

// DetailError error with detail
type DetailError struct {
	KindError
	Detail string
}

func (m KindError) Error() string {
	msg := fmt.Sprintf("[%s] %s", m.Kind, m.Reason)
	return msg
}

func (d DetailError) Error() string {
	msg := d.KindError.Error()
	msg = fmt.Sprintf("%s detail: %s ", msg, d.Detail)
	return msg
}

// AKindError a kind error
func AKindError() error {
	err := &KindError{
		Kind:   "testing",
		Reason: "mock error",
	}
	return err
}

func ADetailError() error {
	err := &DetailError{
		KindError: KindError{
			Kind:   "testing",
			Reason: "mock detailed error",
		},
		Detail: "detail error type using nesting type",
	}
	return err
}
