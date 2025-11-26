package box

import (
	"testing"

	"github.com/Genekkion/gogogadgets/pkg/ptr"
	"github.com/Genekkion/gogogadgets/pkg/test"
	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared/title"
	"github.com/genekkion/theHermit/utils"
)

func TestGenerateLeftPadding(t *testing.T) {
	type Data struct {
		name     string
		input    string
		limit    int
		expected string
	}

	testData := []Data{
		{
			name:     "simple ASCII",
			input:    "ABCDEFGHIJKLMNOPQRSTUVQXYZ",
			limit:    5,
			expected: "ABCDE",
		},
		{
			name:     "emoji",
			input:    "🌟🎉🎨✨🚀",
			limit:    3,
			expected: "🌟🎉🎨",
		},
		{
			name:     "ANSI colored",
			input:    "\x1b[31mRed\x1b[0m\x1b[32mGreen\x1b[0m",
			limit:    2,
			expected: "\x1b[31mRe",
		},
		{
			name:     "emoji and ANSI",
			input:    "🌟\x1b[31mRed\x1b[0m✨",
			limit:    2,
			expected: "🌟\x1b[31mR",
		},
		{
			name:     "CJK wide char",
			input:    "你好世界",
			limit:    2,
			expected: "你好",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			t.Parallel()

			chars := utils.SplitColumns(data.input)
			generated := generateLeftPadding(chars, data.limit)
			test.AssertEqual(t, "Unexpected padding", data.expected, generated)
		})
	}
}

func TestTopBorder(t *testing.T) {
	type Data struct {
		name     string
		width    int
		title    *title.Title
		style    lipgloss.Style
		expected string
	}

	testData := []Data{
		{
			name:  "simple with title even",
			width: 14,
			title: ptr.New(title.New("Hello", lipgloss.NewStyle())),
			style: lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()),
			expected: "┌── Hello ───┐",
		},
		{
			name:  "simple with title odd",
			width: 15,
			title: ptr.New(title.New("Hello", lipgloss.NewStyle())),
			style: lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()),
			expected: "┌─── Hello ───┐",
		},
		{
			name:  "simple without title even",
			width: 14,
			title: nil,
			style: lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()),
			expected: "┌────────────┐",
		},
		{
			name:  "simple without title odd",
			width: 15,
			title: nil,
			style: lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()),
			expected: "┌─────────────┐",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			t.Parallel()

			test.AssertEqual(t, "Unexpected render",
				data.expected,
				topBorder(data.width, data.style, data.title),
			)
		})
	}
}

func TestBottomBorder(t *testing.T) {
	type Data struct {
		name     string
		width    int
		style    lipgloss.Style
		expected string
	}

	testData := []Data{
		{
			name:  "simple",
			width: 15,
			style: lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()),
			expected: "└─────────────┘",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			t.Parallel()

			test.AssertEqual(t, "Unexpected render",
				data.expected,
				bottomBorder(data.width, data.style),
			)
		})
	}
}
