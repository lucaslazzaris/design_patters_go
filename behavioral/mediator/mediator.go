package main

import "fmt"

type DialogDirector interface {
	showDialog()
	widgetChanged(Widget)
	createWidgets()
}

type FontDialogDirector struct{
	ok *Button
	cancel *Button
	fontList *ListBox
	fontName *EntryField
}

func (director *FontDialogDirector) showDialog() {
	fmt.Println("Dialog")
}

func (director *FontDialogDirector) widgetChanged(widget Widget) {
	if widget == director.ok {
		fmt.Println("Ok")
	} else if widget == director.cancel {
		fmt.Println("Cancel")
	}
	if widget == director.fontList {
		fmt.Println("Font List")
	}
	if widget == director.fontName {
		fmt.Println("Font Name")
	}
}

func (director *FontDialogDirector) createWidgets() {
	director.ok = &Button{director: director}
	director.cancel = &Button{director: director}
	director.fontList = &ListBox{director: director}
	director.fontName = &EntryField{director: director}
}

type Widget interface {
	changed()
}

type ListBox struct {
	director DialogDirector
}

func (listBox *ListBox) changed() {
	listBox.director.widgetChanged(listBox)
}

type EntryField struct {
	director DialogDirector
}

func (entryField *EntryField) changed() {
	entryField.director.widgetChanged(entryField)
}

type Button struct {
	director DialogDirector
}

func (button *Button) changed() {
	button.director.widgetChanged(button)
}

func main() {
	fmt.Println("Hello world")
}
