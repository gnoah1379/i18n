package i18n

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var supportFileExtension = []string{
	"json",
	"yaml", "yml",
	"toml",
}

type Config struct {
	DefaultLanguage string                       `json:"defaultLanguage" yaml:"defaultLanguage"`
	LanguagePaths   map[string][]string          `json:"languagePaths" yaml:"languagePaths"`
	DataInject      map[string]map[string]string `json:"dataInject" yaml:"dataInject"`
	Delimiter       string                       `json:"delimiters" yaml:"delimiters"`
}

func NewConfig(defaultLanguage string) Config {
	return Config{DefaultLanguage: defaultLanguage, Delimiter: "."}
}

func (c Config) Build() (*LanguageBundle, error) {
	if c.DefaultLanguage == "" {
		return nil, ErrInvalidLanguage
	}
	bundle := newLanguageBundle(c.DefaultLanguage)
	for lang, paths := range c.LanguagePaths {
		for _, path := range paths {
			data, err := parseFile(c.Delimiter, path)
			if err != nil {
				return nil, err
			}
			for k, v := range data {
				bundle.set(lang, k, v)
			}
		}
	}
	for lang, data := range c.DataInject {
		for k, v := range data {
			bundle.set(lang, k, v)
		}
	}
	return bundle, nil
}

func (c *Config) InjectLanguageData(lang string, languageData map[string]interface{}) error {
	if lang == "" {
		return ErrInvalidLanguage
	}
	if languageData == nil {
		return ErrInvalidLanguageData
	}
	if c.DataInject == nil {
		c.DataInject = make(map[string]map[string]string)
	}
	c.DataInject[lang] = unpackMessage(c.Delimiter, languageData)
	return nil
}

func (c *Config) FromFile(filePath string) error {
	f, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	return c.addFile(f, filePath)
}

func (c *Config) FromDir(dirPath string) error {
	filesInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, f := range filesInfo {
		_ = c.addFile(f, filepath.Join(dirPath, f.Name()))
	}
	return nil
}

func (c *Config) addFile(f os.FileInfo, filePath string) (err error) {
	if err = validateFileExtension(f); err == nil {
		lang, err := getFileLanguage(f)
		if err != nil {
			return err
		}
		if lang == "" {
			return ErrInvalidLanguage
		}
		if c.LanguagePaths == nil {
			c.LanguagePaths = make(map[string][]string)
		}
		c.LanguagePaths[lang] = append(c.LanguagePaths[lang], filePath)
	}
	return err
}
