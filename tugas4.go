package main

import "fmt"

func main() {
	tampil_mahasiswa(182, 178)
}

func tampil_mahasiswa(x int, y int) (int, int) {
	var mahasiswa = map[string]int{"Aldo": x, "Yosep": y}

	for key, val := range mahasiswa {
		fmt.Println(key, ":", val, "cm")
	}
	return x, y
}
