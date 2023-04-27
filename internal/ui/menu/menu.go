package menu

import (
	"sync"

	"github.com/lucaskatayama/pigo/internal/infra/icons"
	"github.com/lucaskatayama/pigo/internal/ui/display"
)

type Menu struct {
	mu sync.RWMutex

	Sel   int
	Items []Item
	_len  int
}

func (m *Menu) Next() {
	m.mu.Lock()
	m.Sel += 1
	if m.Sel > m.Len()-1 {
		m.Sel = 0
	}
	m.mu.Unlock()
}

func (m *Menu) Previous() {
	m.mu.Lock()
	m.Sel -= 1
	if m.Sel < 0 {
		m.Sel = m.Len() - 1
	}
	m.mu.Unlock()
}

func (m *Menu) Len() int {
	return m._len
}

func (m *Menu) Get(i int) Item {
	return m.Items[i]
}

func (m *Menu) PrevItem() Item {
	m.mu.RLock()
	i := m.Sel - 1
	m.mu.RUnlock()
	if i < 0 {
		i = m.Len() - 1
	}
	return m.Get(i)
}

func (m *Menu) NextItem() Item {
	m.mu.RLock()
	i := m.Sel + 1
	m.mu.RUnlock()

	if i > m.Len()-1 {
		i = 0
	}
	return m.Get(i)
}

func (m *Menu) CurrItem() Item {
	return m.Get(m.Sel)
}

// Draw draws menu to display
func (m *Menu) Draw(s display.Display) {
	s.Begin()
	//
	s.DrawImageBytes(icons.ScrollbarBackground, 121, 0)
	s.DrawImageBytes(icons.ItemSelOutline, -1, 21)

	prev := m.PrevItem()
	curr := m.CurrItem()
	next := m.NextItem()

	// Previous Item
	s.DrawImageBytes(prev.Icon, 4, 4)
	s.Text(prev.Label, 18, 4)

	// Selected Item
	s.DrawImageBytes(curr.Icon, 4, 23)
	s.Text(curr.Label, 18, 24)

	// Next Item
	s.DrawImageBytes(next.Icon, 4, 44)
	s.Text(next.Label, 18, 44)

	// Scroll
	s.DrawRect(125, (64/m.Len())*m.Sel, 3, 64/m.Len())

	// Flush page
	s.Flush()
}

func (m *Menu) OnEnter() {

}

// OnClick handles click event for menu
func (m *Menu) OnClick() {
	m.Next()
}

// OnDoubleClick handles double click event for menu
func (m *Menu) OnDoubleClick() {
	m.Previous()
}

// OnPress handles long press event for menu
func (m *Menu) OnPress() {
	s := m.CurrItem().Sub
	s.OnEnter()
	SetDraw(s)
}
