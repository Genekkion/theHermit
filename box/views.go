package box

import (
	"hash/maphash"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/genekkion/theHermit/shared/title"
)

func (m Model) View() string {
	if !m.isShown {
		// If the box is not shown, return the parent's view without obstruction.
		return m.parent.View()
	} else if m.winDims.Width == 0 || m.winDims.Height == 0 {
		// If the window is flattened, we do not render anything.
		return ""
	}

	// Calculate where to insert the box
	//startIndex := (m.winDims.Height / 2) - (m.dims.Height / 2) + 1
	//endIndex := startIndex + m.dims.Height

	//for _, line := range m.parentLines {
	//}

	return ""
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

func (m *Model) cache12() {

	//chars := utils.SplitColumns(line)

}

func (m *Model) generateLeftPadding(chars []string) string {
	limit := min(m.winDims.Width, m.cache.leftPadding)
	return generateLeftPadding(chars, limit)
}

func generateLeftPadding(chars []string, limit int) string {
	return strings.Join(chars[:limit], "")
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
	builder.WriteString(border.TopLeft)

	{
		top := style.UnsetBorderStyle().Render(border.Top)
		builder.WriteString(strings.Repeat(top, remainingWidth/2))
		builder.WriteString(ttStr)
		builder.WriteString(strings.Repeat(top, remainingWidth/2))
		if remainingWidth%2 == 1 {
			builder.WriteString(top)
		}
	}
	builder.WriteString(border.TopRight)

	return builder.String()
}
