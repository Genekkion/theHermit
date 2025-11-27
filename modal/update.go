package modal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/genekkion/theHermit/utils"
)

// Update implements the tea.Model interface.
func (m Model) Update(msg tea.Msg) (model tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowResize(msg)
	default:
		cmds := make([]tea.Cmd, 0, 2)
		if m.isShown {
			// Note that the child is only updated if the modal is shown.
			m.child, cmd = m.child.Update(msg)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
		m.parent, cmd = m.parent.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}

		m.triggerCacheViews()

		return m, tea.Batch(cmds...)
	}
}

// handleWindowResize handles window resizes, accounting for the max dims if set.
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

	// Update the cached values
	m.cache.leftPadWidth = (msg.Width - m.dims.Width) / 2
	m.cache.startIndex = max(
		0,
		(m.winDims.Height-m.dims.Height)/2,
	)
	m.triggerCacheViews()

	return m, tea.Batch(cmds...)
}

// triggerCacheViews triggers the cache views to be recalculated.
func (m *Model) triggerCacheViews() {
	m.cacheDepsViews()
	m.cache.endIndex = min(
		len(m.cache.parent.lines),
		m.winDims.Height,
		m.cache.startIndex+m.dims.Height,
	)

	m.cacheViews()
	m.cache.flags.reset()
}

// cacheViews caches the views based on the cached dependencies.
func (m *Model) cacheViews() {
	if m.cache.flags.parentModified {
		// If the parent model is modified, we need to regenerate
		// ALL views because the left and right paddings are from
		// the parent view.
		m.cache.topSpacer = m.generateTopSpacer()
		m.cache.topBorder = m.generateTopBorder()
		m.cache.content = m.generateContent()
		m.cache.bottomBorder = m.generateBottomBorder()
		m.cache.bottomSpacer = m.generateBottomSpacer()
	} else if m.cache.flags.childModified {
		m.cache.content = m.generateContent()
	}
}

// cacheDepsViews caches the dependencies of the views.
func (m *Model) cacheDepsViews() {
	// Because it is likely that the parent's view will not
	// change while the box is displayed, we can try caching
	// the parent view's lines so we don't have to split
	// it on every render.
	m.cache.flags.parentModified = m.cacheParentView()

	// While it may not be significant, we can also cache the
	// child view since updates will trigger rerenders which
	// may result in unnecessary recalculations.
	m.cache.flags.childModified = m.cacheChildView()
}

// cacheParentView caches the parent's view by using its hash. Returns if
// the parent's view has changed since the last time it was cached.
func (m *Model) cacheParentView() bool {
	defer m.cache.hash.Reset()
	view := m.parent.View()
	m.cache.hash.WriteString(view)
	value := m.cache.hash.Sum64()

	if m.cache.parent.hash == value {
		return false
	}
	m.cache.parent.hash = value

	// We need to split the parent's view into line by line
	// as we will have to modify some lines to render the box instead.
	m.cache.parent.lines, m.cache.parent.widths, m.cache.parent.maxWidth = utils.Lines(view)
	return true
}

// cacheChildView is similar to cacheParentView, but caches the child's view instead.
func (m *Model) cacheChildView() bool {
	defer m.cache.hash.Reset()
	view := m.child.View()
	m.cache.hash.WriteString(view)
	value := m.cache.hash.Sum64()

	if m.cache.child.hash == value {
		return false
	}
	m.cache.child.hash = value

	m.cache.child.lines, m.cache.child.widths, m.cache.child.maxWidth = utils.Lines(view)
	return true
}
