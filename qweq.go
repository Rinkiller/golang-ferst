package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	indexOfslash := strings.LastIndex(filePth, "/")
	indexOfpoint := strings.LastIndex(filePth, ".")
	if indexOfslash == -1 {
		fileName = filePth
	} else {
		if indexOfpoint == -1 {
			fileExt = ""
			fileName = filePth[indexOfslash:]
		} else if indexOfpoint < indexOfslash {
			fmt.Println("Ошибочный вводные данные! Код приостановлен!!!")
			os.Exit(1)
		} else {
			fileName = filePth[indexOfslash+1 : indexOfpoint]
			fileExt = filePth[indexOfpoint+1:]
		}
	}
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
