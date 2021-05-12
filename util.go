package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func validateFileExtension(f os.FileInfo) error {
	filename := strings.Split(f.Name(), ".")
	if len(filename) >= 2 {
		fileExtension := filename[len(filename)-1]
		for _, ext := range supportFileExtension {
			if ext == fileExtension {
				return nil
			}
		}
	}
	return ErrFileExtensionNotSupport
}

func parseFile(delimiter string, path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	inPackMessage := make(map[string]interface{})
	paths := strings.Split(path, ".")
	ext := paths[len(paths)-1]
	switch ext {
	case "json":
		err = json.Unmarshal(data, &inPackMessage)
	case "toml":
		err = toml.Unmarshal(data, &inPackMessage)
	case "yaml", "yml":
		err = yaml.Unmarshal(data, &inPackMessage)
	default:
		err = ErrFileExtensionNotSupport
	}
	if err != nil {
		return nil, err
	}
	return unpackMessage(delimiter, inPackMessage), nil
}

func getFileLanguage(f os.FileInfo) (string, error) {
	filename := strings.Split(f.Name(), ".")
	if len(filename) < 2 {
		return "", ErrInvalidLanguage
	}
	lang := filename[len(filename)-2]
	return lang, nil
}

func unpackMessage(delimiter string, src map[string]interface{}) map[string]string {
	if delimiter == "" {
		delimiter = "."
	}
	root := make(map[string]string)
	for key, value := range src {
		switch value.(type) {
		case string:
			root[key] = value.(string)
		case map[string]interface{}:
			child := unpackMessage(delimiter, value.(map[string]interface{}))
			for k, v := range child {
				root[key+delimiter+k] = v
			}
		default:
			root[key] = fmt.Sprintf("%v", value)
		}
	}
	return root
}
