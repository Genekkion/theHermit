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

	parentLines := strings.Split(m.parent.View(), "\n")
	//for i, line := range parentLines {
	//	fmt.Printf("PARENT %d: '%s'\n", i, line)
	//}
	if len(parentLines) == 0 || lipgloss.Width(m.parent.View()) == 0 {
		return "EMPTY"
	}

	// Calculate where to insert the box
	startIndex := (m.winDims.Height / 2) - (m.dims.Height / 2)
	for _, line := range parentLines[:startIndex] {
		m.builder.WriteString(line)
		m.builder.WriteByte('\n')
	}
	m.writeTopBorder(parentLines[startIndex])

	// CONTENT
	childLines := strings.Split(m.child.View(), "\n")
	childLimit := min(len(childLines), m.dims.Height-2)
	for i, line := range childLines[:childLimit] {
		leftPad := m.generateLeftPadding(utils.SplitColumns(parentLines[startIndex+i+1]))
		m.builder.WriteString(leftPad)
		border, _, _, _, _ := m.style.GetBorder()
		m.builder.WriteString(
			m.style.UnsetBorderStyle().Render(border.Left),
		)

		m.builder.WriteString(line)
		spacer := strings.Repeat(" ", m.dims.Width-2-lipgloss.Width(line))
		spacer = m.style.UnsetBorderStyle().Render(spacer)
		m.builder.WriteString(spacer)

		m.builder.WriteString(
			m.style.UnsetBorderStyle().Render(border.Right),
		)

		rightPad := m.generateRightPadding(utils.SplitColumns(parentLines[startIndex+i+1]))
		m.builder.WriteString(rightPad)

		m.builder.WriteByte('\n')
	}
	for i := range m.dims.Height - 2 - childLimit {
		leftPad := m.generateLeftPadding(utils.SplitColumns(parentLines[startIndex+childLimit+i+1]))
		m.builder.WriteString(leftPad)

		border, _, _, _, _ := m.style.GetBorder()
		m.builder.WriteString(
			m.style.UnsetBorderStyle().Render(border.Left),
		)

		spacer := strings.Repeat(" ", m.dims.Width-2)
		m.builder.WriteString(
			m.style.UnsetBorderStyle().Render(spacer),
		)
		m.builder.WriteString(
			m.style.UnsetBorderStyle().Render(border.Right),
		)

		rightPad := m.generateRightPadding(utils.SplitColumns(parentLines[startIndex+childLimit+i+1]))
		m.builder.WriteString(rightPad)

		m.builder.WriteString("\n")
	}

	endIndex := startIndex + m.dims.Height
	m.writeBottomBorder(parentLines[endIndex-1])
	for _, line := range parentLines[endIndex:] {
		m.builder.WriteString(line)
		m.builder.WriteByte('\n')
	}

	//for _, line := range m.parentLines {
	//}

	s := m.builder.String()
	return s[:len(s)-1]
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
	style = style.UnsetBorderStyle()

	builder.WriteString(style.Render(border.TopLeft))

	{
		top := style.Render(
			strings.Repeat(border.Top, remainingWidth/2),
		)
		builder.WriteString(top)
		builder.WriteString(ttStr)
		builder.WriteString(top)
		if remainingWidth%2 == 1 {
			builder.WriteString(style.Render(border.Top))
		}
	}
	builder.WriteString(style.Render(border.TopRight))

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
