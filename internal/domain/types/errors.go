package types

import "errors"

var ErrNotFound = errors.New("not found")
var ErrValidation = errors.New("rows is empty or your id is equal zero")
