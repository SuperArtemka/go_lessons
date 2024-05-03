package main

import (
	"fmt"
	"log"
	"os"
)

func splitFileNameAndExt(path string) (string, string) {
	var fileName, fileExt string
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			fileName = path[:i]
			fileExt = path[i+1:]
			return fileName, fileExt
		}
	}
	return path, "" // Если нет расширения
}
func main() {
	filePth := ""
	fmt.Println("Введите путь:")
	fmt.Scan(&filePth)
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth = os.Args[1]

	fileName, fileExt := splitFileNameAndExt(filePth)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
