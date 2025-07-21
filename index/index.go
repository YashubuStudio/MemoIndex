package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"ykvario.com/MemoIndex/config"
)

// ファイル1件をインデックス登録
func IndexFile(absPath string) error {
	// インデックスパス設定
	indexPath := config.AppConfig.IndexPath
	if indexPath == "" {
		indexPath = "./memoindex.bleve"
	}

	// ファイル内容を読み込み
	body, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("ファイル読み込み失敗: %w", err)
	}

	// インデックスを開く（または作成）
	var idx bleve.Index
	idx, err = bleve.Open(indexPath)
	if err != nil {
		idx, err = bleve.New(indexPath, bleve.NewIndexMapping())
		if err != nil {
			return fmt.Errorf("インデックス作成失敗: %w", err)
		}
	}
	defer idx.Close()

	// インデックスキーを memoDirs からの相対パスで決定
	memoDirs := config.AppConfig.MemoDirs
	if len(memoDirs) == 0 {
		memoDirs = []string{"./memo"} // デフォルト
	}

	var relPath string
	found := false
	for _, dir := range memoDirs {
		rel, err := filepath.Rel(dir, absPath)
		if err == nil && !strings.HasPrefix(rel, "..") {
			relPath = filepath.ToSlash(filepath.Join(dir, rel)) // Unix風に統一
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("ファイルがmemo_dirsに含まれていません: %s", absPath)
	}

	doc := map[string]string{
		"body": string(body),
	}

	if err := idx.Index(relPath, doc); err != nil {
		return fmt.Errorf("インデックス登録失敗: %w", err)
	}

	fmt.Println("インデックス登録完了:", relPath)
	return nil
}
