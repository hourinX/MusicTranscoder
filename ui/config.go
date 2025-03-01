package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var sources []string

func InitConfig(src []string) {
	sources = src
}

func NewConfigContentUI(contentArea *fyne.Container) *widget.Button {
	configButton := widget.NewButton("配置", func() {
		// 创建配置表单项
		fileSourceEntry := widget.NewEntry()
		fileSourceEntry.SetPlaceHolder("请输入文件来源路径")

		outputPathEntry := widget.NewEntry()
		outputPathEntry.SetPlaceHolder("请输入输出路径")

		if len(sources) == 0 {
			sources = []string{"网易云音乐"}
		}
		// sourceNames := make([]string, len(sources))
		// for i, v := range sources {
		// 	sourceNames[i] = v.Name
		// }

		// sourceSelect := newIconSelect(musicSources)
		sourceSelect := widget.NewSelect(
			sources,
			func(selected string) {
				// 选择时的回调，可选逻辑
				println("选择来源:", selected)
			},
		)
		sourceSelect.SetSelected(sources[0])

		// 每组 Label 和 Input 使用 HBox 横向排列
		fileSourceRow := container.NewHBox(
			widget.NewLabel("文件来源:"),
			fileSourceEntry,
		)
		outputPathRow := container.NewHBox(
			widget.NewLabel("输出路径:"),
			outputPathEntry,
		)
		sourceSelectRow := container.NewHBox(
			widget.NewLabel("选择来源:"),
			sourceSelect,
		)

		// 保存按钮
		saveButton := widget.NewButton("保存", func() {
			// 这里可以获取输入框的值并处理保存逻辑
			fileSource := fileSourceEntry.Text
			outputPath := outputPathEntry.Text
			selectedSource := sourceSelect.Selected // 选择来源下拉框
			// 示例：打印输入值，实际可以替换为保存逻辑
			println("文件来源:", fileSource)
			println("输出路径:", outputPath)
			println("选择来源:", selectedSource)
		})

		// 将所有表单项垂直排列
		configForm := container.NewVBox(
			fileSourceRow,
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

// func newIconSelect(sourceMusic []model.MusicSource) *model.IconSelect {
// 	names := make([]string, len(sources))
// 	icons := make(map[string]*canvas.Image)

// 	for i, v := range sourceMusic {
// 		names[i] = v.Name
// 		img := canvas.NewImageFromFile(v.Icon)
// 		img.FillMode = canvas.ImageFillContain
// 		img.SetMinSize(fyne.NewSize(12, 12))
// 		icons[v.Name] = img
// 	}

// 	is := &model.IconSelect{
// 		Icons: icons,
// 	}
// 	is.ExtendBaseWidget(is)
// }
