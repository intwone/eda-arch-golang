package err

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	transl "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		en := en.New()
		unicode := ut.New(en, en)
		translator, _ = unicode.GetTranslator("en")
		transl.RegisterDefaultTranslations(val, translator)
	}
}

func ErrorValidation(validationError error) *RestError {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		causes := []Cause{
			{Field: "", Message: "invalid field type"},
		}

		return NewBadRequestError("invalid field type", causes)
	}

	if errors.As(validationError, &jsonValidationError) {
		causes := []Cause{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := Cause{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}

			causes = append(causes, cause)
		}

		return NewBadRequestError("invalid fields", causes)
	}

	causes := []Cause{
		{Field: "", Message: "convert fields error"},
	}

	return NewBadRequestError("invalid fields", causes)
}
