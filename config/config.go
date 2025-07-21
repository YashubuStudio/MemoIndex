package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MemoDirs  []string `yaml:"memo_dirs"` // 複数対応に変更
	IndexPath string   `yaml:"index_path"`
	Editor    string   `yaml:"editor"`
}

var AppConfig Config

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("設定ファイル読み込み失敗: %v", err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("設定パース失敗: %v", err)
	}
}
