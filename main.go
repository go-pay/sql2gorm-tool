package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-pay/sql2gorm-tool/model"
	"github.com/go-pay/sql2gorm-tool/themes"
	"github.com/go-pay/xlog"
)

func main() {
	xlog.SetLevel(xlog.WarnLevel)
	a := app.NewWithID("com.jerry.sql2gorm")
	// 监听前台焦点事件
	//logLifecycle(a)
	// 设置中文主题
	a.Settings().SetTheme(&themes.ChineseTheme{Variant: theme.VariantDark})
	// 新建窗口
	w := a.NewWindow("SQL TO GORM 工具")
	// tool 菜单
	//w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	var (
		ddlDataEntry *widget.Entry
		structEntry  *widget.Entry
		btnConvert   = widget.NewButtonWithIcon("转换", theme.MediaPlayIcon(), nil)
		btnCopy      = widget.NewButtonWithIcon("复制结果", theme.ContentCopyIcon(), nil)
	)

	devLabel := widget.NewLabel(model.Copyright)
	devLabel.Importance = widget.SuccessImportance

	// 主题选择
	darkOrLight := widget.NewSelect([]string{"黑色主题", "白色主题"}, nil)
	darkOrLight.SetSelected("黑色主题")
	darkOrLight.OnChanged = func(s string) {
		switch s {
		case "黑色主题":
			a.Settings().SetTheme(&themes.ChineseTheme{Variant: theme.VariantDark})
			devLabel.Importance = widget.SuccessImportance
			devLabel.Refresh()
		case "白色主题":
			a.Settings().SetTheme(&themes.ChineseTheme{Variant: theme.VariantLight})
			devLabel.Importance = widget.DangerImportance
			devLabel.Refresh()
		}
	}
	themeChange := container.NewBorder(nil, nil, nil, darkOrLight, container.NewCenter(widget.NewLabel("SQL 转 GORM Struct")))

	// DDL Input
	ddlDataEntry = widget.NewMultiLineEntry()
	ddlDataEntry.Wrapping = fyne.TextWrapWord
	ddlDataEntry.SetMinRowsVisible(15)
	ddlDataEntry.SetPlaceHolder("create table ddl sql")
	//ddlLayout := container.NewWithoutLayout(ddlDataEntry)
	//ddlDataEntry.Move(fyne.NewPos(10, 0))
	//ddlDataEntry.Resize(fyne.NewSize(1175, 300))

	// BTN Convert
	centerBTN := container.NewCenter(container.NewHBox(btnConvert, widget.NewLabel("     "), btnCopy))
	//btnLayout := container.NewWithoutLayout(centerBTN)
	//centerBTN.Move(fyne.NewPos(500, 240))

	// Struct Output
	structEntry = widget.NewMultiLineEntry()
	structEntry.Wrapping = fyne.TextWrapWord
	structEntry.SetMinRowsVisible(15)
	structEntry.SetPlaceHolder("struct result")
	//structLayout := container.NewWithoutLayout(structEntry)
	//structEntry.Move(fyne.NewPos(10, 220))
	//structEntry.Resize(fyne.NewSize(1175, 300))

	w.SetContent(
		container.NewBorder(
			container.NewVBox(themeChange, widget.NewSeparator()),
			container.NewVBox(widget.NewSeparator(), container.NewCenter(devLabel)), nil, nil,
			container.NewVBox(
				ddlDataEntry,
				centerBTN,
				structEntry,
			),
		),
	)

	w.Resize(fyne.NewSize(1200, 760))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.ShowAndRun()
}
