package i18n

import (
	"bytes"
	"text/template"
)

type LanguageBundle struct {
	languages   map[string]map[string]string
	defaultLang string
}

func newLanguageBundle(defaultLang string) *LanguageBundle {
	l := &LanguageBundle{languages: map[string]map[string]string{
		defaultLang: make(map[string]string),
	}, defaultLang: defaultLang}
	return l
}

func (l LanguageBundle) Get(lang, key string, defaultValue ...string) string {
	return l.get(lang, key, defaultValue)
}
func (l *LanguageBundle) get(langKey, key string, defaultValue []string) string {
	lang, isNotDefault := l.getLang(langKey)
	message, ok := lang[key]
	if !ok {
		if isNotDefault {
			if message, ok = l.languages[l.defaultLang][key]; ok {
				return message
			}
		}
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return message
}

func (l LanguageBundle) MustRender(lang, templateKey string, data interface{}, defaultTemplate ...string) string {
	return l.mustRender(lang, templateKey, data, defaultTemplate)
}
func (l LanguageBundle) mustRender(lang, templateKey string, data interface{}, defaultTemplate []string) string {
	message, err := l.render(lang, templateKey, data, defaultTemplate)
	must(err)
	return message
}

func (l LanguageBundle) Render(lang, templateKey string, data interface{}, defaultTemplate ...string) (string, error) {
	return l.render(lang, templateKey, data, defaultTemplate)
}
func (l LanguageBundle) render(lang, templateKey string, data interface{}, defaultTemplate []string) (string, error) {
	tpl, err := l.getTemplate(lang, templateKey, defaultTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (l LanguageBundle) GetTemplate(lang, templateKey string, defaultTemplate ...string) (*template.Template, error) {
	return l.getTemplate(lang, templateKey, defaultTemplate)
}
func (l LanguageBundle) getTemplate(lang, templateKey string, defaultTemplate []string) (*template.Template, error) {
	tpl := l.get(lang, templateKey, defaultTemplate)
	return template.New(templateKey).Parse(tpl)
}

func (l *LanguageBundle) set(lang, key, value string) {
	if _, ok := l.languages[lang]; !ok {
		l.languages[lang] = make(map[string]string)
	}
	l.languages[lang][key] = value
}

func (l *LanguageBundle) getLang(langKey string) (map[string]string, bool) {
	lang, ok := l.languages[langKey]
	if !ok {
		return l.languages[l.defaultLang], false
	}
	return lang, true
}
