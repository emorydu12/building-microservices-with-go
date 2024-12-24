package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestErrorWhenRequestEmailNotPresent(t *testing.T) {
	validate := validator.New()
	req := Request{
		URL: "https://emorydu.com",
	}

	if err := validate.Struct(&req); err == nil {
		t.Error("Should have raised an error")
	}
}

func TestErrorWhenRequestEmailIsInvalid(t *testing.T) {
	validate := validator.New()
	req := Request{
		Email: "something.com",
		URL:   "https://emorydu.com",
	}

	if err := validate.Struct(&req); err == nil {
		t.Error("Should have raised an error")
	}
}

func TestNoErrorWhenRequestNameNotPresent(t *testing.T) {
	validate := validator.New()
	req := Request{
		Email: "test@emorydu.com",
		URL:   "https://emorydu.com",
	}

	if err := validate.Struct(&req); err != nil {
		t.Error(err)
	}
}
