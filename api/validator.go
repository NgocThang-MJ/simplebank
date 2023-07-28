package api

import (
	"github.com/go-playground/validator/v10"

	"github.com/liquiddev99/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// Check currency is supported
		return util.IsSupportCurrency(currency)
	}

	return false
}
