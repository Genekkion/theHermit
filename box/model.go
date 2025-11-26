package box

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/genekkion/theHermit/shared"
)

// Model is the model for the box which implements the tea.Model interface.
type Model struct {
	isShown bool

	dims    shared.Dimensions
	winDims shared.Dimensions
	Config

	parent tea.Model
	child  tea.Model

	// For rendering
	builder strings.Builder
	cache   ViewCache
}

// ViewCache stores any cache-related items such as hashes
// as well as the cached views.
type ViewCache struct {
	parentHash  uint64
	parentLines []string
	childHash   uint64
	childLines  []string

	leftPadding    int
	leftPaddingStr string
}

func New() Model {
	return Model{

		builder: strings.Builder{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
