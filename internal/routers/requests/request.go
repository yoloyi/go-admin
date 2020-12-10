package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go-admin/internal/utils/response"
	re "go-admin/internal/utils/response/e"
	"net/http"
	"reflect"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	localTranslator := zh.New()
	uni = ut.New(localTranslator, localTranslator)
	trans, _ = uni.GetTranslator("localTranslator")
	validate = validator.New()
	// 替换报错名称
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Validate(ctx *gin.Context, s interface{}) bool {
	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response.NewResponse(ctx).ErrorWithMessage(http.StatusOK, re.ParamsErrorCode, e.Translate(trans), nil)
			return false
		}
	}
	return true
}
