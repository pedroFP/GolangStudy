package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note secceeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ", prompt)
	
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // on windows machines line breaks are \n\r
	
	return text
}
