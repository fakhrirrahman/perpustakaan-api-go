package dto

type Respone[T any] struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success"`
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

// Specific response types for Swagger documentation
type ResponeAuthResponse struct {
	Code    int          `json:"code" example:"200"`
	Message string       `json:"message" example:"Success"`
	Data    AuthResponse `json:"data"`
}

type ResponeBookData struct {
	Code    int      `json:"code" example:"200"`
	Message string   `json:"message" example:"Success"`
	Data    BookData `json:"data"`
}

type ResponeCustomerData struct {
	Code    int          `json:"code" example:"200"`
	Message string       `json:"message" example:"Success"`
	Data    CustomerData `json:"data"`
}

type ResponeArrayBookData struct {
	Code    int        `json:"code" example:"200"`
	Message string     `json:"message" example:"Success"`
	Data    []BookData `json:"data"`
}

type ResponeArrayCustomerData struct {
	Code    int            `json:"code" example:"200"`
	Message string         `json:"message" example:"Success"`
	Data    []CustomerData `json:"data"`
}

type ResponeString struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success"`
	Data    string `json:"data"`
}