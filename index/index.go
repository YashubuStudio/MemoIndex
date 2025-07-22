package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"ykvario.com/MemoIndex/config"
)

// ファイル1件をインデックス登録（Kagome分かち書き付き）
func IndexFile(absPath string) error {
	indexPath := config.AppConfig.IndexPath
	if indexPath == "" {
		indexPath = "./memoindex.bleve"
	}
	os.MkdirAll(filepath.Dir(indexPath), os.ModePerm)

	bodyBytes, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("ファイル読み込み失敗: %w", err)
	}

	wakatiText, err := Wakati(string(bodyBytes))
	if err != nil {
		return fmt.Errorf("形態素解析失敗: %w", err)
	}
	//fmt.Println("【分かち書き】", wakatiText)

	// インデックスオープン（なければ keyword アナライザで作成）
	idx, err := bleve.Open(indexPath)
	if err != nil {
		idx, err = bleve.New(indexPath, CreateKeywordIndexMapping())
		//idx, err = bleve.New(indexPath, createSimpleIndexMapping())
		if err != nil {
			return fmt.Errorf("インデックス作成失敗: %w", err)
		}
	}
	defer idx.Close()

	// memoDirs から相対パスを生成（IDとして使う）
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

	//fmt.Println("【登録ファイル】", absPath)
	//fmt.Println("【登録キー】", relPath)
	//fmt.Println("【インデックス登録開始】")

	// 登録データ：分かち書き文字列をスペース分割して slice 化
	tokens := strings.Fields(wakatiText)
	doc := map[string]interface{}{
		"body": tokens, // keyword アナライザで各要素がそのまま1トークンになる
	}

	if err := idx.Index(relPath, doc); err != nil {
		return fmt.Errorf("インデックス登録失敗: %w", err)
	}

	fmt.Println("インデックス登録完了:", relPath)
	return nil
}
