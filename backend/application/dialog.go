package application

import (
	"context"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DialogImp struct {
	Ctx context.Context
}

func NewDialog(ctx context.Context) *DialogImp {
	return &DialogImp{
		Ctx: ctx,
	}
}

func (p *DialogImp) Confirm(title, message string) bool {
	res, err := runtime.MessageDialog(p.Ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "Yes",
		CancelButton:  "No",
	})

	return err == nil && res == "Yes"
}

func (p *DialogImp) Show(title, message string) {
	runtime.MessageDialog(p.Ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: message,
	})
}

func (p *DialogImp) Error(err error) {
	p.Show("Error", err.Error())
}

func (p *DialogImp) Save(title, fileName string, filters ...string) (path string, err error) {
	filts := make([]runtime.FileFilter, len(filters))
	for i, f := range filters {
		displayname, pattern, _ := strings.Cut(f, "|")

		filts[i] = runtime.FileFilter{
			DisplayName: displayname,
			Pattern:     pattern,
		}
	}

	return runtime.SaveFileDialog(p.Ctx, runtime.SaveDialogOptions{
		DefaultFilename: fileName,
		Title:           title,
		Filters:         filts,
	})
}

func (p *DialogImp) Open(title string, filters ...string) (path string, err error) {
	filts := make([]runtime.FileFilter, len(filters))
	for i, f := range filters {
		displayname, pattern, _ := strings.Cut(f, "|")

		filts[i] = runtime.FileFilter{
			DisplayName: displayname,
			Pattern:     pattern,
		}
	}
	defaultDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	defaultDirectory += "/ATcnea"
	return runtime.OpenFileDialog(p.Ctx, runtime.OpenDialogOptions{
		Title:            title,
		Filters:          filts,
		DefaultDirectory: defaultDirectory,
	})
}

func (p *DialogImp) OpenDir(title string) (path string, err error) {
	return runtime.OpenDirectoryDialog(p.Ctx, runtime.OpenDialogOptions{
		Title: title,
	})
}
