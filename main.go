package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/cfgcmd"
	"ykvario.com/MemoIndex/config"
	"ykvario.com/MemoIndex/gui"
	"ykvario.com/MemoIndex/i18n"
	"ykvario.com/MemoIndex/index"
	"ykvario.com/MemoIndex/note"
	"ykvario.com/MemoIndex/search"
)

func main() {
	// 設定ファイル読み込み
	config.LoadConfig("config.yaml")
	// i18n ロード
	if err := i18n.Load(config.AppConfig.Language); err != nil {
		fmt.Println(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "memoindex",
		Short: "MemoIndex - メモ検索＆新規作成CLIツール",
	}

	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(note.NewNoteCmd)
	rootCmd.AddCommand(index.ReindexCmd)
	rootCmd.AddCommand(gui.GuiCmd)
	rootCmd.AddCommand(cfgcmd.Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
