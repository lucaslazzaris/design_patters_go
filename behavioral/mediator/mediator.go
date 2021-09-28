package main

import "fmt"

type DialogDirector interface {
	showDialog()
	widgetChanged(Widget) string
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

func (director *FontDialogDirector) widgetChanged(widget Widget) string {
	if widget == director.ok {
		return "Ok"
	} else if widget == director.cancel {
		return "Cancel"
	}
	if widget == director.fontList {
		return "Font List"
	}
	if widget == director.fontName {
		return "Font Name"
	}
	return "Not ok"
}

func (director *FontDialogDirector) createWidgets() {
	director.ok = &Button{director: director}
	director.cancel = &Button{director: director}
	director.fontList = &ListBox{director: director}
	director.fontName = &EntryField{director: director}
}

type Widget interface {
	changed() string
}

type ListBox struct {
	director DialogDirector
}

func (listBox *ListBox) changed() string {
	return listBox.director.widgetChanged(listBox)
}

type EntryField struct {
	director DialogDirector
}

func (entryField *EntryField) changed() string {
	return entryField.director.widgetChanged(entryField)
}

type Button struct {
	director DialogDirector
}

func (button *Button) changed() string {
	return button.director.widgetChanged(button)
}

