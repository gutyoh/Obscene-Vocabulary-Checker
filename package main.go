package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// go program that will read a file and print the contents to the console
	// the file name will be input by the user and is in the same directory as the program

	in := bufio.NewReader(os.Stdin)

	// ask user to input file name

	var fname string
	fmt.Scanln(&fname)

	// create while loop for the program
	// the program will continue to ask for the badWord until the user enters exit
	// if the user enters exit the program will print "Bye!" and exit

	// ask the user to input a sentence to check for bad words
	badSentence, _ := in.ReadString('\n')

	for badSentence != "exit" + "\n" {

		// open the file
		file, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		var words []string

		// create a scanner to read the file
		scanner := bufio.NewScanner(file)

		// read the file
		for scanner.Scan() {
			// split the words into a list
			words = append(words, strings.Split(strings.ToLower(scanner.Text()), " ")...)
		}

		// Create words list with "disgusting", "unpleasant", "ugly", "bad"
		// words = append(words, "disgusting", "unpleasant", "ugly", "bad")

		// Split badSentence into a list and remove the new line character at the end
		var badWordList []string
		badWordList = strings.Split(badSentence, " ")
		// remove the "\n" character at the last element of badWordList
		badWordList[len(badWordList)-1] = badWordList[len(badWordList)-1][:len(badWordList[len(badWordList)-1])-1]

		// look for word in badWordList
		// if word is in badWordList replace it with asterisks equal to len(word)
		// then update badWordList and print the updated badWordList
		for i := 0; i < len(badWordList); i++ {
			for j := 0; j < len(words); j++ {
				if strings.ToLower(badWordList[i]) == words[j] {
					badWordList[i] = strings.Repeat("*", len(badWordList[i]))
				}
			}
		}

		// join the badWordList into a string
		badWordString := strings.Join(badWordList, " ")
		fmt.Println(badWordString)
		// ask the user to input a sentence to check for bad words
		badSentence, _ = in.ReadString('\n')
	}
	// print "Bye!"
	fmt.Println("Bye!")
}
