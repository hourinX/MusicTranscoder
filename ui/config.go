package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var sources []string

func InitConfig(src []string) {
	sources = src
}

func NewConfigContentUI(contentArea *fyne.Container) *widget.Button {
	configButton := widget.NewButton("配置", func() {
		windows := fyne.CurrentApp().Driver().AllWindows()[0]
		// 创建配置表单项
		fileSourceEntry := widget.NewEntry()
		fileSourceEntry.SetPlaceHolder("请输入文件来源路径")

		outputPathEntry := widget.NewEntry()
		outputPathEntry.SetPlaceHolder("请输入输出路径")
		outputPathEntry.Disable()
		outputPathEntryContainer := container.New(
			layout.NewGridWrapLayout(fyne.NewSize(500, 36)),
			outputPathEntry,
		)

		selectFolderButton := widget.NewButton("选择", func() {
			folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
				if err == nil && uri != nil {
					outputPathEntry.SetText(uri.Path())
				}
			}, windows)
			folderDialog.Show()
		})

		if len(sources) == 0 {
			sources = []string{"网易云音乐"}
		}

		sourceSelect := widget.NewSelect(
			sources,
			func(selected string) {
				// 选择时的回调，可选逻辑
				println("选择来源:", selected)
			},
		)
		sourceSelect.SetSelected(sources[0])

		outputPathRow := container.NewHBox(
			widget.NewLabel("输出路径:"),
			outputPathEntryContainer,
			selectFolderButton,
		)
		sourceSelectRow := container.NewHBox(
			widget.NewLabel("选择来源:"),
			sourceSelect,
		)

		// 保存按钮
		saveButton := widget.NewButton("保存", func() {
			// 这里可以获取输入框的值并处理保存逻辑
			outputPath := outputPathEntry.Text
			selectedSource := sourceSelect.Selected // 选择来源下拉框
			// 示例：打印输入值，实际可以替换为保存逻辑
			println("输出路径:", outputPath)
			println("选择来源:", selectedSource)
		})

		// 将所有表单项垂直排列
		configForm := container.NewVBox(
			// fileSourceRow,
			outputPathRow,
			sourceSelectRow,
			saveButton,
		)

		// 更新右侧内容区域
		contentArea.Objects = []fyne.CanvasObject{
			configForm,
		}
		contentArea.Refresh()
	})
	return configButton
}
