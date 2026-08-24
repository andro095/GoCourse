package main

import "fmt"

func string_bytes() {
	// var my_string string = "hello world"

	// var my_byte byte = my_string[0]

	// fmt.Println(my_byte)

	// var s2 string = my_string[4:8]
	// var s3 string = my_string[:8]

	// fmt.Println(s2)
	// fmt.Println(s3)
	// var s3 string = my_string[:8]

	// fmt.Println(s2)
	// fmt.Println(s3)

	// strings to slices
	var s string = "Hello world"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)

	fmt.Println(bs)
	fmt.Println(rs)

}
