package note

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"ykvario.com/MemoIndex/config"
	"ykvario.com/MemoIndex/index"
)

var NewNoteCmd = &cobra.Command{
	Use:   "new [filename]",
	Short: "新しいメモを作成します",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := ""
		if len(args) > 0 {
			filename = args[0]
		} else {
			filename = fmt.Sprintf("memo_%s.txt", time.Now().Format("20060102_150405"))
		}

		// メモディレクトリ構成
		memoDirs := config.AppConfig.MemoDirs
		memoDir := "./memo"
		if len(memoDirs) > 0 {
			memoDir = memoDirs[0]
		}
		err := os.MkdirAll(memoDir, os.ModePerm)
		if err != nil {
			log.Fatalf("メモディレクトリ作成失敗: %v", err)
		}

		filepathAbs := filepath.Join(memoDir, filename)

		// ファイル作成
		err = os.WriteFile(filepathAbs, []byte(""), 0644)
		if err != nil {
			log.Fatalf("ファイル作成に失敗: %v", err)
		}
		fmt.Println("作成:", filepathAbs)

		// エディタ起動
		editor := config.AppConfig.Editor
		if editor == "" {
			editor = os.Getenv("EDITOR")
		}
		if editor == "" {
			editor = "notepad"
		}

		cmdEditor := exec.Command(editor, filepathAbs)
		cmdEditor.Stdin = os.Stdin
		cmdEditor.Stdout = os.Stdout
		cmdEditor.Stderr = os.Stderr
		cmdEditor.Run()

		// インデックス登録（外部関数へ委譲）
		err = index.IndexFile(filepathAbs)
		if err != nil {
			log.Fatalf("インデックス登録失敗: %v", err)
		}
	},
}
