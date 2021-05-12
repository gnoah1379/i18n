package i18n

import "text/template"

var root *LanguageBundle

func Build(c Config) error {
	bundle, err := c.Build()
	if err == nil {
		root = bundle
	}
	return err
}

// clone of root
func Clone() *LanguageBundle {
	if root == nil {
		return nil
	}
	bundle := newLanguageBundle(root.defaultLang)
	for lang, messages := range root.languages {
		for k, v := range messages {
			bundle.set(lang, k, v)
		}
	}
	return bundle
}

func Get(lang, templateKey string, defaultTemplate ...string) string {
	if root == nil {
		panic(ErrRootBundleDoesNotInit)
	}
	return root.get(lang, templateKey, defaultTemplate)
}

func Render(lang, templateKey string, data interface{}, defaultTemplate ...string) (string, error) {
	if root == nil {
		return "", ErrRootBundleDoesNotInit
	}
	return root.render(lang, templateKey, data, defaultTemplate)
}

func MustRender(lang, templateKey string, data interface{}, defaultTemplate ...string) string {
	if root == nil {
		panic(ErrRootBundleDoesNotInit)
	}
	return root.mustRender(lang, templateKey, data, defaultTemplate)
}

func GetTemplate(lang, templateKey string, defaultTemplate ...string) (*template.Template, error) {
	if root == nil {
		return nil, ErrRootBundleDoesNotInit
	}
	return root.getTemplate(lang, templateKey, defaultTemplate)
}
