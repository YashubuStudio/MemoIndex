package main

import (
	"fmt"
	"os"

	// For Windows API calls
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal" // Still useful for general terminal detection
	"ykvario.com/MemoIndex/cfgcmd"
	"ykvario.com/MemoIndex/config"
	"ykvario.com/MemoIndex/gui"
	"ykvario.com/MemoIndex/i18n"
	"ykvario.com/MemoIndex/index"
	"ykvario.com/MemoIndex/note"
	"ykvario.com/MemoIndex/search"
)

// isTerminal は、プログラムが対話的なターミナルで実行されているかを判定します。
func isTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

func main() {
	// Load configurations and i18n first
	config.LoadConfig("config.yaml")
	if err := i18n.Load(config.AppConfig.Language); err != nil {
		fmt.Fprintln(os.Stderr, err) // Print to stderr for errors
		// Consider if you want to exit here or proceed with default language
	}

	// Check if we should directly launch GUI.
	// This takes precedence over Cobra's parsing if no CLI arguments are given.
	if len(os.Args) == 1 && !isTerminal() {
		// Only hide console on Windows when directly launching GUI
		hideConsole()
		gui.Run()
		return // Exit after GUI run
	}

	// Initialize Cobra for CLI commands
	rootCmd := &cobra.Command{
		Use:   "memoindex",
		Short: "MemoIndex - メモ検索＆新規作成CLIツール",
		Long:  `MemoIndexは、メモの検索、新規作成、インデックス再構築などを行えるツールです。GUIまたはCLIとして利用できます。`,
		// The Run function for the root command will only execute if
		// no specific subcommand is provided AND it's run from a terminal
		// or if arguments are present but don't match a command.
		Run: func(cmd *cobra.Command, args []string) {
			// If we reach here and it's a terminal, but no specific command was given,
			// or if an invalid command was given, show help.
			_ = cmd.Help()
		},
		SilenceUsage:  true, // Suppress usage on error
		SilenceErrors: true, // Suppress Cobra's default error message
	}

	// Add all your subcommands
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(note.NewNoteCmd)
	rootCmd.AddCommand(index.ReindexCmd)
	rootCmd.AddCommand(gui.GuiCmd) // Keep this for explicit 'memoindex gui'
	rootCmd.AddCommand(cfgcmd.Cmd)

	// Execute the root command. This will parse arguments and call the appropriate command's Run function.
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err) // Print error to stderr
		os.Exit(1)
	}
}
