package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	NoneCommand = 1
)

type Content struct {
	numberLine  int
	numberWords int
	numberBytes int
	numberChars int
}

func main() {

	var content Content

	if len(os.Args) == NoneCommand {
		content = parserStreamOfRead()
	} else {
		content = parseArgs()
	}

	if content.numberChars != 0 {
		fmt.Printf("Кол. сим: %d \n", content.numberChars)
	}

	if content.numberWords != 0 {
		fmt.Printf("Кол. слов: %d \n", content.numberWords)
	}

	if content.numberLine != 0 {
		fmt.Printf("Кол. линий: %d \n", content.numberLine)
	}

	if content.numberBytes != 0 {
		fmt.Printf("Кол. байтов: %d \n", content.numberBytes)
	}

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

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(
			"-w 	Отобразить колличество слов в объекте \n" +
				"-l	Вывести колличество строк в объекте\n" +
				"-m	Показать количество символов в объекте\n" +
				"-c	Отобразить размер объекта в байтах\n\n" +
				"Без параметров: Если аргументов нету то подсчитывается все вышепречисленное из свободного ввод в командную строку." +
				"Если указан путь к файлу: подсчитывается все вышепречисленное для файла")
		return Content{}
	}

	file := OpenFile(os.Args[len(os.Args)-1])

	defer file.Close()

	for _, arg := range os.Args[1:] {

		if arg == "-l" {
			content.numberLine = CalculateLineOfFile(file)
			break
		} else if arg == "-w" {
			content.numberWords = CalculateWordsOfFile(file)
			break
		} else if arg == "-m" {
			content.numberChars = CalculateCountCharecters(file)
			break
		} else if arg == "-c" {
			content.numberBytes = CalculateFileOfBytes(file)
			break
		} else {
			content = AnalyzingAllFile(file)
			break
		}
	}
	return content
}

func AnalyzingAllFile(file *os.File) Content {

	scaner := bufio.NewScanner(file)

	lineNumber := 0
	wordsNumber := 0
	bytesNumber := CalculateFileOfBytes(file)

	for scaner.Scan() {
		lineNumber++
		wordsNumber += len(strings.Fields(scaner.Text()))
	}

	return Content{
		numberLine:  lineNumber,
		numberWords: wordsNumber,
		numberBytes: bytesNumber,
	}
}

func OpenFile(fileName string) *os.File {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err, "Файл не был найден: "+fileName)
	}

	return file
}

/*
Вычисление байтов в файле
*/
func CalculateFileOfBytes(file *os.File) int {
	info, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	return int(info.Size())
}

/*
Вычислить слова из файла
*/
func CalculateWordsOfFile(file *os.File) int {
	scaner := bufio.NewScanner(file)
	wordsNumber := 0

	for scaner.Scan() {
		wordsNumber += len(strings.Fields(scaner.Text()))
	}

	return wordsNumber
}

/*
Вычислить количество строк
*/
func CalculateLineOfFile(file *os.File) int {
	data, err := os.ReadFile(file.Name())

	if err != nil {
		log.Fatal(err)
	}

	text := string(data)
	lineNumber := strings.Count(string(data), "\n")

	if len(text) > 0 && text[len(text)-1] != '\n' {
		lineNumber++
	}

	return lineNumber
}

/*
Подсчёт символов из файла
*/
func CalculateCountCharecters(file *os.File) int {
	data, err := os.ReadFile(file.Name())

	if err != nil {
		log.Fatal(err)
	}

	text := string(data)
	return utf8.RuneCountInString(text)
}
