package gui

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cobra"

	"ykvario.com/MemoIndex/i18n"
	"ykvario.com/MemoIndex/index"
	"ykvario.com/MemoIndex/note"
	"ykvario.com/MemoIndex/search"
)

type centerPercentLayout struct{ percent float32 }

func (l *centerPercentLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}
	min := objects[0].MinSize()
	return fyne.NewSize(min.Width, min.Height)
}

func (l *centerPercentLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) == 0 {
		return
	}
	obj := objects[0]
	w := size.Width * l.percent / 100
	h := obj.MinSize().Height
	x := (size.Width - w) / 2
	obj.Resize(fyne.NewSize(w, h))
	obj.Move(fyne.NewPos(x, 0))
}

type buttonRowLayout struct{ buttonWidth float32 }

func (l *buttonRowLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	h := float32(0)
	for _, o := range objects {
		if ms := o.MinSize().Height; ms > h {
			h = ms
		}
	}
	return fyne.NewSize(l.buttonWidth*float32(len(objects)), h)
}

func (l *buttonRowLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	count := len(objects)
	if count == 0 {
		return
	}
	spacing := (size.Width - float32(count)*l.buttonWidth) / float32(count+1)
	if spacing < 0 {
		spacing = 0
	}
	for i, o := range objects {
		h := o.MinSize().Height
		x := spacing + float32(i)*(l.buttonWidth+spacing)
		o.Resize(fyne.NewSize(l.buttonWidth, h))
		o.Move(fyne.NewPos(x, 0))
	}
}

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
	w.Resize(fyne.NewSize(640, 360)) // 初期ウィンドウサイズを設定

	entry := widget.NewEntry()
	entry.SetPlaceHolder(i18n.T("input_field", nil))

	resultBox := widget.NewLabel("")
	resultBox.Wrapping = fyne.TextWrapBreak

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

	reindexButton := widget.NewButton(i18n.T("reindex", nil), func() {
		count, err := index.ReindexAll()
		if err != nil {
			log.Println(err)
			resultBox.SetText(i18n.T("error", map[string]interface{}{"Err": err}))
			return
		}
		resultBox.SetText(i18n.T("reindex_done", map[string]interface{}{"Count": count}))
	})

	newButton := widget.NewButton(i18n.T("new_note", nil), func() {
		filename := ""
		title := strings.TrimSpace(entry.Text)
		if title != "" {
			if !strings.HasSuffix(title, ".txt") {
				filename = title + ".txt"
			} else {
				filename = title
			}
		}
		path, err := note.CreateNewNote(filename)
		if err != nil {
			log.Println(err)
			resultBox.SetText(i18n.T("error", map[string]interface{}{"Err": err}))
			return
		}
		resultBox.SetText(i18n.T("created", map[string]interface{}{"Path": path}))
	})

	entryRow := container.New(&centerPercentLayout{percent: 95}, entry)
	btnWidth := float32(7) * theme.TextSize()
	buttonRow := container.New(&buttonRowLayout{buttonWidth: btnWidth}, reindexButton, searchButton, newButton)

	content := container.NewVBox(entryRow, buttonRow, resultBox)

	w.SetContent(content)
	w.ShowAndRun()
}
