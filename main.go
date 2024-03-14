package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// model is a struct that holds all the stuff we need to store our
// applications state.
type model struct {
	
}

// initialModel initializes and returns the model. This is not to 
// be confused with the bubbletea.Init() method. This one prepares
// our model, the other let's start with some kind of command.
func initialModel() model {
	return model{}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	// For now, we do nothing here, so we hae nothing to return
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	// s is a string that will be returned with our view for rendering.
	s := ""

	// Build the view
	s += fmt.Sprintf("%s %s\n\n%s\n", ">", "Hello from 'tea saucer' ;)", "Press q to Quit.")

	// Return the view for rendering.
	return s
}

// The entry point of the appliction.
func main() {
	log.Println("Program 'tea saucer' started.")

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal("Error running tea saucer:", err)
	}

	log.Println("Program 'tea saucer' terminated.")
}
