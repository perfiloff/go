// На сервер приходит HTTP-запрос. Тело запроса парсится и мапится в модель.
// Сразу работать с моделью нельзя, потому что данные могут быть неверными.
// Реализуйте функцию Validate(req UserCreateRequest) string, которая валидирует запрос и возвращает строку с ошибкой "invalid request", если модель невалидна.
// Если запрос корректный, то функция возвращает пустую строку. Правила валидации и структура модели описаны ниже.
// Не используйте готовые библиотеки и опишите правила самостоятельно.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "aa",
		Age:       1,
	}))
}

type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

func Validate(req UserCreateRequest) string {
	switch {
	case ValFirstName(req.FirstName) == false:
		return "invalid request"
	case ValAge(req.Age) == false:
		return "invalid request"
	default:
		return ""
	}
}

func ValFirstName(name string) bool {
	if name == "" {
		return false
	} else if strings.Contains(name, " ") {
		return false
	} else {
		return true
	}
}

func ValAge(age int) bool {
	if 150 >= age && age > 0 {
		return true
	} else {
		return false
	}
}
