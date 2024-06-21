package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/list"
)

type MainModel struct {
	list         list.Model
	selectedText string
	style        lipgloss.Style
}

type ListItem struct {
	title string
}

func (item ListItem) Title() string {
	return item.title
}

func (item ListItem) FilterValue() string {
	return item.title
}

var defaultListItems = []list.Item{
	ListItem{
		title: "hello world!",
	},
	ListItem{
		title: "these are all items!",
	},
	ListItem{
		title: "press 'enter' to select an item",
	},
}

func DefaultModel() MainModel {
	return MainModel{
		selectedText: "Nothing is selected",
		style: lipgloss.NewStyle().
			Background(lipgloss.Color("#23283B")),

		list: list.NewDefault(defaultListItems),
	}
}

func (model MainModel) Init() tea.Cmd {
	return nil
}

func (model MainModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		// log.Println(message.Width, message.Height)
		model.style = model.style.
			Width(message.Width).
			Height(message.Height)
		listModel, _ := model.list.Update(message)
		model.list = listModel.(list.Model)
		return model, nil

	case tea.KeyMsg:
		switch message.String() {
		case "ctrl+c":
			return model, tea.Quit

		case "ctrl+e":
			model.list.SetIsShown(!model.list.GetIsShown())
			return model, nil

		case "ctrl+r":
			model.list.SetIsNumbered(!model.list.GetIsNumbered())
			return model, nil

		}

		if model.list.GetIsShown() {
			listModel, command := model.list.Update(message)
			model.list = listModel.(list.Model)
			if message.String() == "enter" {
				model.selectedText = model.list.GetSelectedItem().Title()
			}
			return model, command
		}
	}

	return model, nil
}

func (model MainModel) View() string {
	builder := strings.Builder{}
	builder.WriteString("Selected text: " + model.selectedText + "\n\n")
	builder.WriteString("Press 'ctrl+e' to toggle the simple list\n")
	builder.WriteString("Press 'ctrl+r' to toggle the numbering\n")
	builder.WriteString("Press 'ctrl+c' to quit!\n")

	currentHeight := lipgloss.Height(builder.String())
	for range model.style.GetHeight() - currentHeight {
		builder.WriteString("-\n")
	}
	builder.WriteString("-")

	model.list.SetView(model.style.Render(builder.String()))
	return model.list.View()
}

func main() {
	program := tea.NewProgram(
		DefaultModel(),
		tea.WithAltScreen(),
	)
	_, err := program.Run()
	if err != nil {
		log.Fatal(err)
	}
}
