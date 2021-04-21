package validator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/translations/en"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

type customValidator struct {
	Validator *validator.Validate
}

const (
	FIRST_NAME = "[A-Z][a-zA-Z]*"
	SURNAME = "[A-Z][a-zA-Z]*"
	PASSWORD = `(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z\\d$@$!%*?&].{7,}`

)

const (
	FIRST_NAME_ERROR_MSG = "{0} must be in valid format"
	SURNAME_ERR_MSG = "{0} must be in valid format"
	PASSWORD_ERR_MSG = "Password must have minimum 1 uppercase letter, 1 lowercase letter, 1 digit and 1 special character and needs to be minimum 8 characters long"
)

func NewCustomValidator() *customValidator {
	cv := &customValidator{validator.New()}
	err := registerNameValdation(cv)
	err = registerSurnameValidation(cv)
	err = registerPasswordValidation(cv)
	if err != nil {
		return &customValidator{}
	}
	return cv
}

func (cv *customValidator) RegisterEnTranslation() (ut.Translator, error) {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	registerEnNameTranslation(trans, cv)
	registerEnSurnameTranslation(trans, cv)
	registerEnPasswordTranslation(trans, cv)
	return trans, enTranslations.RegisterDefaultTranslations(cv.Validator, trans)
}

func (cv *customValidator) TranslateError(err error, translator ut.Translator) (translatedErrors []error){
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(translator))
		translatedErrors = append(translatedErrors, translatedErr)
	}
	return translatedErrors
}

func (cv *customValidator) GetErrorsString(errs []error) (errsString []string) {
	for _, e := range errs {
		errsString = append(errsString, e.Error())
	}
	return errsString
}

func registerNameValdation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("name", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(FIRST_NAME, []byte(f1.Field().String()))
		return mathced
	})

}

func registerSurnameValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("surname", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(SURNAME, []byte(f1.Field().String()))
		return mathced
	})
}

func registerPasswordValidation(cv *customValidator) error {
	return cv.Validator.RegisterValidation("password", func(f1 validator.FieldLevel) bool {
		mathced, _ := regexp.Match(PASSWORD, []byte(f1.Field().String()))
		return mathced
	})
}

func registerEnNameTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("name", tr, func(ut ut.Translator) error {
		return ut.Add("name", FIRST_NAME_ERROR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("name", fe.Field())
		return t
	})
}

func registerEnSurnameTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("surname", tr, func(ut ut.Translator) error {
		return ut.Add("surname", SURNAME_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("surname", fe.Field())
		return t
	})
}

func registerEnPasswordTranslation(tr ut.Translator, cv *customValidator) {
	_ = cv.Validator.RegisterTranslation("password", tr, func(ut ut.Translator) error {
		return ut.Add("password", PASSWORD_ERR_MSG, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})
}


