package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func cekGender(gender string) string {
	elements := map[string]string{
		"female": "T",
		"akhwat": "T",
		"male":   "N",
		"ikhwan": "N",
	}
	value, exists := elements[strings.ToLower(gender)]
	if !exists {
		fmt.Println(errors.New("gender code unknown"))
		return ""
	}

	return value
}

func generatorNIK(gend string, thn int, numb int) []string {

	var results []string
	gcode := cekGender(gend)
	const prefix = "AR"

	// ambil 2 digit dari tahun
	year := thn % 100
	ycode := strconv.Itoa(year)

	//cek Semester
	bulan := int(time.Now().Month())
	var scode string
	if bulan <= 6 {
		scode = "1"
	} else {
		scode = "2"
	}

	nikpref := prefix + gcode + ycode + scode + "-"
	for i := 1; i <= numb; i++ {
		nik := fmt.Sprintf("%s%05d", nikpref, i)
		results = append(results, nik)
	}

	return results
}

func lanjutNIK(a string, jml int) []string {
	var results []string
	word := strings.Split(a, "-")
	nikPref := word[0]
	nikNum, _ := strconv.Atoi(word[1])
	newnik := nikNum
	for i := 1; i <= jml; i++ {
		newnik++
		element := fmt.Sprintf("%s%s%05d", nikPref, "-", newnik)
		//element := nikPref + "-" + strconv.Itoa(newnik)
		results = append(results, element)
	}
	return results
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Gender:")
	//input, _ := reader.ReadString('\n')
	//fmt.Println("your text:", input)
	//
	//fmt.Println("Input Tahun:")
	//inputTahun, _ := reader.ReadString('\n')
	//inputTahun = strings.TrimSpace(inputTahun)
	//tahun, err := strconv.Atoi(inputTahun)
	//if err != nil {
	//	fmt.Println("Error!", err)
	//	return
	//}
	//fmt.Printf("Tahun anda: %d\n", tahun)
	//
	//fmt.Println("Jumlah yang digenerate:")
	//jmlInput, _ := reader.ReadString('\n')
	//// Remove newline character and convert to int
	//jmlInput = strings.TrimSpace(jmlInput)
	//jml, err := strconv.Atoi(jmlInput)
	//if err != nil {
	//	fmt.Println("Invalid input for jumlah:", err)
	//	return
	//}
	//fmt.Printf("Jumlah: %d\n", jml)
	//
	//now := time.Now()
	//month := now.Month()
	//
	//fmt.Println(month)
	fmt.Println("Tugas 2 Nomor 1 : Generator NIK")
	daftarNIK := generatorNIK("male", 2024, 2)
	for i := 0; i < len(daftarNIK); i++ {
		fmt.Println(daftarNIK[i])
	}

	fmt.Println("Tugas 2 Nomor 2 : lanjutkan NIK")
	daftarBaruNIK := lanjutNIK("ARN192-00334", 5)
	for i := 0; i < len(daftarBaruNIK); i++ {
		fmt.Println(daftarBaruNIK[i])
	}

}
