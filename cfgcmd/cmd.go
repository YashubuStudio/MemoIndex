package cfgcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"ykvario.com/MemoIndex/config"
	"ykvario.com/MemoIndex/i18n"
)

// Cmd is the root command for configuration operations.
var Cmd = &cobra.Command{
	Use:   "config",
	Short: "アプリ設定を変更します",
}

var langCmd = &cobra.Command{
	Use:   "lang [locale]",
	Short: "使用言語を設定します",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lang := args[0]
		config.AppConfig.Language = lang
		if err := config.SaveConfig("config.yaml"); err != nil {
			return err
		}
		if err := i18n.Load(lang); err != nil {
			return err
		}
		fmt.Println(i18n.T("language_set", map[string]interface{}{"Lang": lang}))
		return nil
	},
}

var addDirCmd = &cobra.Command{
	Use:   "add-dir [path]",
	Short: "インデックス対象ディレクトリを追加します",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		for _, d := range config.AppConfig.MemoDirs {
			if d == dir {
				fmt.Println(i18n.T("already_exists", map[string]interface{}{"Dir": dir}))
				return nil
			}
		}
		config.AppConfig.MemoDirs = append(config.AppConfig.MemoDirs, dir)
		if err := config.SaveConfig("config.yaml"); err != nil {
			return err
		}
		fmt.Println(i18n.T("dir_added", map[string]interface{}{"Dir": dir}))
		return nil
	},
}

var removeDirCmd = &cobra.Command{
	Use:   "remove-dir [path]",
	Short: "インデックス対象ディレクトリを削除します",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		newDirs := []string{}
		found := false
		for _, d := range config.AppConfig.MemoDirs {
			if d == dir {
				found = true
				continue
			}
			newDirs = append(newDirs, d)
		}
		if !found {
			fmt.Println(i18n.T("dir_not_found", map[string]interface{}{"Dir": dir}))
			return nil
		}
		config.AppConfig.MemoDirs = newDirs
		if err := config.SaveConfig("config.yaml"); err != nil {
			return err
		}
		fmt.Println(i18n.T("dir_removed", map[string]interface{}{"Dir": dir}))
		return nil
	},
}

func init() {
	Cmd.AddCommand(langCmd)
	Cmd.AddCommand(addDirCmd)
	Cmd.AddCommand(removeDirCmd)
}
