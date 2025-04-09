package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// sec 6: les 101
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// sec 6: les 103
//
//	type outputtable interface {
//		Save() error
//		Display()
//	}
type outputtable interface {
	saver
	Display()
}

func main() {
	printSometing(1)
	printSometing(1.5)
	printSometing("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

// sec 6: les 104
// this func doesnt need in this programm, its only demonstration to "Any Value Allowed" Type feature
func printSometing(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Int:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println(stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Int:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
