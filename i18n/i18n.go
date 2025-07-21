package i18n

import (
	"embed"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

//go:embed *.yaml
var localeFS embed.FS

// Load initializes the localization bundle and sets the active language.
func Load(lang string) error {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	entries, err := localeFS.ReadDir(".")
	if err != nil {
		return err
	}
	for _, e := range entries {
		if _, err := bundle.LoadMessageFileFS(localeFS, e.Name()); err != nil {
			return err
		}
	}
	localizer = i18n.NewLocalizer(bundle, lang)
	return nil
}

// T localizes the given message ID using optional template data.
func T(id string, data map[string]interface{}) string {
	if localizer == nil {
		return id
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: data})
	if err != nil {
		return id
	}
	return msg
}
