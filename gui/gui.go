package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cobra"

	"ykvario.com/MemoIndex/note"
	"ykvario.com/MemoIndex/search"
)

// GuiCmd defines the CLI command to start the GUI.
var GuiCmd = &cobra.Command{
	Use:   "gui",
	Short: "GUIアプリを起動します",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

// Run launches the Fyne based GUI application.
func Run() {
	a := app.New()
	w := a.NewWindow("MemoIndex")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("検索ワード")

	resultBox := widget.NewMultiLineLabel("")
	resultBox.Wrapping = fyne.TextWrapWord

	searchButton := widget.NewButton("検索", func() {
		results, err := search.ExecuteSearch(entry.Text, 3)
		if err != nil {
			log.Println(err)
			return
		}
		if len(results) == 0 {
			resultBox.SetText("検索結果がありません。")
			return
		}
		text := ""
		for i, r := range results {
			text += fmt.Sprintf("%d. %s\n   ...%s...\n", i+1, r.Path, r.Fragment)
		}
		resultBox.SetText(text)
	})

	newButton := widget.NewButton("新規メモ", func() {
		path, err := note.CreateNewNote("")
		if err != nil {
			log.Println(err)
			resultBox.SetText(fmt.Sprintf("エラー: %v", err))
			return
		}
		resultBox.SetText(fmt.Sprintf("作成: %s", path))
	})

	control := container.NewHBox(entry, searchButton, newButton)
	content := container.NewVBox(control, resultBox)

	w.SetContent(content)
	w.ShowAndRun()
}
