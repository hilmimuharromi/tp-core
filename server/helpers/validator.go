package helpers

import (
	"github.com/google/uuid"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (this *CustomValidator) Init() error {
	err := this.Validator.RegisterValidation("valid_uuid", isValidUUID)
	if err != nil {
		return err
	}

	// Register JSON fields.
	this.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return nil
}

func (this *CustomValidator) Validate(i interface{}) error {
	return this.Validator.Struct(i)
}

func isValidUUID(fl validator.FieldLevel) bool {
	strUuid := fl.Field().String()
	if strUuid == "" {
		return false
	}

	_, err := uuid.Parse(strUuid)
	if err != nil {
		return false
	}

	return true
}
