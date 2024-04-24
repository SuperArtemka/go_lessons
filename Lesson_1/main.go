package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	_ "unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordLine, symbolNumber, bytesNumber int
	// Напишите код здесь
	bytesNumber = len(text)

	words := strings.Fields(text)
	wordLine = len(words)
	for _, s := range words {
		symbolNumber += len([]rune(s))
	}
	// Подсказка. Чтобы обойти все символы в строке используйте цикл for _, v := range value
	// Помните вам нужно вывести количество букв, а не символов.
	// Используйте unicode.IsSpace() для определения пробела
	// Используйте unicode.IsLetter() для определения буквы
	fmt.Printf("Количество слов: %d\n", wordLine)
	fmt.Printf("Количество букв: %d\n", symbolNumber)
	fmt.Printf("Количество байт: %d\n", bytesNumber)
}
