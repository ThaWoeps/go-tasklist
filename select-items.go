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


//let's actually do something
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {
	
	// Is it a key press?
	case tea.KeyMsg:
		
		// Excellent,  so what's the actual key pressed?
		switch msg.String() {

		// these keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit
		
		// the "up" and "k" key move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--  
			}

		// the "down" and "j" key move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		
		// the "Enter" key and spacebar (a literal space) toggel
		// the selected state for the item that the cursor is pointing at
		case "enter", " ":
			_, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string { //this STRING is our UI!
    // The header
    s := "What should we buy at the market?\n\n"

    // Iterate over our choices
    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func main() { //Execute the bloody program
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}