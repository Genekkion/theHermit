package title

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Title is a simple struct that holds the title value and the style to render it
// in.
type Title struct {
	value string
	style lipgloss.Style
}

// New creates a new Title struct with the given value and style.
func New(value string, style lipgloss.Style) Title {
	value = fmt.Sprintf(" %s ", value)

	return Title{
		value: value,
		style: style,
	}
}

// NewDefault creates a new Title struct with the given value and default style.
func NewDefault(value string) Title {
	return New(value, lipgloss.NewStyle())
}

// Value returns the title value.
func (t Title) Value() string {
	return t.value
}

// SetValue sets the title value.
func (t *Title) SetValue(value string) {
	t.value = value
}

// Style returns the style of the title.
func (t Title) Style() lipgloss.Style {
	return t.style
}

// SetStyle sets the style of the title.
func (t *Title) SetStyle(style lipgloss.Style) {
	t.style = style
}

// Render returns the rendered title.
func (t Title) Render() string {
	return t.style.Render(t.value)
}
