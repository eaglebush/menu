package menu

import "strings"

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
}

// RootMenu - the root of the menus (not rendered)
type RootMenu struct {
	BaseMenu
	Children    []MainMenu
	Initialized bool
}

// Menu - a common menu that can be rendered
type Menu struct {
	BaseMenu
	Headers            []Header
	Text               string
	Icon               string
	Link               string
	Visible            bool
	Visibility         VisibilityType
	HeaderSubscription string
	Ordinal            int
	Children           []Menu
}

// Header - a header for the menu
type Header struct {
	Text       string
	Subscripts []Menu //lists the menus that are under this header
}

// MainMenu - a menu that cannot have links
type MainMenu struct {
	BaseMenu
	Headers  []Header
	Children []Menu
	Ordinal  int
}

// AddMainMenu - adds a main menu to the root menu
func (pm *RootMenu) AddMainMenu(menu *MainMenu) *MainMenu {
	if !pm.Initialized {
		pm = &RootMenu{
			Initialized: true,
		}
	}

	pm.Children = append(pm.Children, *menu)
	pm.HasChildren = true

	idx := len(pm.Children) - 1

	// Append a blank header if the main being added has not yet defined its headers
	cptr := &pm.Children[idx]
	if len(cptr.Headers) == 0 {
		cptr.Headers = append(cptr.Headers, Header{Text: ""})
	}

	return cptr
}

// AddMainMenuHeader - add headers to the menu
func (mm *MainMenu) AddMainMenuHeader(header *Header, parentMenu *MainMenu) (*Header, error) {
	//check if the header text is already in this menu
	idx := -1
	for i, h := range parentMenu.Headers {
		if strings.ToLower(h.Text) == strings.ToLower(header.Text) {
			idx = i
			break
		}
	}

	if idx != -1 {
		parentMenu.Headers = append(parentMenu.Headers, *header)
		idx = 0
		hptr := &parentMenu.Headers[idx]
		return hptr, nil
	}

	hptr := &parentMenu.Headers[idx]
	return hptr, nil
}

// AddMenuHeader - add headers to the menu
func (m *Menu) AddMenuHeader(header *Header) (*Header, error) {
	//check if the header text is already in this menu
	idx := -1
	for i, h := range m.Headers {
		if strings.ToLower(h.Text) == strings.ToLower(header.Text) {
			idx = i
			break
		}
	}

	if idx != -1 {
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
	for i, h := range mm.Headers {
		if strings.ToLower(h.Text) == strings.ToLower(mptr.HeaderSubscription) {
			idx = i
			break
		}
	}

	// If header is not found, we refer to the default header which is blank
	if idx == -1 {
		idx = 0
	}

	mm.Headers[idx].Subscripts = append(mm.Headers[idx].Subscripts, *mptr)

	return mptr
}
