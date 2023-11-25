// Для выполнения этого задания потребуется функция json.Unmarshal, которая декодирует JSON байты в структуру:

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type HelloWorld struct {
//     Hello string
// }

// func main() {
//     hw := HelloWorld{}

//   // первым аргументом передается JSON-строка в виде слайса байт. Вторым аргументом указатель на структуру, в которую нужно декодировать результат.
//     err := json.Unmarshal([]byte("{\"hello\":\"world\"}"), &hw)

//     fmt.Printf("error: %s, struct: %+v\n", err, hw) // error: %!s(<nil>), struct: {Hello:world}
// }

// В API методах часто используются запросы с телом в виде JSON. Такие тела нужно декодировать в структуры и валидировать. Хоть это и не лучшая практика делать функции, в которых происходит несколько действий, но для простоты примера реализуйте функцию DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error), которая декодирует тело запроса из JSON в структуру CreateUserRequest и валидирует ее. Если приходит невалидный JSON или структура заполнена неверно, функция должна вернуть ошибку.
// Структура запроса:

// type CreateUserRequest struct {
//     Email                string `json:"email"`
//     Password             string `json:"password"`
//     PasswordConfirmation string `json:"password_confirmation"`
// }

// Список ошибок, которые нужно возвращать из функции:

// // validation errors
// var (
//     errEmailRequired                = errors.New("email is required") // когда поле email не заполнено
//     errPasswordRequired             = errors.New("password is required") // когда поле password не заполнено
//     errPasswordConfirmationRequired = errors.New("password confirmation is required") // когда поле password_confirmation не заполнено
//     errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation") // когда поля password и password_confirmation не совпадают
// )

package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type Validator interface {
	Validate()
}

var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func (cur *CreateUserRequest) Validate() (CreateUserRequest, error) {
	if cur.Email == "" {
		return CreateUserRequest{}, errEmailRequired
	} else if cur.Password == "" {
		return CreateUserRequest{}, errPasswordRequired
	} else if cur.PasswordConfirmation == "" {
		return CreateUserRequest{}, errPasswordConfirmationRequired
	} else if cur.Password != cur.PasswordConfirmation {
		return CreateUserRequest{}, errPasswordDoesNotMatch
	} else {
		return *cur, nil
	}
}

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	// CreateUserRequest.Validate()
	cur := CreateUserRequest{}
	json.Unmarshal(requestBody, &cur)
	res, err := cur.Validate()
	// fmt.Println(res, err)
	return res, err
}

func main() {
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "email is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "password is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}")))                               // CreateUserRequest{}, "password confirmation is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")))                          // CreateUserRequest{}, "password does not match with the confirmation"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}"))) // CreateUserRequest{Email:"email@test.com", Password:"passwordtest", PasswordConfirmation:"passwordtest"}, nil
}

// func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
// 	req := CreateUserRequest{}

// 	err := json.Unmarshal(requestBody, &req)
// 	if err != nil {
// 		return CreateUserRequest{}, err
// 	}

// 	err = validateCreateUserRequest(req)
// 	if err != nil {
// 		return CreateUserRequest{}, err
// 	}

// 	return req, nil
// }

// func validateCreateUserRequest(req CreateUserRequest) error {
// 	if req.Email == "" {
// 		return errEmailRequired
// 	}

// 	if req.Password == "" {
// 		return errPasswordRequired
// 	}

// 	if req.PasswordConfirmation == "" {
// 		return errPasswordConfirmationRequired
// 	}

// 	if req.Password != req.PasswordConfirmation {
// 		return errPasswordDoesNotMatch
// 	}

// 	return nil
// }
