package index

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/config"
)

// コマンド：再インデックス化
var ReindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "メモフォルダ内のすべてのファイルを再インデックス化します",
	Run: func(cmd *cobra.Command, args []string) {
		count, err := ReindexAll()
		if err != nil {
			log.Printf("再インデックス中にエラー: %v", err)
		}
		fmt.Printf("再インデックス完了: %d 件登録\n", count)
	},
}

// ReindexAll reindexes all files under configured memo directories.
// The number of indexed files is returned.
func ReindexAll() (int, error) {
	memoDirs := config.AppConfig.MemoDirs
	if len(memoDirs) == 0 {
		memoDirs = []string{"./memo"} // デフォルト
	}

	count := 0
	for _, memoDir := range memoDirs {
		if _, err := os.Stat(memoDir); os.IsNotExist(err) {
			if mkErr := os.MkdirAll(memoDir, os.ModePerm); mkErr != nil {
				log.Printf("ディレクトリ作成失敗: %v", mkErr)
				continue
			}
			log.Printf("メモフォルダを作成しました: %s", memoDir)
		}

		err := filepath.WalkDir(memoDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if ext != ".txt" && ext != ".md" && ext != ".html" {
				return nil
			}

			if err := IndexFile(path); err != nil {
				log.Printf("インデックス登録失敗: %v", err)
			} else {
				count++
			}
			return nil
		})
		if err != nil {
			return count, err
		}
	}

	return count, nil
}
