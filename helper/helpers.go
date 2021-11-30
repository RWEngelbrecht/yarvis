package helper

import (
	"errors"
)

func OutputError(errText string) {
	panic(errors.New(errText).Error())
}
