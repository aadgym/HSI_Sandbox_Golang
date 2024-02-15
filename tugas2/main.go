package main

import (
	"errors"
	"fmt"
	"sort"
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

func generateStudentGroups(taAdmin, tahunAjaran []int) map[int][]string {
	mapAdminIkhawan := make(map[int][]string)
	mapSantriIkhwan := make(map[int][]string)
	mapKelompok := make(map[int][]string)

	for _, thnAdm := range taAdmin {
		listNIkadmin := generatorNIK("male", thnAdm, 1)
		mapAdminIkhawan[thnAdm] = listNIkadmin
	}

	for _, thnA := range tahunAjaran {
		listNIKIkh := generatorNIK("male", thnA, 5)
		nikAdm := mapAdminIkhawan[thnA-1]
		mapSantriIkhwan[thnA] = listNIKIkh
		mapSantriIkhwan[thnA] = append(mapSantriIkhwan[thnA], nikAdm...)
		sort.Strings(mapSantriIkhwan[thnA])
		mapKelompok[thnA] = mapSantriIkhwan[thnA]
	}

	return mapKelompok
}

func main() {
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

	//
	//
	fmt.Println("Tugas 2 Nomor 3 ")

	tahunAjaran := [6]int{2019, 2020, 2021, 2022, 2023, 2024}
	taAdmin := [6]int{2018, 2019, 2020, 2021, 2022, 2023}

	allKelompok := generateStudentGroups(taAdmin[:], tahunAjaran[:])
	fmt.Println(allKelompok)
}
