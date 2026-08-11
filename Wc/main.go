package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	NoneCommand = 1
)

type Content struct {
	numberLine  int
	numberWords int
	numberBytes int
}

func main() {

	var content Content

	if len(os.Args) == NoneCommand {
		content = parserStreamOfRead()
	} else {
		content = parseArgs()
	}
	fmt.Printf(" Кол. слов: %d \n Кол. строк: %d \n Кол.байт: %d \n", content.numberWords, content.numberLine, content.numberBytes)
}

/*
Анализирует поток ввода
*/
func parserStreamOfRead() Content {

	var builder strings.Builder
	var numberLine int
	var numberBytes int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		builder.WriteString(text)
		builder.WriteString(" ")
		numberLine++
	}

	words := strings.Fields(builder.String())

	for _, word := range words {
		numberBytes += len(word)
	}

	return Content{numberWords: len(words), numberLine: numberLine, numberBytes: numberBytes}
}

/*
Анализирует переданные аргументы в командную строку
*/
func parseArgs() Content {

	content := Content{}

	for _, arg := range os.Args {

		if arg == "-h" {
			fmt.Println(
				"-w 	Отобразить колличество слов в объекте \n" +
					"-l		Вывести колличество строк в объекте\n" +
					"-m		Показать количество символов в объекте\n" +
					"-c		Отобразить размер объекта в байтах\n\n" +
					"Без параметров: Если аргументов нету то подсчитывается все вышепречисленное из свободного ввод в командную строку." +
					"Если указан путь к файлу: подсчитывается все вышепречисленное для файла")

		} else {
			content = readFile(arg)
		}
	}
	return content
}

func readFile(fileName string) Content {
	file, err := os.Open(fileName)
	content := Content{}

	if err != nil {
		log.Fatal(err, "Файл не был найден: "+fileName)
		os.Exit(1)
	}

	defer file.Close()

	scaner := bufio.NewScanner(file)
	strBuilder := strings.Builder{}

	for scaner.Scan() {
		strBuilder.WriteString(scaner.Text())
		strBuilder.WriteString(" ")
		content.numberLine++
		content.numberBytes += len(scaner.Bytes())
	}

	words := strings.Fields(strBuilder.String())

	content.numberWords = len(words)

	return content
}
