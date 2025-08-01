package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MemoDirs  []string `yaml:"memo_dirs"` // 複数対応に変更
	IndexPath string   `yaml:"index_path"`
	Editor    string   `yaml:"editor"`
	Language  string   `yaml:"language"`
}

var AppConfig Config

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			sample := path + ".sample"
			if sampleData, sErr := os.ReadFile(sample); sErr == nil {
				if wErr := os.WriteFile(path, sampleData, 0644); wErr == nil {
					data = sampleData
				} else {
					log.Fatalf("設定ファイル作成失敗: %v", wErr)
				}
			} else {
				defaultCfg := Config{
					MemoDirs:  []string{"./memo"},
					IndexPath: "./memoindex.bleve",
					Editor:    "notepad",
					Language:  "ja",
				}
				bytes, mErr := yaml.Marshal(&defaultCfg)
				if mErr != nil {
					log.Fatalf("設定初期化失敗: %v", mErr)
				}
				if wErr := os.WriteFile(path, bytes, 0644); wErr != nil {
					log.Fatalf("設定ファイル作成失敗: %v", wErr)
				}
				data = bytes
			}
		} else {
			log.Fatalf("設定ファイル読み込み失敗: %v", err)
		}
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("設定パース失敗: %v", err)
	}
}

func SaveConfig(path string) error {
	data, err := yaml.Marshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("設定シリアライズ失敗: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("設定ファイル書き込み失敗: %w", err)
	}
	return nil
}
