package main

import (
	stringsconverters "Testes/stringsConverters"
	"fmt"
)

func main() {

	str := "402923--4093840 3984924 39482934 3.4.5.5.6.34.34.343.4"

	fmt.Println(stringsconverters.RemoveNonNumber(str))
}
