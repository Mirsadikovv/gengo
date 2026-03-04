package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fobus1289/ufa_shared/gengo/internal"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln(errors.New("usage: gengo --new | --add"))
	}

	switch os.Args[1] {
	case "--new":
		projectName := promptInput("Enter project name: ")
		moduleName := promptInput("Enter module name:  ")
		entityName := promptInput("Enter entity name:  ")
		modPath := promptInput("Enter mod path:     ")
		internal.NewProject(projectName, moduleName, entityName, modPath)

	case "--add":
		moduleName := promptInput("Enter module name: ")
		entityName := promptInput("Enter entity name: ")
		internal.AddModule(moduleName, entityName)

	default:
		log.Fatalln(fmt.Errorf("unknown flag %q, use --new or --add", os.Args[1]))
	}
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
