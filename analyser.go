package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cipher struct {
	totalText     string
	encryptedText string
	cipherMap     map[string][]int
	delimmiter    string
}

const delimmiter = "ab"

func main() {

	var textPath string
	flag.StringVar(&textPath, "textPath", "default", "path to the file to encrypt text")
	defaultText := "apple A ball cat dog ear flag game high it join king land mine noon operation purchase queen rusty vaseline zebra win xmas yaarana B J D H I S C E E F B G H I J K L M N TTTTTTTTTT O P Q R W X Y Z ZZZZZZ BBBBBBBBB EEEEEEEEEEEEEEEEEE"

	var encrypt string
	flag.StringVar(&encrypt, "e", "", "text to encrypt")

	var decrypt string
	flag.StringVar(&decrypt, "d", "", "text to decrypt")

	h := flag.Bool("help", false, "show help")

	flag.Parse()

	if *h {
		help()
		os.Exit(0)
	}

	if encrypt == "" && decrypt == "" {
		fmt.Println("Error: Text to encrypt is not provided, this is an encryption tool, not much to be done without something to decrypt")
		os.Exit(1)
	}

	if encrypt != "" && decrypt != "" {
		fmt.Println("Error: encrypt and decrypt at the same time?..are you sure?")
		os.Exit(1)
	}

	var c cipher
	if textPath == "default" {
		fmt.Println("Using default text to encrypt/decrypt")
		c = createCipherFromDefault(defaultText)
	} else {
		c = extractTotalText("sample_text\\random.txt")
	}

	if decrypt != "" {
		c.encryptedText = decrypt
		c.delimmiter = delimmiter
		fmt.Println(c.decrypt())
		os.Exit(0)
	}

	re := regexp.MustCompile(`\s+`)

	encrypt = re.ReplaceAllString(encrypt, "")

	encryptedText := ""
	for _, char := range encrypt {
		l := c.cipherMap[string(char)]
		randomIndex := rand.Intn(len(l))
		randomValue := l[randomIndex]
		encryptedText = encryptedText + delimmiter + strconv.Itoa(randomValue)
	}

	fmt.Println(encryptedText)

	// fmt.Print(c.decrypt())

}

func (c cipher) decrypt() string {
	d := strings.Split(c.encryptedText, c.delimmiter)

	decryptedText := ""
	for index, j := range d {
		if index == 0 {
			continue
		}
		i, _ := strconv.Atoi(j)
		for count, text := range c.totalText {
			if count == i {
				decryptedText = decryptedText + string(text) + " "
				break
			}
		}
	}

	return decryptedText
}

func extractTotalText(path string) cipher {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	cipherMap := make(map[string][]int)
	var count int = 0
	totalText := ""
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		s := string(char)
		if s == " " {
			continue
		}
		totalText = totalText + s

		cipherMap[s] = append(cipherMap[s], count)

		count = count + 1
	}

	return cipher{
		totalText: totalText,
		cipherMap: cipherMap,
	}
}

func createCipherFromDefault(text string) cipher {
	cipherMap := make(map[string][]int)
	for count, text := range text {
		s := string(text)
		if s == " " {
			continue
		}
		cipherMap[s] = append(cipherMap[s], count)
	}
	return cipher{
		totalText: text,
		cipherMap: cipherMap,
	}
}

func help() {
	fmt.Println("book-cipher is a cli tool that uses the old fashioned book to encrypt plan text, you can provide a path to text file or not provide it and the tool will use the default text to encrypt")

	fmt.Println("Usage to encrypt: .\\adaptive-cipher.exe -e=\"<your text>\"")

	fmt.Println("Usage to decrypt: .\\adaptive-cipher.exe -d=\"<your decrypted text>\"")

	fmt.Println("Note: if you are using a custom file path to encrypt make sure to provide the same for decrypting as well")
}
