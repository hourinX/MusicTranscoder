package main

import (
	"MusicTranscoder/ui"
	"strings"

	"github.com/spf13/viper"
)

func read() {
	viper.SetConfigName("conf")
	viper.SetConfigType("ini")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件无法加载")
	}

	sourceStr := viper.GetString("origin.musicSource")
	if sourceStr == "" {
		sourceStr = "网易云音乐"
	}
	sources := strings.Split(strings.TrimSpace(sourceStr), ",")
	for item, index := range sources {
		sources[item] = strings.TrimSpace(index)
	}
	// iconStr := viper.GetString("origin.musicIcon")
	// if iconStr == "" {
	// 	iconStr = "../static/cloud.png"
	// }
	// icons := strings.Split(strings.TrimSpace(iconStr), ",")
	// for item, index := range icons {
	// 	icons[item] = strings.TrimSpace(index)
	// }

	// if len(sources) != len(icons) {
	// 	panic("icon数量异常")
	// }

	// musicSource := make([]model.MusicSource, len(sources))
	// for i := range sources {
	// 	musicSource[i] = model.MusicSource{
	// 		Icon: sources[i],
	// 		Name: icons[i],
	// 	}
	// }
	ui.InitConfig(sources)
}
