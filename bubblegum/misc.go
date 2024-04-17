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

func (model Bubblegum) GetHeight() int {
	return model.height
}

func (model *Bubblegum) SetHeight(height int) {
	model.height = height
}

func (model Bubblegum) GetWidth() int {
	return model.width
}

func (model *Bubblegum) SetWidth(width int) {
	model.width = width
}

func (model Bubblegum) GetMaxHeight() int {
	return model.maxHeight
}

func (model *Bubblegum) SetMaxHeight(maxHeight int) {
	model.maxHeight = maxHeight
}

func (model Bubblegum) GetMaxWidth() int {
	return model.maxWidth
}

func (model *Bubblegum) SetMaxWidth(maxWidth int) {
	model.maxWidth = maxWidth
}

func (model Bubblegum) GetSelectedItem() Item {
	return model.items[model.cursor]
}

func (model Bubblegum) Cursor() int {
	return model.cursor
}

func (model Bubblegum) GetItems() []Item {
	return model.items
}

func (model *Bubblegum) SetItems(items []Item) {
	model.items = items
}

func (model Bubblegum) Title() string {
	return model.title
}

func (model *Bubblegum) SetTitle(title string) {
	model.title = title
}

func (model *Bubblegum) SetView(view string) {
	model.view = view
}

func (model Bubblegum) GetView() string {
	return model.view
}

func (model *Bubblegum) SetBorder(border lipgloss.Border) {
	model.border = border
}

func (model Bubblegum) GetBorder() lipgloss.Border {
	return model.border
}

func (model *Bubblegum) SetBorderForeground(color lipgloss.Color) {
	model.borderStyle = model.borderStyle.BorderForeground(color)
}

func (model *Bubblegum) SetBorderBackground(color lipgloss.Color) {
	model.borderStyle = model.borderStyle.BorderBackground(color)
}

func (model *Bubblegum) SetTitleForeground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Foreground(color)
}

func (model *Bubblegum) SetTitleBackground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Background(color)
}

func (model *Bubblegum) SetTitleBold(bold bool) {
	model.titleStyle = model.titleStyle.Bold(bold)
}

func (model *Bubblegum) SetItemForeground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Foreground(color)
}

func (model *Bubblegum) SetItemBackground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Background(color)
}

func (model *Bubblegum) SetItemBold(bold bool) {
	model.itemStyle = model.itemStyle.Bold(bold)
}

func (model *Bubblegum) SetSelectedForeground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Foreground(color)
}

func (model *Bubblegum) SetSelectedBackground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Background(color)
}

func (model *Bubblegum) SetSelectedBold(bold bool) {
	model.selectedStyle = model.selectedStyle.Bold(bold)
}

