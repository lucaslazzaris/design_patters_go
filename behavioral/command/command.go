package main

import "fmt"

type Command interface {
	execute()
}

type Application struct {
	documents []*Document
}

func (application *Application) add(document *Document) {
	application.documents = append(application.documents, document)
}

type Document struct {
	name string
}

func (document *Document) open() {
	fmt.Println("Doc", document.name, "opened")
}

func (document *Document) paste() {
	fmt.Println("Doc", document.name, "pasted data from clipboard")
}

type OpenCommand struct {
	application *Application
	response string
}

func newOpenCommand(application *Application) OpenCommand {
	return OpenCommand{application: application}
}

func (openCommand *OpenCommand) execute() {
	name := openCommand.askUser()

	if (name != "" ){
		document := &Document{name: name}
		openCommand.application.add(document)
		document.open()
	}
}

func (openCommand *OpenCommand) askUser() string {
	return "UserData.pdf"
}

type PasteCommand struct {
	document *Document
}

func newPasteCommand(document *Document) PasteCommand {
	return PasteCommand{document: document}
}

func (pasteCommand PasteCommand) execute() {
	pasteCommand.document.paste()
}

type MacroCommand struct {
	commands []Command
}

func (macroCommand *MacroCommand) add(command Command) {
	macroCommand.commands = append(macroCommand.commands, command)
}

func (macroCommand *MacroCommand) remove(index int) {
	remainingCommands := make([]Command, 0)
	remainingCommands = append(remainingCommands, macroCommand.commands[:index]...)
	macroCommand.commands = append(remainingCommands, macroCommand.commands[index+1:]...)
}

func (macroCommand *MacroCommand) execute() {
	size := len(macroCommand.commands)
	for i := 0; i < size; i++ {
		macroCommand.commands[i].execute()
	}

}
