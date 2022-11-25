package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
)

//adding boiler plate code from the bubbletea github,  adapting later
type model struct { // #LEARN check why this is a struct and not a map
	choices []string				//items on the todo list
	cursor int						// which to-do list item our cursor is pointing on
	selected map[int]struct{}		// which to-do items are selected
}

func initialModel() model {  // #LEARN what are models in go
	return model{
		// Our to-do list is a grovery list
		choices: []string{"Buy carrots", "Buy celery", "Buy radish"},

		// a map which indicates which choices are slected.
		// the map like a methermatical set.  The keys refer to the indexes
		// of the choises slice, above.
		selected: make(map[int]struct{}),

	}
}

func (m model) Init() tea.Cmd {
	// Just return 'nil', which means "no I/O right now, please."
	return nil
}