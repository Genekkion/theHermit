package box

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared/title"
	"github.com/genekkion/theHermit/utils"
)

func (m Model) View() string {
	m.isShown = true

	if !m.isShown {
		// If the box is not shown, return the parent's view without obstruction.
		if m.parent != nil {
			return m.parent.View()
		}
		// Should never happen
		panic("parent is nil")

	} else if m.winDims.Width == 0 || m.winDims.Height == 0 ||
		m.cache.parent.maxWidth == 0 || len(m.cache.parent.lines) == 0 {

		// If the window is flattened, we do not render anything.
		return ""
	}

	// Unfortunately, the builder has to be reset every time we render
	// since the number of bytes per line is not constant.
	m.builder.Reset()

	m.writeTopSpacer()
	m.writeTopBorder()
	m.writeContent()
	m.writeBottomBorder()
	m.writeBottomSpacer()

	return m.builder.String()
}

func (m *Model) writeTopSpacer() {
	//m.builder.WriteString(m.generateTopSpacer())
	m.builder.WriteString(m.cache.topSpacer)
}

func (m *Model) generateTopSpacer() string {
	builder := strings.Builder{}

	for _, line := range m.cache.parent.lines[:m.cache.startIndex] {
		builder.WriteString(line)
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (m *Model) writeBottomSpacer() {
	//m.builder.WriteString(m.generateBottomSpacer())
	m.builder.WriteString(m.cache.bottomSpacer)
}

func (m *Model) generateBottomSpacer() string {
	builder := strings.Builder{}

	parentLines := m.cache.parent.lines[m.cache.endIndex:]
	n := len(parentLines) - 1
	for i, line := range parentLines {
		builder.WriteString(line)
		if i != n {
			builder.WriteByte('\n')
		}
	}

	return builder.String()
}

// m.builder.WriteStringContent(parentLines, startIndex)
func (m *Model) writeContent() {
	//m.builder.WriteString(m.generateContent())
	m.builder.WriteString(m.cache.content)
}

func (m *Model) generateContent() string {
	builder := strings.Builder{}

	unsetStyle := m.style.
		UnsetPadding().
		UnsetMargins().
		UnsetBorderStyle()
	border, _, _, _, _ := m.style.GetBorder()

	startIndex := m.cache.startIndex
	parentLines := m.cache.parent.lines[startIndex+1:]
	childLines := m.cache.child.lines
	childWidths := m.cache.child.widths
	childLimit := max(0, min(len(childLines), m.dims.Height-2))
	if childLimit == 0 {
		return ""
	}

	for i, line := range childLines[:childLimit] {
		leftPad := m.generateLeftPadding(
			utils.SplitColumns(parentLines[i]),
		)
		builder.WriteString(leftPad)

		builder.WriteString(unsetStyle.Render(border.Left))

		line = line[:max(0, min(len(line), m.dims.Width-2))]
		builder.WriteString(line)
		if childWidths[i] < m.dims.Width-2 {
			spacer := strings.Repeat(" ", m.dims.Width-2-childWidths[i])
			builder.WriteString(unsetStyle.Render(spacer))
		}

		builder.WriteString(unsetStyle.Render(border.Right))

		rightPad := m.generateRightPadding(
			utils.SplitColumns(parentLines[i]),
		)
		builder.WriteString(rightPad)

		builder.WriteByte('\n')
	}

	parentLines = parentLines[childLimit:]
	for i := range m.dims.Height - 2 - childLimit {
		leftPad := m.generateLeftPadding(
			utils.SplitColumns(parentLines[i]),
		)
		builder.WriteString(leftPad)

		builder.WriteString(unsetStyle.Render(border.Left))

		spacer := strings.Repeat(" ", m.dims.Width-2)
		builder.WriteString(unsetStyle.Render(spacer))

		builder.WriteString(unsetStyle.Render(border.Right))

		rightPad := m.generateRightPadding(
			utils.SplitColumns(parentLines[i]),
		)
		builder.WriteString(rightPad)

		builder.WriteByte('\n')
	}

	return builder.String()
}

func (m *Model) writeTopBorder() {
	//m.builder.WriteString(m.generateTopBorder())
	m.builder.WriteString(m.cache.topBorder)
}

func (m *Model) generateTopBorder() string {
	builder := strings.Builder{}
	line := m.cache.parent.lines[m.cache.startIndex]
	chars := utils.SplitColumns(line)
	builder.WriteString(m.generateLeftPadding(chars))
	builder.WriteString(generateTopBorder(m.dims.Width, m.style, m.title))
	builder.WriteString(m.generateRightPadding(chars))
	builder.WriteByte('\n')
	return builder.String()
}

func (m *Model) writeBottomBorder() {
	//m.builder.WriteString(m.generateBottomBorder())
	m.builder.WriteString(m.cache.bottomBorder)
}

func (m *Model) generateBottomBorder() string {
	builder := strings.Builder{}
	line := m.cache.parent.lines[m.cache.endIndex-1]
	chars := utils.SplitColumns(line)
	builder.WriteString(m.generateLeftPadding(chars))
	builder.WriteString(generateBottomBorder(m.dims.Width, m.style))
	builder.WriteString(m.generateRightPadding(chars))
	builder.WriteByte('\n')
	return builder.String()
}

func (m *Model) writeLeftPadding(chars []string) {
	m.builder.WriteString(m.generateLeftPadding(chars))
}

func (m *Model) generateLeftPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadWidth)
	return generateLeftPadding(chars, limit)
}

func generateLeftPadding(chars []string, width int) string {
	return strings.Join(chars[:min(len(chars), width)], "")
}

func (m *Model) generateRightPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadWidth+m.dims.Width)
	return generateRightPadding(chars, limit)
}

func generateRightPadding(chars []string, width int) string {
	if len(chars) <= width {
		return strings.Repeat(" ", width-len(chars))
	}
	return strings.Join(chars[width:], "")
}

func generateTopBorder(width int, style lipgloss.Style, tt *title.Title) string {
	ttStr := ""
	if tt != nil {
		ttStr = tt.Render()
	}

	availableWidth := width - 2
	ttWidth := lipgloss.Width(ttStr)
	if ttWidth >= availableWidth {
		return ttStr[:availableWidth]
	}
	remainingWidth := availableWidth - ttWidth

	builder := strings.Builder{}
	border, _, _, _, _ := style.GetBorder()
	unsetStyle := style.
		UnsetPadding().
		UnsetMargins().
		UnsetBorderStyle()

	builder.WriteString(unsetStyle.Render(border.TopLeft))

	{
		renderedBorderTop := unsetStyle.Render(
			strings.Repeat(border.Top, remainingWidth/2),
		)
		builder.WriteString(renderedBorderTop)
		builder.WriteString(ttStr)
		builder.WriteString(renderedBorderTop)
		if remainingWidth%2 == 1 {
			builder.WriteString(unsetStyle.Render(border.Top))
		}
	}
	builder.WriteString(unsetStyle.Render(border.TopRight))

	return builder.String()
}

func generateBottomBorder(width int, style lipgloss.Style) string {
	builder := strings.Builder{}
	border, _, _, _, _ := style.GetBorder()
	style = style.UnsetBorderStyle()

	builder.WriteString(style.Render(border.BottomLeft))
	bottom := style.Render(
		strings.Repeat(border.Bottom, width-2),
	)
	builder.WriteString(style.Render(bottom))
	builder.WriteString(style.Render(border.BottomRight))

	return builder.String()
}
