package main

import "fmt"

func main()  {
	tinggi:=map[string]string{"Aldo": "182 cm", "Yosep": "178 cm"}
	tampil_mahasiswa(tinggi)
}

func tampil_mahasiswa(tinggi map[string]string) {
	for name, height:= range tinggi{
		fmt.Println(name, ":", height)
	}
}