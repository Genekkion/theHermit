package box

import tea "github.com/charmbracelet/bubbletea"

// Update implements the tea.Model interface.
func (m Model) Update(msg tea.Msg) (model tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowResize(msg)
	default:
		cmds := make([]tea.Cmd, 2)
		m.child, cmd = m.child.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		m.parent, cmd = m.parent.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}
}

// handleWindowResize handles window resizes, accounting for the max dims
// if set.
func (m Model) handleWindowResize(msg tea.WindowSizeMsg) (model tea.Model, cmd tea.Cmd) {
	// Update the height
	if msg.Height >= m.dims.Height {
		m.dims.Height = min(m.maxDimensions.Height, msg.Height)
	} else {
		m.dims.Height = min(m.dims.Height, msg.Height)
	}

	// Update the width
	if msg.Width >= m.dims.Width {
		m.dims.Width = min(m.maxDimensions.Width, msg.Width)
	} else {
		m.dims.Width = min(m.dims.Width, msg.Width)
	}

	m.cacheViews(msg)

	// Update the latest window dims
	m.winDims.Height = msg.Height
	m.winDims.Width = msg.Width

	// Propagate the message to the parent and child
	cmds := make([]tea.Cmd, 0, 2)
	m.parent, cmd = m.parent.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	m.child, cmd = m.child.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) cacheViews(msg tea.WindowSizeMsg) {
	// Update the padding
	m.updateLeftPadding(msg.Width)

	// Because it is likely that the parent's view will not
	// change while the box is displayed, we can try caching
	// the parent view's lines so we don't have to split
	// it on every render.
	//parentModified := m.cacheParentView()

	// While it may not be significant, we can also cache the
	// child view since updates will trigger rerenders which
	// may result in unnecessary recalculations.
	//childModified := m.cacheChildView()
}

func (m *Model) updateLeftPadding(msgWidth int) {
	// Update the padding
	newPadding := (msgWidth - m.dims.Width) / 2
	if m.cache.leftPadding == newPadding {
		return
	}

	m.cache.leftPadding = newPadding
	//m.generateLeftPadding()
}
