package components

import (
	"log"

	"github.com/gdamore/tcell/v2"
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
	Application *tview.Application
	Pages       *tview.Pages
	Page        string

	// Windows
	Content   *tview.TextView
	Sidebar   *tview.TextView
	Statusbar *tview.TextView

	// Modes for keeping track of toggle
	Mode mode
}

func NewApp() *App {
	app := App{
		Pages: tview.NewPages(),
		Mode:  Normal,
	}

	// Create fullscreen app
	app.Application = tview.NewApplication()

	// Create home as default
	app.Page = "home"
	page := tview.NewGrid().SetRows(1, -1, 1).SetColumns(-1, -3)

	// Content
	app.Content = tview.NewTextView().
		SetChangedFunc(func() {
			app.Application.Draw()
		})
	app.Content.SetBorder(true)
	page.AddItem(app.Content, 1, 1, 1, 1, 0, 0, true)

	// Sidebar
	app.Sidebar = tview.NewTextView().
		SetChangedFunc(func() {
			app.Application.Draw()
		})
	app.Sidebar.SetBorder(true)
	page.AddItem(app.Sidebar, 1, 0, 1, 1, 0, 0, false)

	// Sidebar
	app.Statusbar = tview.NewTextView().
		SetChangedFunc(func() {
			app.Application.Draw()
		})
	page.AddItem(app.Statusbar, 2, 0, 1, 2, 0, 0, false)

	// Home page
	app.Pages.AddPage("home", page, true, true)

	// Set global keymaps
	app.SetKeymaps()

	// Add it to page and display
	app.Application.SetRoot(app.Pages, true).SetFocus(app.Content)

	return &app
}

// Quit, help, toggle
func (a *App) SetKeymaps() {
	a.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		// Exit
		case tcell.KeyEsc:
			a.Application.Stop()
			return nil
		}

		switch event.Rune() {
		// Basics
		case 'q':
			a.Application.Stop()
			return nil
		}
		return event
	})
}

func main() {
	// Get path from args
	app := NewApp()

	// Run the application
	err := app.Application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
