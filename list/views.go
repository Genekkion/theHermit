package list

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Regex for identifiying styles used. Mainly used for
	// splitting the styles within a single line.
	colorRegex = regexp.MustCompile("\033\\[[0-9;]+m")
)

// Writes the left padding for the list to the string builder.
func (model *List) writeLeftPadding(stringBuilder *strings.Builder, chars *[]string) {
	index := 0
	limit := min(model.windowWidth, model.leftPadding)

	// Iterate through the list by width due to
	// ANSI codes placed within the line
	for lipgloss.Width(stringBuilder.String()) < limit {
		stringBuilder.WriteString((*chars)[index])
		index++
	}
	// When content is less than the spacer
	for lipgloss.Width(stringBuilder.String()) < model.leftPadding {
		stringBuilder.WriteByte(' ')
	}
}

// Auxillary function used by writeRightPadding to find the
// valid chars at the end of the line.
func paddingLength(array []string) int {
	stringBuilder := strings.Builder{}
	for _, char := range array {
		stringBuilder.WriteString(char)
	}
	return lipgloss.Width(stringBuilder.String())
}

// Searches through an array produced by a regex's
// FindAllStringIndex function, to check if the extra
// bytes are from ANSI codes.
func isCode(regexMatches [][]int, index int) bool {
	for _, match := range regexMatches {
		if index >= match[0] && index <= match[1] {
			return true
		}
	}
	return false
}

// Writes the right padding for the list to the string builder.
func (model *List) writeRightPadding(stringBuilder *strings.Builder, chars *[]string,
	line *string) {

	currentWidth := lipgloss.Width(stringBuilder.String())
	if currentWidth == model.windowWidth {
		stringBuilder.WriteByte('\n')
		return
	}

	colorCodes := colorRegex.FindAllStringIndex(*line, -1)

	rightPaddingLength := model.windowWidth - currentWidth - 1
	index := len(*chars) - 1
	rightPadding := []string{}
	for isCode(colorCodes, index) || paddingLength(rightPadding) < rightPaddingLength {
		rightPadding = append([]string{(*chars)[index]}, rightPadding...)
		index--
	}
	if len(colorCodes) != 0 {
		colorIndexStart := colorCodes[0][0]
		colorIndexEnd := colorCodes[0][1]
		for i, code := range colorCodes {
			if i == 0 {
				continue
			} else if code[0] < index {
				colorIndexStart = code[0]
				colorIndexEnd = code[1]
			} else {
				break
			}
		}
		for _, char := range (*chars)[colorIndexStart:colorIndexEnd] {
			stringBuilder.WriteString(char)
		}
	}
	for _, char := range (*chars)[index:] {
		stringBuilder.WriteString(char)
	}
	stringBuilder.WriteByte('\n')
}

// Returns the string for the top border of the list,
// accounting for the background text.
func (model *List) topBorder(line *string) string {
	text := strings.Builder{}
	chars := strings.Split(*line, "")

	// Left padding from background
	model.writeLeftPadding(&text, &chars)

	// Account for border
	availableWidth := model.width - 2

	title := ""

	// Account for no title
	if len(model.title) > 0 {
		title = fmt.Sprintf(" %s ", model.title)
	}

	// Account for title overflow
	if lipgloss.Width(title) > availableWidth {
		title = title[:availableWidth]
	}

	stringBuilder := strings.Builder{}

	// Start writing left side
	stringBuilder.WriteString(model.border.TopLeft)
	for range (availableWidth - lipgloss.Width(title)) / 2 {
		stringBuilder.WriteString(model.border.Top)
	}

	// Write title
	stringBuilder.WriteString(model.titleStyle.Render(title))
	text.WriteString(model.borderStyle.Render(stringBuilder.String()))

	// Write right side
	borderLimit := availableWidth - lipgloss.Width(stringBuilder.String())
	borderCounter := 0
	stringBuilder.Reset()
	for borderCounter <= borderLimit {
		stringBuilder.WriteString(model.border.Top)
		borderCounter++
	}
	stringBuilder.WriteString(model.border.TopRight)

	text.WriteString(model.borderStyle.Render(stringBuilder.String()))

	// Right padding from background
	model.writeRightPadding(&text, &chars, line)
	return text.String()
}

