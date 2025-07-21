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
		}

		if _, err := CreateNewNote(filename); err != nil {
			log.Fatalf("%v", err)
		}
	},
}

// CreateNewNote creates a new memo file, opens it with the configured
// editor and indexes the file. The created absolute path is returned.
func CreateNewNote(filename string) (string, error) {
	if filename == "" {
		filename = fmt.Sprintf("memo_%s.txt", time.Now().Format("20060102_150405"))
	}

	memoDirs := config.AppConfig.MemoDirs
	memoDir := "./memo"
	if len(memoDirs) > 0 {
		memoDir = memoDirs[0]
	}
	if err := os.MkdirAll(memoDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("メモディレクトリ作成失敗: %w", err)
	}

	filepathAbs := filepath.Join(memoDir, filename)
	if err := os.WriteFile(filepathAbs, []byte(""), 0644); err != nil {
		return "", fmt.Errorf("ファイル作成に失敗: %w", err)
	}
	fmt.Println("作成:", filepathAbs)

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
	if err := cmdEditor.Run(); err != nil {
		return "", fmt.Errorf("エディタ起動失敗: %w", err)
	}

	if err := index.IndexFile(filepathAbs); err != nil {
		return "", fmt.Errorf("インデックス登録失敗: %w", err)
	}

	return filepathAbs, nil
}

// OpenFile opens the specified file with the configured editor.
func OpenFile(path string) error {
	editor := config.AppConfig.Editor
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "notepad"
	}
	cmdEditor := exec.Command(editor, path)
	return cmdEditor.Start()
}
