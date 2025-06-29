package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T)map[string]string{
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TranslateTag(v)
		}
	}
	return res
}

func TranslateTag(fd validator.FieldError) string {
	switch fd.Tag() {
	case "required":
		return fmt.Sprintf("field %s wajib diisi", fd.StructField())
	case "min":
		return fmt.Sprintf("field %s minimal %s karakter", fd.StructField(), fd.Param())
	case "max":
		return fmt.Sprintf("field %s maksimal %s karakter", fd.StructField(), fd.Param())
	case "email":
		return fmt.Sprintf("field %s harus berformat email yang valid", fd.StructField())
	case "uuid":
		return fmt.Sprintf("field %s harus berformat UUID yang valid", fd.StructField())
	case "unique":
		return fmt.Sprintf("field %s harus unik", fd.StructField())
	}
	return "validasi gagal"
}