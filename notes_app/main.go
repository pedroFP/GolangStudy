package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note_app/note"
	"example.com/note_app/todo"
)

type saver interface { // convetion name for interfaces with one method
	Save() error
}

func main() {
	todoText := getUserInput("Todo Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()

	if saveData(todo) != nil {
		return
	}

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	if saveData(userNote) != nil {
		return
	}

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note secceeded!")

	return nil
}


func getNoteData() (string, string) {
	Title := getUserInput("Note Title:")
	Content := getUserInput("Note Content:")
	
	return Title, Content
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
