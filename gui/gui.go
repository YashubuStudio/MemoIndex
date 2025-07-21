package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cobra"

	"ykvario.com/MemoIndex/i18n"
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
	entry.SetPlaceHolder(i18n.T("search_placeholder", nil))

	resultBox := widget.NewLabel("")
	resultBox.Wrapping = fyne.TextWrapWord

	searchButton := widget.NewButton(i18n.T("search", nil), func() {
		results, err := search.ExecuteSearch(entry.Text, 3)
		if err != nil {
			log.Println(err)
			return
		}
		if len(results) == 0 {
			resultBox.SetText(i18n.T("no_results", nil))
			return
		}
		text := ""
		for i, r := range results {
			text += fmt.Sprintf("%d. %s\n   ...%s...\n", i+1, r.Path, r.Fragment)
		}
		resultBox.SetText(text)
	})

	newButton := widget.NewButton(i18n.T("new_note", nil), func() {
		path, err := note.CreateNewNote("")
		if err != nil {
			log.Println(err)
			resultBox.SetText(i18n.T("error", map[string]interface{}{"Err": err}))
			return
		}
		resultBox.SetText(i18n.T("created", map[string]interface{}{"Path": path}))
	})

	control := container.NewHBox(entry, searchButton, newButton)
	content := container.NewVBox(control, resultBox)

	w.SetContent(content)
	w.ShowAndRun()
}
