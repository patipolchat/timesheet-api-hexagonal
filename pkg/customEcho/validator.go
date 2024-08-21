package customEcho

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	validateInstance *validator.Validate
	validateOnce     *sync.Once
)

func GetValidator() *validator.Validate {
	validateOnce.Do(func() {
		validateInstance = validator.New()
	})
	return validateInstance
}
