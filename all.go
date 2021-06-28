package main

import "fmt"

func main () {
	//conversion

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//type data number

	var kkm int = 10
	fmt.Println(kkm)

	//boolean
	var pelajar = true
	fmt.Println(pelajar)

	//perbandingan

	var a = 10
	var b = 10
	fmt.Println(a == b)

	//operasi matematika

	var x = 100
	var y = 100
	var z = x + y
	fmt.Println(z)

	var i = 10
	i++
	fmt.Println(i)

	var j = 20
	j += 20
	fmt.Println(j)

	//variable

	var variab = "this"
	fmt.Println(variab)

	var (
		desa = "tejoasri"
		dusun = "tejo"
	)
	fmt.Println(desa)
	fmt.Println(dusun)

}