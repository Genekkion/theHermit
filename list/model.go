package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	isShown bool

	// Height & width of the entire model
	height int
	width  int

	// Max limits on the height & width
	maxHeight int
	maxWidth  int

	// Window dimensions which are updated during
	// the Update function
	windowHeight int
	windowWidth  int

	// Precalculated leftPadding length during
	// window resizes
	leftPadding int

	// Misc details
	title      string
	isNumbered bool

	// Working mechanisms for the list
	cursor int
	offset int
	items  []Item
	view   string

	// Styles
	borderStyle   lipgloss.Style
	titleStyle    lipgloss.Style
	selectedStyle lipgloss.Style
	itemStyle     lipgloss.Style

	// Border is specificied seperately and DOES NOT
	// come from any of the above styles
	border lipgloss.Border
}

// Initialises a List with sensible defaults.
func NewDefault(items []Item) Model {
	width := 81

	return Model{
		cursor:     0,
		height:     14,
		width:      width,
		maxHeight:  14,
		maxWidth:   width,
		isNumbered: true,
		isShown:    false,

		items:  items,
		offset: 0,

		border: lipgloss.NormalBorder(),
		borderStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("#16161D")).
			BorderBackground(lipgloss.Color("#16161D")).
			Foreground(lipgloss.Color("#2D4F67")),
		titleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7FB4CA")).
			Bold(true),
		itemStyle: lipgloss.NewStyle().
			Width(width - 2).
			Background(lipgloss.Color("#16161D")).
			Foreground(lipgloss.Color("#DCD7BA")),
		selectedStyle: lipgloss.NewStyle().
			Width(width - 2).
			Background(lipgloss.Color("#DCD7BA")).
			Foreground(lipgloss.Color("#1F1F28")).
			Bold(true),
	}
}

// Initialises a List instance
func New(height int, width int, items []Item) Model {
	return Model{
		cursor:     0,
		height:     height,
		width:      width,
		maxHeight:  height,
		maxWidth:   width,
		isNumbered: true,
		isShown:    false,

		items:  items,
		offset: 0,

		border: lipgloss.NormalBorder(),
		borderStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("#16161D")).
			BorderBackground(lipgloss.Color("#16161D")).
			BorderForeground(lipgloss.Color("#2D4F67")),
		titleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7FB4CA")).
			Bold(true),
		itemStyle: lipgloss.NewStyle().
			Width(width - 2).
			BorderBackground(lipgloss.Color("#16161D")).
			BorderForeground(lipgloss.Color("#2D4F67")),
		selectedStyle: lipgloss.NewStyle().
			Width(width - 2).
			Background(lipgloss.Color("#DCD7BA")).
			Foreground(lipgloss.Color("#1F1F28")).
			Bold(true),
	}
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		// Updates the height and width accordingly to fit
		// fit all screen sizes.
		// WARN: May be buggy for really tiny screens :(
		if message.Height >= model.height {
			model.height = min(model.maxHeight, message.Height)
		} else {
			model.height = min(model.height, message.Height)
		}

		if message.Width >= model.width {
			model.width = min(model.maxWidth, message.Width)
		} else {
			model.width = min(model.width, message.Width)
		}

		model.leftPadding = (message.Width - model.width) / 2
		model.windowHeight = message.Height
		model.windowWidth = message.Width

	case tea.KeyMsg:
		switch message.String() {
		case "up":
			if model.cursor > 0 {
				model.cursor--
				if model.cursor < model.offset {
					model.offset--
				}
			}
		case "down":
			availableHeight := model.height - 2
			if len(model.items) <= availableHeight {
				if model.cursor < len(model.items)-1 {
					model.cursor++
				}
			} else {
				limit := model.offset + availableHeight
				if model.cursor >= model.offset && model.cursor < limit {
					model.cursor++
				} else if model.cursor >= limit && limit < len(model.items)-1 {
					model.cursor++
					model.offset++
				}
			}
		}
	}
	return model, nil
}
