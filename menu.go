package menu

import (
	"strings"
)

//VisibilityType - the type of visibility
type VisibilityType int

const (
	// BOTH - visible on both logged or logged status
	BOTH VisibilityType = 0
	// NOTLOGGED - visible when the status is not logged
	NOTLOGGED VisibilityType = 1
	// LOGGED - visible when the status is logged
	LOGGED VisibilityType = 2
)

// BaseMenu - base menu struct (Not rendered)
type BaseMenu struct {
	MenuID      string
	HasChildren bool
	Text        string
	Icon        string
	Link        string
	OnClick     string
	Visible     bool
	Visibility  VisibilityType
	Ordinal     int
	Children    []Menu
	Headers     []Header
}

// RootMenu - the root of the menus (not rendered)
type RootMenu struct {
	Children    []MainMenu
	Initialized bool
	Logged      bool
}

// Menu - a common menu that can be rendered
type Menu struct {
	BaseMenu
	HeaderSubscription string
}

// Header - a header for the menu
type Header struct {
	Text       string
	Subscripts []Menu //lists the menus that are under this header
}

// MainMenu - a menu that cannot have links
type MainMenu struct {
	BaseMenu
}

// Add - adds a main menu to the root menu
func (rm *RootMenu) Add(menu *MainMenu) *MainMenu {
	if !rm.Initialized {
		rm = &RootMenu{
			Initialized: true,
		}
	}

	rm.Children = append(rm.Children, *menu)
	idx := len(rm.Children) - 1

	// Append a blank header if the main being added has not yet defined its headers
	cptr := &rm.Children[idx]
	if len(cptr.Headers) == 0 {
		cptr.Headers = append(cptr.Headers, Header{Text: ""})
	}

	return cptr
}

// Evaluate - evaluates various properties set before being assigned to the template
func (rm *RootMenu) Evaluate() {
	for _, m := range rm.Children {
		m.Visible = true
		switch m.Visibility {
		case LOGGED:
			if !rm.Logged {
				m.Visible = false
			}
		case NOTLOGGED:
			if rm.Logged {
				m.Visible = false
			}
		}

		for _, h := range m.Headers {
			for _, s := range h.Subscripts {
				s.Visible = true
				switch s.Visibility {
				case LOGGED:
					if !rm.Logged {
						s.Visible = false
					}
				case NOTLOGGED:
					if rm.Logged {
						s.Visible = false
					}
				}
			}
		}
	}
}

// AddHeader - add headers to the menu
func (mm *MainMenu) AddHeader(header *Header) (*Header, error) {
	//check if the header text is already in this menu
	idx := -1

	if len(mm.Headers) == 0 {
		mm.Headers = append(mm.Headers, Header{})
	}

	if len(mm.Headers) > 0 {
		for i, h := range mm.Headers {
			if strings.ToLower(h.Text) == strings.ToLower(header.Text) {
				idx = i
				break
			}
		}
	}

	if idx == -1 {
		mm.Headers = append(mm.Headers, *header)
		idx = 0
		hptr := &mm.Headers[idx]
		return hptr, nil
	}

	hptr := &mm.Headers[idx]
	return hptr, nil
}

// AddHeader - add headers to the menu
func (m *Menu) AddHeader(header *Header) (*Header, error) {
	//check if the header text is already in this menu
	idx := -1

	if len(m.Headers) == 0 {
		m.Headers = append(m.Headers, Header{})
	}

	if len(m.Headers) > 0 {
		for i, h := range m.Headers {
			if strings.ToLower(h.Text) == strings.ToLower(header.Text) {
				idx = i
				break
			}
		}
	}

	if idx == -1 {
		m.Headers = append(m.Headers, *header)
		idx = 0
		hptr := &m.Headers[idx]
		return hptr, nil
	}

	hptr := &m.Headers[idx]
	return hptr, nil
}

// AddMenu - add a sub menu to a main menu
func (mm *MainMenu) AddMenu(menu *Menu) *Menu {
	mm.Children = append(mm.Children, *menu)
	mm.HasChildren = true

	idx := len(mm.Children) - 1
	mptr := &mm.Children[idx]

	// Look for the header added as indicated from the menu on the parentMenu's list of header
	idx = -1

	if len(mm.Headers) == 0 {
		mm.Headers = append(mm.Headers, Header{})
	}

	if len(mm.Headers) > 0 {
		for i, h := range mm.Headers {
			if strings.ToLower(h.Text) == strings.ToLower(mptr.HeaderSubscription) {
				idx = i
				break
			}
		}
	}

	// If header is not found, we refer to the default header which is blank
	if idx == -1 {
		idx = 0
	}

	mm.Headers[idx].Subscripts = append(mm.Headers[idx].Subscripts, *mptr)

	return mptr
}
