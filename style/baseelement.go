package core

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

// BaseElement renders a styled primitive element.
type BaseElement struct {
	Token  string
	Prefix string
	Suffix string
	Style  StylePrimitive
}

// Get tcell.Style with BackgroundColor, Foreground, Bold, ...
func GetStyle(rules StylePrimitive) tcell.Style {
	style := tcell.StyleDefault
	if rules.Color != nil {
		style.Foreground(tcell.GetColor(*rules.Color))
	}
	if rules.BackgroundColor != nil {
		style.Background(tcell.GetColor(*rules.BackgroundColor))
	}
	if rules.Underline != nil && *rules.Underline {
		style.Underline(true)
	}
	if rules.Bold != nil && *rules.Bold {
		style.Bold(true)
	}
	if rules.Italic != nil && *rules.Italic {
		style.Italic(true)
	}
	if rules.CrossedOut != nil && *rules.CrossedOut {
		style.StrikeThrough(true)
	}
	if rules.Overlined != nil && *rules.Overlined {
		// No implementation in tcell
	}
	if rules.Inverse != nil && *rules.Inverse {
		style.Reverse(true)
	}
	if rules.Blink != nil && *rules.Blink {
		style.Blink(true)
	}
	return style
}

// Decorate string with prefix, suffix, case, ...
func DecorateString(rules StylePrimitive, s string) string {
	if rules.Upper != nil && *rules.Upper {
		s = strings.ToUpper(s)
	}
	if rules.Lower != nil && *rules.Lower {
		s = strings.ToLower(s)
	}
	if rules.Title != nil && *rules.Title {
		s = strings.Title(s)
	}
	return s
}
