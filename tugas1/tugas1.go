package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa membagi dengan 0")
	}
	return a / b, nil
}

func ToDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Tidak bisa membagi %d dengan 0", a)
	}
	return a / b, nil
}

func testPrint() {
	fmt.Println("Hello")                // muncul output di layar
	out := fmt.Sprintln("Hello SPrint") // tidak ada output di layar pada perintah ini
	fmt.Println(out)
}

func main() {
	// tugas nomor 2
	testPrint()

	//tugas nomor 3
	result, err := divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}

	result, err = ToDivide(8, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}

	//fmt.Println("Hello World")
}
