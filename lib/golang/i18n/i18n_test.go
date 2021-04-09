package i18n

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func TestText(t *testing.T) {
	t.Log(InitI18nYaml())
}

func InitI18n() string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("./zh-TW.toml")
	localizer := i18n.NewLocalizer(bundle, "zh-TW")
	return localizer.MustLocalize((&i18n.LocalizeConfig{MessageID: "personCats"}))
}

func InitI18nYaml() string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.LoadMessageFile("./zh-TW.yaml")
	localizer := i18n.NewLocalizer(bundle, "zh-TW")
	pc := &PersonCat{
		Name:  "test",
		Count: 1,
		Test:  Test{},
	}
	return localizer.MustLocalize((&i18n.LocalizeConfig{
		MessageID:    "chatbot_personCats",
		TemplateData: pc,
		PluralCount:  1,
	}))
}

type PersonCat struct {
	Name  string
	Count int
	Test  Test
}

type Test struct {
}

func (Test) String() string {
	return "this is testing"
}
