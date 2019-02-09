package menu

import (
	"fmt"
	"testing"
)

func TestEarlyMenuAdding(t *testing.T) {
	// create a main menu
	dm := MainMenu{}
	dm.MenuID = "domain"
	dm.Text = "Domain Management"
	dm.AddHeader(&Header{Text: "DOMAIN USERS"})
	dm.AddHeader(&Header{Text: "DOMAIN APPLICATIONS"})
	dm.AddHeader(&Header{Text: "DOMAIN COMPANIES & GROUPS"})

	// Add a sub menu with header subscription at users
	dms := Menu{HeaderSubscription: "DOMAIN USERS"}
	dms.Text = "Manage Users"
	dms.Icon = "user"
	dms.MenuID = "user"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Application Master"
	dms.Icon = "app"
	dms.MenuID = "app"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Module Master"
	dms.Icon = "mod"
	dms.MenuID = "mod"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Permission Master"
	dms.Icon = "perm"
	dms.MenuID = "perm"
	dm.AddMenu(&dms)

	// Add a new main menu
	em := MainMenu{}
	em.MenuID = "licenses"
	em.Text = "Application Licenses"

	// Add a sub menu with no header subscription
	dms = Menu{}
	dms.Text = "Application"
	dms.Icon = "licenses"
	dms.MenuID = "lic"
	em.AddMenu(&dms)

	hm := MainMenu{}
	hm.MenuID = "workflow"
	hm.Text = "Workflow Management"

	// Add a sub menu with no header subscription
	dms = Menu{}
	dms.Text = "Application"
	dms.Icon = "appwf"
	dms.MenuID = "appwf"
	hm.AddMenu(&dms)

	dms = Menu{}
	dms.Text = "Route Template"
	dms.Icon = "rtwf"
	dms.MenuID = "rtwf"
	hm.AddMenu(&dms)

	dms = Menu{}
	dms.Text = "Live Routes"
	dms.Icon = "lrwf"
	dms.MenuID = "lrwf"
	hm.AddMenu(&dms)

	rm := RootMenu{Initialized: true}
	rm.Add(&dm)
	rm.Add(&em)
	rm.Add(&hm)

	// Render
	for _, m := range rm.Children {
		fmt.Println("main: " + m.Text)
		for _, h := range m.Headers {
			fmt.Println("head:--" + h.Text)
			for _, s := range h.Subscripts {
				fmt.Println("menu:----" + s.Text)
			}
		}
	}
}

func TestEarlyLateAdding(t *testing.T) {
	rm := RootMenu{Initialized: true}
	dm := rm.Add(&MainMenu{})

	dm.MenuID = "domain"
	dm.Text = "Domain Management"
	dm.AddHeader(&Header{Text: "DOMAIN USERS"})
	dm.AddHeader(&Header{Text: "DOMAIN APPLICATIONS"})
	dm.AddHeader(&Header{Text: "DOMAIN COMPANIES & GROUPS"})

	// Add a sub menu with header subscription at users
	dms := Menu{HeaderSubscription: "DOMAIN USERS"}
	dms.Text = "Manage Users"
	dms.Icon = "user"
	dms.MenuID = "user"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Application Master"
	dms.Icon = "app"
	dms.MenuID = "app"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Module Master"
	dms.Icon = "mod"
	dms.MenuID = "mod"
	dm.AddMenu(&dms)

	dms = Menu{HeaderSubscription: "DOMAIN APPLICATIONS"}
	dms.Text = "Permission Master"
	dms.Icon = "perm"
	dms.MenuID = "perm"
	dm.AddMenu(&dms)

	// Add a new main menu
	em := rm.Add(&MainMenu{})
	em.MenuID = "licenses"
	em.Text = "Application Licenses"

	// Add a sub menu with no header subscription
	dms = Menu{}
	dms.Text = "Application"
	dms.Icon = "licenses"
	dms.MenuID = "lic"
	em.AddMenu(&dms)

	hm := rm.Add(&MainMenu{})
	hm.MenuID = "workflow"
	hm.Text = "Workflow Management"

	// Add a sub menu with no header subscription
	dms = Menu{}
	dms.Text = "Application"
	dms.Icon = "appwf"
	dms.MenuID = "appwf"
	hm.AddMenu(&dms)

	dms = Menu{}
	dms.Text = "Route Template"
	dms.Icon = "rtwf"
	dms.MenuID = "rtwf"
	hm.AddMenu(&dms)

	dms = Menu{}
	dms.Text = "Live Routes"
	dms.Icon = "lrwf"
	dms.MenuID = "lrwf"
	hm.AddMenu(&dms)

	// Render
	for _, m := range rm.Children {
		fmt.Println("main: " + m.Text)
		for _, h := range m.Headers {
			fmt.Println("head:--" + h.Text)
			for _, s := range h.Subscripts {
				fmt.Println("menu:----" + s.Text)
			}
		}
	}
}
