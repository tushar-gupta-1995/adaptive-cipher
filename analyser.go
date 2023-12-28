package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("sample_text\\chatGPTStory.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	cipherMap := make(map[string][]int)
	// Read the file character by character.
	var count int = 0
	totalText := ""
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break // End of file
		}
		s := string(char)
		if s == " " {
			continue
		}
		totalText = totalText + s

		cipherMap[s] = append(cipherMap[s], count)
		// fmt.Print(s + " - " + strconv.Itoa(count))

		count = count + 1

		// fmt.Print(string(char))
	}

	encrypt := "this is text to encrypt don"
	re := regexp.MustCompile(`\s+`)

	encrypt = re.ReplaceAllString(encrypt, "")

	encryptedText := ""
	for _, char := range encrypt {
		l := cipherMap[string(char)]
		randomIndex := rand.Intn(len(l))
		randomValue := l[randomIndex]
		// fmt.Print(strconv.Itoa(randomValue) + " - " + string(char))
		encryptedText = encryptedText + "ab" + strconv.Itoa(randomValue)
	}

	// fmt.Print(encryptedText)

	d := strings.Split(encryptedText, "ab")

	for index, j := range d {
		if index == 0 {
			continue
		}
		i, _ := strconv.Atoi(j)
		// fmt.Print(i)
		// fmt.Println()
		for count, text := range totalText {
			if count == i {
				fmt.Print(string(text) + " ")
				break
			}
		}
	}

}
