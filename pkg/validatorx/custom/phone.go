package custom

import (
	"fgzs-single/pkg/util/validutil"
	"github.com/go-playground/validator/v10"
)

type Phone struct {
}

func NewPhone() *Phone {
	return &Phone{}
}

func (p *Phone) Tag() string {
	return "phone"
}

func (p *Phone) ZhTranslation() string {
	return "{0} 错误的手机格式"

}

func (p *Phone) EnTranslation() string {
	return "{0} wrong mobile phone format"
}

func (p *Phone) Validate(fl validator.FieldLevel) bool {
	return validutil.IsPhone(fl.Field().String())
}
