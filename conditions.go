package main

import (
	"fmt"
)

func main() {
	fmt.Println(DomainForLocale("code-basics.com", "vn"))
}

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)
}
