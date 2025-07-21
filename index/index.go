package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"ykvario.com/MemoIndex/config"
)

// ファイル1件をインデックス登録（MeCab分かち書き付き）
func IndexFile(absPath string) error {
	indexPath := config.AppConfig.IndexPath
	if indexPath == "" {
		indexPath = "./memoindex.bleve"
	}

	bodyBytes, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("ファイル読み込み失敗: %w", err)
	}

	wakatiText, err := Wakati(string(bodyBytes))
	if err != nil {
		return fmt.Errorf("形態素解析失敗: %w", err)
	}

	idx, err := bleve.Open(indexPath)
	if err != nil {
		idx, err = bleve.New(indexPath, bleve.NewIndexMapping())
		if err != nil {
			return fmt.Errorf("インデックス作成失敗: %w", err)
		}
	}
	defer idx.Close()

	memoDirs := config.AppConfig.MemoDirs
	if len(memoDirs) == 0 {
		memoDirs = []string{"./memo"}
	}

	var relPath string
	found := false
	for _, dir := range memoDirs {
		rel, err := filepath.Rel(dir, absPath)
		if err == nil && !strings.HasPrefix(rel, "..") {
			relPath = filepath.ToSlash(filepath.Join(dir, rel))
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("ファイルがmemo_dirsに含まれていません: %s", absPath)
	}

	doc := map[string]string{
		"body": wakatiText,
	}

	if err := idx.Index(relPath, doc); err != nil {
		return fmt.Errorf("インデックス登録失敗: %w", err)
	}

	fmt.Println("インデックス登録完了:", relPath)
	return nil
}
