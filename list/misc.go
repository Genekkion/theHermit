package list

import (
	"github.com/charmbracelet/lipgloss"
)

func (model Model) GetIsNumbered() bool {
	return model.isNumbered
}

func (model *Model) SetIsNumbered(isNumbered bool) {
	model.isNumbered = isNumbered
}

func (model Model) GetIsShown() bool {
	return model.isShown
}

func (model *Model) SetIsShown(isShown bool) {
	model.isShown = isShown
}

func (model Model) GetHeight() int {
	return model.height
}

func (model *Model) SetHeight(height int) {
	model.height = height
}

func (model Model) GetWidth() int {
	return model.width
}

func (model *Model) SetWidth(width int) {
	model.width = width
}

func (model Model) GetMaxHeight() int {
	return model.maxHeight
}

func (model *Model) SetMaxHeight(maxHeight int) {
	model.maxHeight = maxHeight
}

func (model Model) GetMaxWidth() int {
	return model.maxWidth
}

func (model *Model) SetMaxWidth(maxWidth int) {
	model.maxWidth = maxWidth
}

func (model Model) GetSelectedItem() Item {
	return model.items[model.cursor]
}

func (model Model) Cursor() int {
	return model.cursor
}

func (model Model) SetCursor(cursor int) Model {
	model.cursor = max(min(cursor, len(model.items)-1), 0)
	return model
}

func (model Model) GetItems() []Item {
	return model.items
}

func (model *Model) SetItems(items []Item) {
	model.items = items
}

func (model Model) Title() string {
	return model.title
}

func (model *Model) SetTitle(title string) {
	model.title = title
}

func (model *Model) SetView(view string) {
	model.view = view
}

func (model Model) GetView() string {
	return model.view
}

func (model *Model) SetBorder(border lipgloss.Border) {
	model.border = border
}

func (model Model) GetBorder() lipgloss.Border {
	return model.border
}

func (model *Model) SetBorderForeground(color lipgloss.Color) {
	model.borderStyle = model.borderStyle.Foreground(color)
}

func (model *Model) SetBorderBackground(color lipgloss.Color) {
	// model.borderStyle = model.borderStyle.BorderBackground(color)
	model.borderStyle = model.borderStyle.Background(color)
}

func (model *Model) SetTitleForeground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Foreground(color)
}

func (model *Model) SetTitleBackground(color lipgloss.Color) {
	model.titleStyle = model.titleStyle.Background(color)
}

func (model *Model) SetTitleBold(bold bool) {
	model.titleStyle = model.titleStyle.Bold(bold)
}

func (model *Model) SetItemForeground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Foreground(color)
}

func (model *Model) SetItemBackground(color lipgloss.Color) {
	model.itemStyle = model.itemStyle.Background(color)
}

func (model *Model) SetItemBold(bold bool) {
	model.itemStyle = model.itemStyle.Bold(bold)
}

func (model *Model) SetSelectedForeground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Foreground(color)
}

func (model *Model) SetSelectedBackground(color lipgloss.Color) {
	model.selectedStyle = model.selectedStyle.Background(color)
}

func (model *Model) SetSelectedBold(bold bool) {
	model.selectedStyle = model.selectedStyle.Bold(bold)
}
