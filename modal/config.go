package modal

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared"
	"github.com/genekkion/theHermit/shared/title"
)

type Config struct {
	title         *title.Title
	maxDimensions *shared.Dimensions
	style         lipgloss.Style
}

func defaultConfig() Config {
	return Config{
		title: nil,
	}
}

type Option func(*Config)

func WithTitle(t title.Title) Option {
	return func(c *Config) {
		c.title = &t
	}
}

func WithMaxDimensions(dimensions shared.Dimensions) Option {
	return func(c *Config) {
		c.maxDimensions = &dimensions
	}
}

func WithStyle(style lipgloss.Style) Option {
	return func(c *Config) {
		c.style = style
	}
}
