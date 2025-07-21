package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MemoDir   string `yaml:"memo_dir"`
	IndexPath string `yaml:"index_path"`
	Editor    string `yaml:"editor"`
}

var AppConfig Config

func LoadConfig(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 設定ファイルが無ければデフォルト設定で生成
		defaultCfg := Config{
			MemoDir:   "./memo",
			IndexPath: "./memoindex.bleve",
			Editor:    "notepad",
		}
		data, err := yaml.Marshal(&defaultCfg)
		if err != nil {
			log.Fatalf("設定ファイル生成失敗: %v", err)
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			log.Fatalf("設定ファイル生成失敗: %v", err)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("設定ファイル読み込み失敗: %v", err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("設定パース失敗: %v", err)
	}

	// 設定されたパスを絶対パスへ展開
	if AppConfig.MemoDir != "" && !filepath.IsAbs(AppConfig.MemoDir) {
		if abs, err := filepath.Abs(AppConfig.MemoDir); err == nil {
			AppConfig.MemoDir = abs
		}
	}
	if AppConfig.IndexPath != "" && !filepath.IsAbs(AppConfig.IndexPath) {
		if abs, err := filepath.Abs(AppConfig.IndexPath); err == nil {
			AppConfig.IndexPath = abs
		}
	}
}
