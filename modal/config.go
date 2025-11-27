package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared"
	"github.com/genekkion/theHermit/shared/title"
)

// Config holds the configuration for the modal.
type Config struct {
	title         *title.Title
	maxDimensions *shared.Dimensions
	style         lipgloss.Style
}

// defaultConfig returns a default configuration for the modal.
func defaultConfig() Config {
	return Config{
		title:         nil,
		maxDimensions: nil,
		style:         lipgloss.NewStyle(),
	}
}

// Option is a functional option for configuring the modal.
type Option func(*Config)

// WithTitle sets the title of the modal.
func WithTitle(t title.Title) Option {
	return func(c *Config) {
		c.title = &t
	}
}

// WithMaxDimensions sets the maximum dimensions of the modal.
func WithMaxDimensions(dimensions shared.Dimensions) Option {
	return func(c *Config) {
		c.maxDimensions = &dimensions
	}
}

// WithStyle sets the style of the modal.
func WithStyle(style lipgloss.Style) Option {
	return func(c *Config) {
		c.style = style
	}
}
