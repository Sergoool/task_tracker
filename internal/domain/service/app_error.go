package service

type Code string
const (
	CodeValidation Code = "validation_error" // неверные данные
	CodeNotFound   Code = "not_found"        // не найдено
	CodeInternal   Code = "internal_error"   // внутренняя ошибка
)

type AppError struct {
	Code    Code
	Details any
	err     error
}

func (e *AppError) Error() string {
	if e.err != nil {
		return string(e.Code) + ": " + e.err.Error()
	}
	return string(e.Code)
}

func (e *AppError) Unwrap() error { return e.err }

// helpers
func Validation(details any) error { return &AppError{Code: CodeValidation, Details: details} } // создать validation
func NotFound(details any) error   { return &AppError{Code: CodeNotFound, Details: details} }  // создать not_found
func Internal(err error) error     { return &AppError{Code: CodeInternal, err: err} }          // создать internal
