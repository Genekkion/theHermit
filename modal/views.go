package modal

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared/title"
	"github.com/genekkion/theHermit/utils"
)

// View implements the tea.Model interface.
func (m Model) View() string {
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

	// Write all values.
	m.writeTopSpacer()
	m.writeTopBorder()
	m.writeContent()
	m.writeBottomBorder()
	m.writeBottomSpacer()

	// Return the rendered string.
	return m.builder.String()
}

// writeTopSpacer writes the top spacer string to the builder.
func (m *Model) writeTopSpacer() {
	m.builder.WriteString(m.cache.topSpacer)
}

// generateTopSpacer returns the top spacer string, i.e. the lines above the
// modal.
func (m *Model) generateTopSpacer() string {
	builder := strings.Builder{}

	for _, line := range m.cache.parent.lines[:m.cache.startIndex] {
		builder.WriteString(line)
		builder.WriteByte('\n')
	}
	return builder.String()
}

// writeBottomSpacer writes the bottom spacer string to the builder.
func (m *Model) writeBottomSpacer() {
	m.builder.WriteString(m.cache.bottomSpacer)
}

// generateBottomSpacer returns the bottom spacer string, i.e. the lines below
// the modal.
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

// writeContent writes the content string to the builder.
func (m *Model) writeContent() {
	m.builder.WriteString(m.cache.content)
}

// generateContent returns the content string, i.e. the lines which include the modal itself
// and the child view. Note that the top and bottom borders are not included and have
// separate functions.
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

		line = line[:max(0, min(childWidths[i], m.dims.Width-2))]
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

// writeTopBorder writes the top border string to the builder.
func (m *Model) writeTopBorder() {
	m.builder.WriteString(m.cache.topBorder)
}

// generateTopBorder returns the top border string, i.e. the line above the modal, which may
// include the title if present.
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

// generateTopBorder returns the top border string.
func generateTopBorder(width int, style lipgloss.Style, tt *title.Title) string {
	ttStr := ""
	if tt != nil {
		ttStr = tt.Render()
	}

	availableWidth := width - 2
	ttCols := utils.SplitColumns(ttStr)
	ttWidth := len(ttCols)
	if ttWidth >= availableWidth {
		return strings.Join(ttCols[:availableWidth], "")
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

// writeBottomBorder writes the bottom border string to the builder.
func (m *Model) writeBottomBorder() {
	m.builder.WriteString(m.cache.bottomBorder)
}

// generateBottomBorder returns the bottom border string, i.e. the line below the modal.
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

// generateBottomBorder returns the bottom border string.
func generateBottomBorder(width int, style lipgloss.Style) string {
	builder := strings.Builder{}
	border, _, _, _, _ := style.GetBorder()
	unsetStyle := style.UnsetBorderStyle()

	builder.WriteString(unsetStyle.Render(border.BottomLeft))
	bottom := unsetStyle.Render(
		strings.Repeat(border.Bottom, width-2),
	)
	builder.WriteString(bottom)
	builder.WriteString(unsetStyle.Render(border.BottomRight))

	return builder.String()
}

// generateLeftPadding returns the left padding string, i.e. the characters from the parent view,
// which are present on the left side of the modal when rendered.
func (m *Model) generateLeftPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadWidth)
	return generateLeftPadding(chars, limit)
}

// generateLeftPadding returns the left padding string.
func generateLeftPadding(chars []string, width int) string {
	return strings.Join(chars[:min(len(chars), width)], "")
}

// generateRightPadding returns the right padding string, i.e. the characters from the parent view,
// which are present on the right side of the modal when rendered.
func (m *Model) generateRightPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadWidth+m.dims.Width)
	return generateRightPadding(chars, limit)
}

// generateRightPadding returns the right padding string.
func generateRightPadding(chars []string, width int) string {
	if len(chars) <= width {
		return strings.Repeat(" ", width-len(chars))
	}
	return strings.Join(chars[width:], "")
}
