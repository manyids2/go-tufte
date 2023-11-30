package components

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/manyids2/go-tufte/core"
	"github.com/rivo/tview"
)

type mode int

const (
	Normal mode = iota
	Command
	Toggle
)

// View of app
type App struct {
	// Necessities
	Doc         *core.Document
	Application *tview.Application
	Pages       *tview.Pages
	Page        string

	// Windows
	Content   *tview.TextView
	Sidebar   *tview.TextView
	Statusbar *tview.TextView

	// Modes for keeping track of toggle
	Mode mode

	// Keep track of focus
	focusCycle   []*tview.TextView
	focusCurrent int

	// Keep track of sections
	focusSection int
}

func NewApp(doc *core.Document) *App {
	app := App{
		Doc:   doc,
		Pages: tview.NewPages(),
		Mode:  Normal,
	}

	// Create fullscreen app
	app.Application = tview.NewApplication()

	// Add home
	app.AddHomePage()

	// Set global keymaps
	app.focusCycle = []*tview.TextView{app.Sidebar, app.Content}
	app.SetKeymaps()

	// Add it to page and display
	app.Application.SetRoot(app.Pages, true).SetFocus(app.Sidebar)

	// Show info
	app.SetStatusbar()

	// Table of contents
	app.SetSidebar()

	// Try to print rendered string
	app.SetContent()

	return &app
}

func (a *App) SetContent() {
	if len(a.Doc.Sections) > 0 {
		s := a.Doc.Sections[a.focusSection]
		a.Content.Clear()
		// TODO: Why cant I slice buffer??
		content := string(*a.Doc.Buffer)[s.StartByte:s.EndByte]
		fmt.Fprintf(a.Content, content)
	}
}

func (a *App) SetStatusbar() {
	a.Statusbar.Clear()
	title := ""
	if len(a.Doc.Sections) > 0 {
		title = a.Doc.Sections[a.focusSection].Title
	}
	fmt.Fprintf(a.Statusbar, fmt.Sprintf("  %s > %s > %s", a.Doc.Path, a.Doc.Title, title))
}

func (a *App) SetSidebar() {
	a.Sidebar.Clear()
	for i, s := range a.Doc.Sections {
		indent := strings.Repeat("  ", s.Level-1)
		if i == a.focusSection {
			fmt.Fprintf(a.Sidebar, fmt.Sprintf("|%s|> %s\n", indent, s.Title))
		} else {
			fmt.Fprintf(a.Sidebar, fmt.Sprintf(" %s   %s\n", indent, s.Title))
		}
	}
}

func (a *App) Update() {
	a.SetStatusbar()
	a.SetSidebar()
	a.SetContent()
}

func (a *App) HandleSidebar(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyDown:
		a.focusSection = (a.focusSection + 1) % len(a.Doc.Sections)
		a.Update()
		return nil
	case tcell.KeyUp:
		a.focusSection = a.focusSection - 1
		if a.focusSection < 0 {
			a.focusSection = len(a.Doc.Sections) - 1
		}
		a.Update()
		return nil
	}
	switch event.Rune() {
	case 'j':
		a.focusSection = (a.focusSection + 1) % len(a.Doc.Sections)
		a.Update()
		return nil
	case 'k':
		a.focusSection = a.focusSection - 1
		if a.focusSection < 0 {
			a.focusSection = len(a.Doc.Sections) - 1
		}
		a.Update()
		return nil
	}
	return event
}

func (a *App) AddHomePage() {
	// Create home as default
	a.Page = "home"
	page := tview.NewGrid().SetRows(1, -1, 1).SetColumns(-1, -2)

	// Content
	a.Content = tview.NewTextView().
		SetChangedFunc(func() {
			a.Application.Draw()
		})
	a.Content.SetBorder(true)
	page.AddItem(a.Content, 1, 1, 1, 1, 0, 0, true)

	// Sidebar
	a.Sidebar = tview.NewTextView().
		SetChangedFunc(func() {
			a.Application.Draw()
		})
	a.Sidebar.SetBorder(true)
	a.Sidebar.SetWrap(false)
	page.AddItem(a.Sidebar, 1, 0, 1, 1, 0, 0, false)
	a.Sidebar.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return a.HandleSidebar(event)
	})

	// Statusbar
	a.Statusbar = tview.NewTextView().
		SetChangedFunc(func() {
			a.Application.Draw()
		})
	page.AddItem(a.Statusbar, 2, 0, 1, 2, 0, 0, false)

	// Home page
	a.Pages.AddPage("home", page, true, true)
}

// ToggleFocus
func (a *App) ToggleFocus() {
	a.focusCurrent = (a.focusCurrent + 1) % len(a.focusCycle)
	a.Application.SetFocus(a.focusCycle[a.focusCurrent])
}

// Quit, help, toggle
func (a *App) SetKeymaps() {
	a.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			a.Application.Stop()
			return nil
		case tcell.KeyTab:
			a.ToggleFocus()
			return nil
		}
		switch event.Rune() {
		case 'q':
			a.Application.Stop()
			return nil
		}
		return event
	})
}
