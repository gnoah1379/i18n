# i18n 
The easiest solution for multi-language support for your Golang project.

## What is i18n?
i18n is a Go package that helps you translate Go programs into multiple languages.
* Supports reading message file from files or directories.
* Supports reading from JSON, YAML, TOML, and HCL formats.
* Supports nested messages key.
* Supports render template with text/template.

## Install
> go get github.com/gnoah1379 /i18n

## Message file naming convention

> File name format: 
> > anything.{language}.{extension} 
> 
> File extension need in: json, yaml, yml, toml, hcl 

## LanguageBundle
```go
type LanguageBundle{
	// private field
}

// Get keyword data in the corresponding language.
//
// When the language does not exist or cannot find a keyword in the default language
// If not found this keyword in the default language, 
// return the first value of default values if default value existed
// or an empty string.
func (l LanguageBundle) Get(language, key string, defaultValue ...string) string

// Get template of keyword via GetTemplate method and Execute this template with data
func (l LanguageBundle) Render(lang, templateKey string, data interface{}, defaultTemplate ...string) (string, error)

// call Render and panic when error
func (l LanguageBundle) MustRender(lang, templateKey string, data interface{}, defaultTemplate ...string) string

// Get value of keyword via Get method and Parse value become Template
func (l LanguageBundle) GetTemplate(lang, templateKey string, defaultTemplate ...string) (*template.Template, error)
```
## Config
```go
type Config struct {
    DefaultLanguage string                       `json:"defaultLanguage" yaml:"defaultLanguage"`
    LanguagePaths   map[string][]string          `json:"languagePaths" yaml:"languagePaths"`
    DataInject      map[string]map[string]string `json:"dataInject" yaml:"dataInject"`
    Delimiter       string                       `json:"delimiters" yaml:"delimiters"`
}


// Load any files with filename match with convention.
func (c *Config) FromDir(dirpath string)error

// Load file if filename match with convention.
func (c *Config) FromFile(filepath string)error

// Data will inject into the bundle after loading from the file done. 
func (c *Config) InjectLanguageData(lang string, languageData map[string]interface{}) error

// Build config become a LanguageBundle.
//
// When data inject or message unmarshal is nested
// Will convert it become key-value format (map[string]string) with delimiter (default is ".").
func (c Config) Build() (*LanguageBundle, error)
```

# Usage
```go
 cfg := i18n.NewConfig("default language") 
 // config message dir
 err := cfg.FromDir("language messages directory")
 ...
 bundle,err := cfg.Build()
 msg := bundle.Get("language name","key","default value")
 msg = bundle.MustRender("language name","key",map[string]interface{}{
 	"bind": "value"
},"default template")
 ```