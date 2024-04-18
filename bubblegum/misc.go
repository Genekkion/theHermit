package bubblegum

import "github.com/charmbracelet/lipgloss"

/*
	type Bubblegum struct {
		// Height & width of the entire model
		height int
		width  int

		// Max limits on the height & width
		maxHeight int
		maxWidth  int

		// Window dimensions which are updated during
		// the Update function
		windowHeight int
		windowWidth  int

		// Precalculated leftPadding length during
		// window resizes
		leftPadding int

		// Misc details
		title      string
		isNumbered bool

		// Working mechanisms for the list
		cursor int
		offset int
		items  []Item
		view   string

		// Styles
		borderStyle   lipgloss.Style
		titleStyle    lipgloss.Style
		selectedStyle lipgloss.Style
		itemStyle     lipgloss.Style

		// Border is specificied seperately and DOES NOT
		// come from any of the above styles
		border lipgloss.Border
	}
*/

func (model List) GetHeight() int {
	return model.height
}

func (model *List) SetHeight(height int) {
	model.height = height
}

func (model List) GetWidth() int {
	return model.width
}

func (model *List) SetWidth(width int) {
	model.width = width
}

func (model List) GetMaxHeight() int {
	return model.maxHeight
}

func (model *List) SetMaxHeight(maxHeight int) {
	model.maxHeight = maxHeight
}

func (model List) GetMaxWidth() int {
	return model.maxWidth
}

func (model *List) SetMaxWidth(maxWidth int) {
	model.maxWidth = maxWidth
}

func (model List) GetSelectedItem() Item {
	return model.items[model.cursor]
}

func (model List) Cursor() int {
	return model.cursor
}

func (model List) GetItems() []Item {
	return model.items
}

func (model *List) SetItems(items []Item) {
	model.items = items
}

func (model List) Title() string {
	return model.title
}

func (model *List) SetTitle(title string) {
	model.title = title
}

func (model *List) SetView(view string) {
	model.view = view
}

func (model List) GetView() string {
	return model.view
}

func (model *List) SetBorder(border lipgloss.Border) {
	model.border = border
}

func (model List) GetBorder() lipgloss.Border {
	return model.border
}

func (model *List) SetBorderForeground(color lipgloss.Color) {
	model.borderStyle = model.borderStyle.BorderForeground(color)
}

func (model *List) SetBorderBackground(color lipgloss.Color) {
	model.borderStyle = model.borderStyle.BorderBackground(color)
}

func (model *List) SetTitleForeground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Foreground(color)
}

func (model *List) SetTitleBackground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Background(color)
}

func (model *List) SetTitleBold(bold bool) {
	model.titleStyle = model.titleStyle.Bold(bold)
}

func (model *List) SetItemForeground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Foreground(color)
}

func (model *List) SetItemBackground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Background(color)
}

func (model *List) SetItemBold(bold bool) {
	model.itemStyle = model.itemStyle.Bold(bold)
}

func (model *List) SetSelectedForeground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Foreground(color)
}

func (model *List) SetSelectedBackground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Background(color)
}

func (model *List) SetSelectedBold(bold bool) {
	model.selectedStyle = model.selectedStyle.Bold(bold)
}
