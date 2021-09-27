package main

const (
	no_help_topic = -1
	print_topic = 1
	paper_orientation_topic = 2
	application_topic = 3
)

type HelpHandler interface {
	hasHelp() bool
	setHandler(HelpHandler)
	handleHelp() string
}

type ConcreteHelpHandler struct {
	successor HelpHandler
	topic int
}

func (helpHandler *ConcreteHelpHandler) hasHelp() bool {
	return helpHandler.topic != no_help_topic
}

func (helpHandler *ConcreteHelpHandler) handleHelp() string {
	if (helpHandler.successor != nil) { 
		helpHandler.successor.handleHelp()
	}
	return ""
}

// Implements HelpHandler
type Button struct {
	successor HelpHandler
	topic int
	parent HelpHandler
}

func (button *Button) hasHelp() bool {
	return button.topic != no_help_topic
}

func (button *Button) handleHelp() string {
	if (button.hasHelp()) {
		return "Button help"
	} else {
		return button.successor.handleHelp()
	}
}

func newButton(helpHandler HelpHandler, topic int) *Button {
	return &Button{topic: topic, parent: helpHandler}
}

type Dialog struct {
	successor HelpHandler
	topic int
}

func (dialog *Dialog) hasHelp() bool {
	return dialog.topic == print_topic
}

func (dialog *Dialog) handleHelp() string {
	if (dialog.hasHelp()) {
		return "Dialog help"
	} else {
		return dialog.successor.handleHelp()
	}
}

func (dialog *Dialog) setHandler(helpHandler HelpHandler) {
	dialog.successor = helpHandler
}

func newDialog(helpHandler HelpHandler, topic int) *Dialog {
	return &Dialog{topic: topic, successor: helpHandler}
}

type Application struct {
	successor HelpHandler
	topic int
}

func (application *Application) hasHelp() bool {
	return application.topic != no_help_topic
}

func (application *Application) handleHelp() string {
	if (application.hasHelp()) {
		return "Application help"
	} else {
		return application.successor.handleHelp()
	}
}

func (application *Application) setHandler(helpHandler HelpHandler) {
	application.successor = helpHandler
}

func newApplication(topic int) *Application {
	return &Application{topic: topic}
}

func main(){
	application := newApplication(application_topic)
	dialog := newDialog(application, print_topic)
	button := newButton(dialog, paper_orientation_topic)
	button.handleHelp()
}