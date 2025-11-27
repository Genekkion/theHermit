package box

import (
	"hash/maphash"
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
		return ""
	} else if m.winDims.Width == 0 || m.winDims.Height == 0 {
		// If the window is flattened, we do not render anything.
		return ""
	}

	parentLines, _, parentWidth := utils.Lines(m.parent.View())
	if len(parentLines) == 0 || parentWidth == 0 {
		return ""
	}

	m.builder.Reset()

	// Calculate where to insert the box
	startIndex := (m.winDims.Height / 2) - (m.dims.Height / 2)
	endIndex := startIndex + m.dims.Height

	m.writeTopSpacer(parentLines[:startIndex])
	m.writeTopBorder(parentLines[startIndex])
	m.writeContent(parentLines, startIndex)
	m.writeBottomBorder(parentLines[endIndex-1])
	m.writeBottomSpacer(parentLines[endIndex:])

	return m.builder.String()
}

func (m *Model) writeTopSpacer(parentLines []string) {
	for _, line := range parentLines {
		m.builder.WriteString(line)
		m.builder.WriteByte('\n')
	}
}

func (m *Model) writeBottomSpacer(parentLines []string) {
	n := len(parentLines) - 1
	for i, line := range parentLines {
		m.builder.WriteString(line)
		if i != n {
			m.builder.WriteByte('\n')
		}
	}
}

func (m *Model) writeContent(parentLines []string, startIndex int) {
	unsetStyle := m.style.
		UnsetPadding().
		UnsetMargins().
		UnsetBorderStyle()
	border, _, _, _, _ := m.style.GetBorder()

	childLines, childWidths, _ := utils.Lines(m.child.View())
	childLimit := min(len(childLines), m.dims.Height-2)

	for i, line := range childLines[:childLimit] {
		leftPad := m.generateLeftPadding(
			utils.SplitColumns(parentLines[startIndex+i+1]),
		)
		m.builder.WriteString(leftPad)

		m.builder.WriteString(unsetStyle.Render(border.Left))

		m.builder.WriteString(line[:m.dims.Width-2])
		if childWidths[i] < m.dims.Width-2 {
			spacer := strings.Repeat(" ", m.dims.Width-2-childWidths[i])
			m.builder.WriteString(unsetStyle.Render(spacer))
		}

		m.builder.WriteString(unsetStyle.Render(border.Right))

		rightPad := m.generateRightPadding(
			utils.SplitColumns(parentLines[startIndex+i+1]),
		)
		m.builder.WriteString(rightPad)

		m.builder.WriteByte('\n')
	}

	for i := range m.dims.Height - 2 - childLimit {
		leftPad := m.generateLeftPadding(
			utils.SplitColumns(parentLines[startIndex+childLimit+i+1]),
		)
		m.builder.WriteString(leftPad)

		m.builder.WriteString(unsetStyle.Render(border.Left))

		spacer := strings.Repeat(" ", m.dims.Width-2)
		m.builder.WriteString(unsetStyle.Render(spacer))

		m.builder.WriteString(unsetStyle.Render(border.Right))

		rightPad := m.generateRightPadding(
			utils.SplitColumns(parentLines[startIndex+childLimit+i+1]),
		)
		m.builder.WriteString(rightPad)

		m.builder.WriteByte('\n')
	}
}

func (m *Model) writeTopBorder(line string) {
	chars := utils.SplitColumns(line)
	m.builder.WriteString(m.generateLeftPadding(chars))
	m.builder.WriteString(topBorder(m.dims.Width, m.style, m.title))
	m.builder.WriteString(m.generateRightPadding(chars))
	m.builder.WriteByte('\n')
}

func (m *Model) writeBottomBorder(line string) {
	chars := utils.SplitColumns(line)
	m.builder.WriteString(m.generateLeftPadding(chars))
	m.builder.WriteString(bottomBorder(m.dims.Width, m.style))
	m.builder.WriteString(m.generateRightPadding(chars))
	m.builder.WriteByte('\n')
}

// cacheParentView caches the parent's view by using its hash. Returns if
// the parent's view has changed since the last time it was cached.
func (m *Model) cacheParentView() bool {
	hash := maphash.Hash{}
	hash.WriteString(m.parent.View())
	value := hash.Sum64()
	if m.cache.parentHash == value {
		return false
	}
	m.cache.parentHash = value

	// We need to split the parent's view into line by line
	// as we will have to modify some lines to render the box instead.
	m.cache.parentLines = strings.Split(m.parent.View(), "\n")
	return true
}

// cacheChildView is similar to cacheParentView, but caches the child's view instead.
func (m *Model) cacheChildView() bool {
	hash := maphash.Hash{}
	hash.WriteString(m.child.View())
	value := hash.Sum64()
	if m.cache.childHash == value {
		return false
	}
	m.cache.childHash = value

	m.cache.childLines = strings.Split(m.parent.View(), "\n")
	return true
}

func (m *Model) generateLeftPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadding)
	return generateLeftPadding(chars, limit)
}

func (m *Model) writeLeftPadding(chars []string) {
	m.builder.WriteString(m.generateLeftPadding(chars))
}

func generateLeftPadding(chars []string, width int) string {
	return strings.Join(chars[:min(len(chars), width)], "")
}
func (m *Model) generateRightPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadding+m.dims.Width)
	return generateRightPadding(chars, limit)
}

func generateRightPadding(chars []string, width int) string {
	if len(chars) <= width {
		return strings.Repeat(" ", width-len(chars))
	}
	return strings.Join(chars[width:], "")
}

func topBorder(width int, style lipgloss.Style, tt *title.Title) string {
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
		renderedTop := unsetStyle.Render(
			strings.Repeat(border.Top, remainingWidth/2),
		)
		builder.WriteString(renderedTop)
		builder.WriteString(ttStr)
		builder.WriteString(renderedTop)
		if remainingWidth%2 == 1 {
			builder.WriteString(unsetStyle.Render(border.Top))
		}
	}
	builder.WriteString(unsetStyle.Render(border.TopRight))

	return builder.String()
}

func bottomBorder(width int, style lipgloss.Style) string {
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
