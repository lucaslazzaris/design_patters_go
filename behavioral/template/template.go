package main

type ViewTemplate interface {
	doDisplay() string
}

type View struct {
	template ViewTemplate
}

func (view *View) display() string {
	displayText := ""
	displayText += view.setFocus()
	displayText += view.template.doDisplay()
	displayText += view.resetFocus()
	return displayText
}

func (view *View) setFocus() string {
	return "Really Focused"
}

func (view *View) resetFocus() string {
	return " Focus Resetted"
}

type ScreenView struct {
}

func (view *ScreenView) doDisplay() string {
	return " Display with Screen View"
}
