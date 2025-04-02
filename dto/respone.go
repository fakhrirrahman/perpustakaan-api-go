package dto

type Respone[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponeError(message string) Respone[string] {
	return Respone[string]{
		Code:    99,
		Message: message,
		Data:    "",
	}
}

func CreateResponeErrorData(message string, data map[string]string) Respone[map[string]string] {
	return Respone[map[string]string]{
		Code:    99,
		Message: message,
		Data:    data,
	}
}

func CreateResponeSuccess[T any](data T) Respone[T] {
	return Respone[T]{
		Code:    00,
		Message: "Success",
		Data:    data,
	}
}