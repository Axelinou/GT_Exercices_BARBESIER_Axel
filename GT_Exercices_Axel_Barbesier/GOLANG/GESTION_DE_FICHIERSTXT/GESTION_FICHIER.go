package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//IMPORTANT !!! Pensez a faire un "cd .\GOLANG\GESTION_DE_FICHIERTXT\" pour que le programme ce lance correctement. Merci
//IMPORTANT N°2 !!! Veuillez lancer le programme via le terminal en tappant "go run .\GESTION_FICHIER.go" . Merci

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(text string, file *os.File) {
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
}

func read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}

func main() {
	file, err := os.OpenFile("text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	check(err)

	write("Je suis en b1 info\n", file)

	data := read(file.Name())
	fmt.Print(data)
}
