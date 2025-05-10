package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-pay/sql2gorm-tool/model"
	"github.com/go-pay/sql2gorm-tool/parser"
	"github.com/go-pay/sql2gorm-tool/themes"
	"github.com/go-pay/xlog"
	"strings"
	"time"
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
		btnConvert   = widget.NewButtonWithIcon("转换DDL", theme.MediaPlayIcon(), nil)
		btnCopy      = widget.NewButtonWithIcon("复制结果", theme.ContentCopyIcon(), nil)
		btnReset     = widget.NewButtonWithIcon("清空所有", theme.ContentClearIcon(), nil)

		ops = options{
			Charset:        "",
			Collation:      "",
			JsonTag:        true,
			NoNullType:     true,
			NullStyle:      "",
			Package:        "model",
			GormType:       true,
			ForceTableName: true,
			Sql:            "",
		}
		ddlSQL string
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

	// 获取当前活动窗口
	win := fyne.CurrentApp().Driver().AllWindows()
	if len(win) == 0 {
		return
	}
	currentWin := win[0]

	// DDL Input
	ddlDataEntry = widget.NewMultiLineEntry()
	ddlDataEntry.Wrapping = fyne.TextWrapWord
	ddlDataEntry.SetMinRowsVisible(15)
	ddlDataEntry.SetPlaceHolder("请输入一个或多个 MySQL 建表 DDL 语句")
	ddlDataEntry.OnChanged = func(s string) {
		ddlSQL = s
	}
	//ddlLayout := container.NewWithoutLayout(ddlDataEntry)
	//ddlDataEntry.Move(fyne.NewPos(10, 0))
	//ddlDataEntry.Resize(fyne.NewSize(1175, 300))

	// BTN Convert
	centerBTN := container.NewCenter(container.NewHBox(btnConvert, widget.NewLabel("  "), btnCopy, widget.NewLabel("  "), btnReset))
	//btnLayout := container.NewWithoutLayout(centerBTN)
	//centerBTN.Move(fyne.NewPos(500, 240))

	// Struct Output
	structEntry = widget.NewMultiLineEntry()
	structEntry.Wrapping = fyne.TextWrapWord
	structEntry.SetMinRowsVisible(15)
	structEntry.SetPlaceHolder("GORM Struct 结果")
	//structLayout := container.NewWithoutLayout(structEntry)
	//structEntry.Move(fyne.NewPos(10, 220))
	//structEntry.Resize(fyne.NewSize(1175, 300))

	// 按钮事件
	btnConvert.OnTapped = func() {
		if ddlSQL == "" {
			ShowToast(currentWin, "请输入DDL语句", time.Second*2)
			return
		}
		buf := &strings.Builder{}
		data, err := parser.ParseSql(ddlSQL, getOptions(ops)...)
		if err != nil {
			ShowToast(currentWin, "转换失败，检查DDL后重试", time.Second*2)
			xlog.Errorf("转换失败，检查DDL后重试：%s", err.Error())
			return
		}
		for i, v := range data.StructCode {
			buf.WriteString(v)
			if i != len(data.StructCode)-1 {
				buf.WriteByte('\n')
			}
		}
		structEntry.SetText(buf.String())
		structEntry.Refresh()
	}
	// 复制结果
	btnCopy.OnTapped = func() {
		// 复制到剪切板
		if structEntry.Text == "" {
			return
		}
		// 复制文本到剪贴板
		clip := currentWin.Clipboard()
		if clip != nil {
			text := structEntry.SelectedText()
			if text != "" {
				clip.SetContent(text)
			} else {
				clip.SetContent(structEntry.Text)
			}
			ShowToast(currentWin, "复制成功", time.Second*2)
		} else {
			ShowToast(currentWin, "不能复制空内容", time.Second*2)
		}
	}
	// 重置
	btnReset.OnTapped = func() {
		ddlDataEntry.SetText("")
		structEntry.SetText("")
	}

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

// =====================================================================================================================

type options struct {
	Charset   string
	Collation string
	JsonTag   bool
	//TablePrefix    string
	//ColumnPrefix   string
	NoNullType     bool
	NullStyle      string
	Package        string
	GormType       bool
	ForceTableName bool

	Sql string

	MysqlDsn   string
	MysqlTable string
}

func getOptions(args options) []parser.Option {
	opt := make([]parser.Option, 0, 1)
	if args.Charset != "" {
		opt = append(opt, parser.WithCharset(args.Charset))
	}
	if args.Collation != "" {
		opt = append(opt, parser.WithCollation(args.Collation))
	}
	if args.JsonTag {
		opt = append(opt, parser.WithJsonTag())
	}
	//if args.TablePrefix != "" {
	//	opt = append(opt, parser.WithTablePrefix(args.TablePrefix))
	//}
	//if args.ColumnPrefix != "" {
	//	opt = append(opt, parser.WithColumnPrefix(args.ColumnPrefix))
	//}
	if args.NoNullType {
		opt = append(opt, parser.WithNoNullType())
	}
	if args.NullStyle != "" {
		switch args.NullStyle {
		case "sql":
			opt = append(opt, parser.WithNullStyle(parser.NullInSql))
		case "ptr":
			opt = append(opt, parser.WithNullStyle(parser.NullInPointer))
		default:
			fmt.Printf("invalid null style: %s\n", args.NullStyle)
			return nil
		}
	}
	if args.Package != "" {
		opt = append(opt, parser.WithPackage(args.Package))
	}
	if args.GormType {
		opt = append(opt, parser.WithGormType())
	}
	if args.ForceTableName {
		opt = append(opt, parser.WithForceTableName())
	}
	return opt
}

func ShowToast(win fyne.Window, message string, duration time.Duration) {
	if win == nil {
		// 获取当前活动窗口
		wins := fyne.CurrentApp().Driver().AllWindows()
		if len(wins) == 0 {
			return
		}
		// 复制文本到剪贴板
		win = wins[0]
	}
	// 创建一个 Label 作为 Toast 内容
	label := widget.NewLabel(message)
	// 创建一个 PopUp 作为 Toast
	toast := widget.NewPopUp(label, win.Canvas())
	//widget.ShowPopUpAtRelativePosition(label, win.Canvas(), fyne.NewPos(6, 9), nextTo)
	toast.Move(fyne.NewPos(502, 300))
	toast.Show()
	// 设置定时器，duration 后隐藏 Toast

	go func() {
		time.Sleep(duration)
		fyne.DoAndWait(func() {
			toast.Hide()
		})
	}()
}
