package index

import (
	"fmt"
	"os"
	"path/filepath"

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

	// インデックスキーを相対パスにする（memoDirからの）
	memoDir := config.AppConfig.MemoDir
	if memoDir == "" {
		memoDir = "./memo"
	}
	relPath, _ := filepath.Rel(memoDir, absPath)
	doc := map[string]string{"body": string(body)}

	// 登録
	if err := idx.Index(filepath.Join(memoDir, relPath), doc); err != nil {
		return fmt.Errorf("インデックス登録失敗: %w", err)
	}

	fmt.Println("インデックス登録完了:", relPath)
	return nil
}
