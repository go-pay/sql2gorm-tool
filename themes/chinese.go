package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/go-pay/sql2gorm-tool/font"
)

// =============================================主题=====================================================================

var (
	purple       = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	orange       = &color.NRGBA{R: 198, G: 123, B: 0, A: 255}
	grey         = &color.Gray{Y: 123}
	errorColor   = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	successColor = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	warningColor = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
)

// 自定义主题
type ChineseTheme struct {
	Variant fyne.ThemeVariant
}

func (m *ChineseTheme) Font(style fyne.TextStyle) fyne.Resource {
	return font.ResourceNotoSansSCRegularTtf
}

func (m *ChineseTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	primary := fyne.CurrentApp().Settings().PrimaryColor()
	if n == theme.ColorNamePrimary || n == theme.ColorNameHyperlink {
		return primaryColorNamed(primary)
	} else if n == theme.ColorNameFocus {
		return focusColorNamed(primary)
	} else if n == theme.ColorNameSelection {
		return selectionColorNamed(primary)
	}
	if m.Variant == theme.VariantLight {
		return lightPaletColorNamed(n)
	}
	return darkPaletColorNamed(n)
}

func (m *ChineseTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *ChineseTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameInnerPadding:
		return 8
	case theme.SizeNameLineSpacing:
		return 4
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameText:
		return 14
	case theme.SizeNameHeadingText:
		return 24
	case theme.SizeNameSubHeadingText:
		return 18
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInputBorder:
		return 1
	case theme.SizeNameInputRadius:
		return 5
	case theme.SizeNameSelectionRadius:
		return 3
	default:
		return 0
	}
}

func primaryColorNamed(name string) color.NRGBA {
	switch name {
	case theme.ColorRed:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorOrange:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	case theme.ColorYellow:
		return color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0xff}
	case theme.ColorGreen:
		return color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0xff}
	case theme.ColorPurple:
		return color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}
	case theme.ColorBrown:
		return color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0xff}
	case theme.ColorGray:
		return color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0xff}
	}
	// We return the value for ColorBlue for every other value.
	// There is no need to have it in the switch above.
	return color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
}

func selectionColorNamed(name string) color.NRGBA {
	switch name {
	case theme.ColorRed:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x3f}
	case theme.ColorOrange:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x3f}
	case theme.ColorYellow:
		return color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x3f}
	case theme.ColorGreen:
		return color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x3f}
	case theme.ColorPurple:
		return color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x3f}
	case theme.ColorBrown:
		return color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x3f}
	case theme.ColorGray:
		return color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x3f}
	}
	// We return the value for ColorBlue for every other value.
	// There is no need to have it in the switch above.
	return color.NRGBA{R: 0x00, G: 0x6C, B: 0xff, A: 0x40}
}

func focusColorNamed(name string) color.NRGBA {
	switch name {
	case theme.ColorRed:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x7f}
	case theme.ColorOrange:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x7f}
	case theme.ColorYellow:
		return color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x7f}
	case theme.ColorGreen:
		return color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x7f}
	case theme.ColorPurple:
		return color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x7f}
	case theme.ColorBrown:
		return color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x7f}
	case theme.ColorGray:
		return color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x7f}
	}
	// We return the value for ColorBlue for every other value.
	// There is no need to have it in the switch above.
	return color.NRGBA{R: 0x00, G: 0x6C, B: 0xff, A: 0x2a}
}

func darkPaletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 40, G: 40, B: 40, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 64, G: 131, B: 201, A: 0xff}
	case theme.ColorNameDisabled:
		//return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
		return color.NRGBA{R: 150, G: 150, B: 150, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 64, G: 131, B: 201, A: 95}
	case theme.ColorNameError:
		return errorColor
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	case theme.ColorNameHeaderBackground:
		return color.NRGBA{R: 0x1b, G: 0x1b, B: 0x1b, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 35, G: 35, B: 35, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 30, G: 30, B: 30, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x66}
	case theme.ColorNameSuccess:
		return successColor
	case theme.ColorNameWarning:
		return warningColor
	}
	return color.Transparent
}

func lightPaletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 252, G: 120, B: 54, A: 0xff}
	case theme.ColorNameDisabled:
		//return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
		return color.NRGBA{R: 150, G: 150, B: 150, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 252, G: 120, B: 54, A: 95}
	case theme.ColorNameError:
		return errorColor
	case theme.ColorNameForeground:
		return color.NRGBA{R: 40, G: 40, B: 40, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{A: 0x0f}
	case theme.ColorNameHeaderBackground:
		return color.NRGBA{R: 0xf9, G: 0xf9, B: 0xf9, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{A: 0x19}
	case theme.ColorNameScrollBar:
		return color.NRGBA{A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x33}
	case theme.ColorNameSuccess:
		return successColor
	case theme.ColorNameWarning:
		return warningColor
	}
	return color.Transparent
}
