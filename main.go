package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/index"
	"ykvario.com/MemoIndex/note"
	"ykvario.com/MemoIndex/search"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "memoindex",
		Short: "MemoIndex - メモ検索＆新規作成CLIツール",
	}

	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(note.NewNoteCmd)
	rootCmd.AddCommand(index.ReindexCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
