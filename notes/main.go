package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alopes2/go-studies/notes/note"
	"github.com/alopes2/go-studies/notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		panic(err)
	}

	err = outputData(todo)

	if err != nil {
		panic(err)
	}

	userNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	err = outputData(userNote)

	if err != nil {
		panic(err)
	}
}

func outputData(data outputtable) error {
	data.Display()

	return saveData(data)
}

func printSomething(value interface{}) {
	typedVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", typedVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float64: ", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println(stringVal)
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving data failed")
		return err
	}

	fmt.Println("Saving data succeeded")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
