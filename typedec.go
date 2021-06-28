package main

import "fmt"

func main () {
	type ktp string
	type married bool

	var nomor ktp = "1212"
	var sudah married = true

	fmt.Println(nomor)
	fmt.Println(sudah)
}