// Returns the string for the bottom border of the list,
// accounting for the background text.
func (model *List) bottomBorder(line *string) string {
	text := strings.Builder{}
	chars := strings.Split(*line, "")

	// Left padding from background
	model.writeLeftPadding(&text, &chars)

	// Build the border
	text.WriteString(model.borderStyle.Render(model.border.BottomLeft))
	for range model.width - 2 {
		text.WriteString(model.borderStyle.Render(model.border.Bottom))
	}
	text.WriteString(model.borderStyle.Render(model.border.BottomRight))

	// Right padding from background
	model.writeRightPadding(&text, &chars, line)
	return text.String()
}

// Returns the string for the items on the list,
// accounting for background text.
func (model *List) middleBorder(line *string, item *Item, index int) string {
	text := strings.Builder{}
	chars := strings.Split(*line, "")

	// Left padding from background
	model.writeLeftPadding(&text, &chars)

	// Left border
	text.WriteString(model.borderStyle.Render(model.border.Left))

	var itemText string

	// Account for item numbering
	if model.isNumbered {
		itemText = fmt.Sprintf("%d. %s", index+1, (*item).Title())
		if lipgloss.Width(itemText) > model.width-2 {
			itemText = itemText[:model.width-2]
		}

		// Account for selected item
		if index == model.cursor {
			itemText = model.selectedStyle.
				Bold(true).
				Render(itemText)
		} else {
			itemText = model.itemStyle.Render(itemText)
		}

	} else {
		itemText = (*item).Title()
		if lipgloss.Width(itemText) > model.width-2 {
			itemText = itemText[:model.width-2]
		}

		// Account for selected item
		if index == model.cursor {
			itemText = model.selectedStyle.
				Bold(true).
				Render(itemText)
		} else {
			itemText = model.itemStyle.Render(itemText)
		}
	}

	text.WriteString(itemText)

	// Build the empty space remaining in the row
	spacer := strings.Builder{}
	for range model.width - lipgloss.Width(itemText) - 2 {
		spacer.WriteByte(' ')
	}
	text.WriteString(model.borderStyle.Render(spacer.String()))

	// Right border
	text.WriteString(model.borderStyle.Render(model.border.Right))

	// Right padding from background
	model.writeRightPadding(&text, &chars, line)

	return text.String()
}

// FOr use when there are lesser items than there are rows.
func (model List) middleSpacer(line *string) string {
	text := strings.Builder{}
	chars := strings.Split(*line, "")

	// Left padding from background
	model.writeLeftPadding(&text, &chars)

	// Left border
	text.WriteString(model.borderStyle.Render(model.border.Left))

	// Build the spacer
	spacer := strings.Builder{}
	for range model.width - 2 {
		spacer.WriteByte(' ')
	}
	text.WriteString(model.borderStyle.Render(spacer.String()))

	// Right border
	text.WriteString(model.borderStyle.Render(model.border.Right))

	// Right padding from background
	model.writeRightPadding(&text, &chars, line)

	return text.String()
}

// The main view for the list.
func (model List) View() string {
	if model.windowWidth == 0 || model.windowHeight == 0 {
		return ""
	}

	text := strings.Builder{}

	// Needs to be converted to array and processed
	// line by line
	arrayView := strings.Split(model.view, "\n")

	// Calculate where to insert the list
	midPoint1 := model.windowHeight/2 - model.height/2 + 1
	midPoint2 := midPoint1 + model.height

	// Get items based on offset, due to cursor going
	// beyond the initial list
	var items []Item
	if len(model.items) <= model.height-2 {
		items = model.items
	} else {
		items = model.items[model.offset : model.offset+model.height-1]
	}
	itemLength := len(items)

	itemIndex := 0

	length := len(arrayView) - 1
	for i, line := range arrayView {
		switch {
		case i == midPoint1:
			text.WriteString(model.topBorder(&line))
		case i == midPoint2:
			text.WriteString(model.bottomBorder(&line))
		case i > midPoint1 && i < midPoint2:
			if itemIndex < itemLength {
				text.WriteString(model.middleBorder(&line, &items[itemIndex],
					itemIndex+model.offset))
				itemIndex++
			} else {
				text.WriteString(model.middleSpacer(&line))
			}
		default:
			text.WriteString(line)
			if i < length {
				text.WriteByte('\n')
			}
		}
	}
	return text.String()
}
