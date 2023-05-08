package custometheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

/*
var (
	purple = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	orange = &color.NRGBA{R: 198, G: 123, B: 0, A: 255}
	grey   = &color.Gray{Y: 123}
)*/

// customTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type customTheme struct {
}

func (customTheme) Color(c fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {

	return theme.DarkTheme().Color(c, theme.VariantDark)

}

func (customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DarkTheme().Font(style)
}

func (customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (customTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameScrollBar:
		return 10
	case theme.SizeNameScrollBarSmall:
		return 5
	case theme.SizeNameText:
		return 22
	case theme.SizeNameHeadingText:
		return 30
	case theme.SizeNameSubHeadingText:
		return 25
	case theme.SizeNameCaptionText:
		return 15
	case theme.SizeNameInputBorder:
		return 1
	default:
		return 0
	}
}

func NewCustomTheme() fyne.Theme {
	return &customTheme{}
}
