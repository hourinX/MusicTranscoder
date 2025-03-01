package ui

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var iconData []byte

func StartForm() {
	app := app.New()
	form := app.NewWindow("MusicTranscoder")
	form.Resize(fyne.NewSize(900, 600))

	// 右侧内容区域
	contentArea := container.NewMax(
		widget.NewLabel("欢迎来到 MusicTranscoder 主页！"), // 默认显示主页内容
	)

	// 主页按钮
	homeButton := widget.NewButton("主页", func() {
		contentArea.Objects = []fyne.CanvasObject{
			widget.NewLabel("欢迎来到 MusicTranscoder 主页！"),
		}
		contentArea.Refresh()
	})

	// 配置按钮
	configButton := NewConfigContentUI(contentArea)

	// 关于按钮
	aboutButton := widget.NewButton("关于", func() {
		contentArea.Objects = []fyne.CanvasObject{
			widget.NewLabel("MusicTranscoder v1.0\n一款音乐转码工具"),
		}
		contentArea.Refresh()
	})
	// 退出按钮
	quitButton := widget.NewButton("退出", func() {
		app.Quit()
	})
	icon := canvas.NewImageFromFile("../static/icon.png")
	icon.FillMode = canvas.ImageFillContain // 保持图片比例
	icon.SetMinSize(fyne.NewSize(24, 24))

	titleContainer := container.NewHBox(
		icon,
		widget.NewLabel("音乐格式转换器"),
	)

	// 侧边导航栏
	sidebar := container.NewVBox(
		titleContainer,
		homeButton,
		configButton,
		aboutButton,
		widget.NewSeparator(),
		quitButton,
	)
	sidebarContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 700)), sidebar)

	// 创建主布局
	mainLayout := container.NewHBox(
		sidebarContainer, // 左侧固定宽度导航栏
		contentArea,      // 右侧内容区域自适应
	)

	// // 设置窗口内容并显示
	form.SetContent(mainLayout)
	form.ShowAndRun()
	// 定义导航项

}
