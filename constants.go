// В больших проектах часто используется gRPC — высокопроизводительный RPC (Remote Procedure Call)-фреймворк для коммуникации между микросервисами. Ошибки в gRPC стандартизированы и представлены в виде строк для пользователя, либо в виде чисел для компьютера.

// Представим, что нам нужно написать конвертер ошибок в числовой формат для gRPC. Реализуйте функцию ErrorMessageToCode(msg string) int, которая возвращает числовой код для заданного значения. Список сообщений и соответствующих кодов:

// OK = 0
// CANCELLED = 1
// UNKNOWN = 2

// В реальности список намного шире. Мы для простоты ограничимся тремя ошибками. Учтите, что если в функцию передать неизвестную ошибку, она должна вернуть код ошибки для сообщения UNKNOWN.

package main

import (
	"fmt"
)

const (
	OK        = 0
	CANCELLED = 1
	UNKNOWN   = 2
)
const (
	OKmsg        = "OK"
	CANCELLEDmsg = "CANCELLED"
	UNKNOWNmsg   = "UNKNOWN"
)

func main() {
	ErrorMessageToCode("CANCELLED")
}

func ErrorMessageToCode(msg string) int {
	switch msg {
	case OKmsg:
		return OK
	case CANCELLEDmsg:
		return CANCELLED
	default:
		return UNKNOWN
	}
}
