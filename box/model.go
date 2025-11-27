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
	parent ViewCacheModel
	child  ViewCacheModel

	leftPadding    int
	leftPaddingStr string

	startIndex int
	endIndex   int
}

type ViewCacheModel struct {
	hash     uint64
	lines    []string
	widths   []int
	maxWidth int
}

func New(dimensions shared.Dimensions, parent tea.Model, child tea.Model, opts ...Option) (m *Model, err error) {
	if parent == nil {
		return nil, ErrMissingParent
	} else if child == nil {
		return nil, ErrMissingChild
	}

	config := defaultConfig()
	for _, opt := range opts {
		opt(&config)
	}
	if config.maxDimensions == nil {
		config.maxDimensions = &dimensions
	}

	m = &Model{
		dims:    dimensions,
		builder: strings.Builder{},
		Config:  config,

		parent: parent,
		child:  child,
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}
