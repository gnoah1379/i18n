package i18n

import (
	"testing"
)

var (
	enData = map[string]interface{}{
		"SYS_ERR":           "System error: {{.Error}}",
		"ERR_INVALID_PARAM": "invalid parameter",
		"ERR_INVALID_KEY":   "invalid key",
	}
	viData = map[string]interface{}{
		"SYS_ERR":           "Lỗi hệ thống: {{.Error}}",
		"ERR_INVALID_PARAM": "tham số không hợp lệ",
	}
	esData = map[string]interface{}{
		"SYS_ERR":           "Error del sistema: {{.Error}}",
		"ERR_INVALID_PARAM": "parámetro no válido",
	}
)

var (
	EN = "en"
	VI = "vi"
	ES = "es"
)

func getBundle() (*LanguageBundle, error) {
	var cfg = NewConfig(EN)
	if err := cfg.InjectLanguageData(EN, enData); err != nil {
		return nil, err
	}
	if err := cfg.InjectLanguageData(VI, viData); err != nil {
		return nil, err
	}
	if err := cfg.InjectLanguageData(ES, esData); err != nil {
		return nil, err
	}
	return cfg.Build()
}

func TestLanguageBundle_Get(t *testing.T) {
	bundle, err := getBundle()
	if err != nil {
		t.Error(err)
	}
	if value := bundle.Get(EN, "ERR_INVALID_PARAM"); value != enData["ERR_INVALID_PARAM"] {
		t.Errorf("language %s get error value: %s", EN, value)
	}
	if value := bundle.Get(VI, "ERR_INVALID_PARAM"); value != viData["ERR_INVALID_PARAM"] {
		t.Errorf("language %s get error value: %s", VI, value)
	}
	if value := bundle.Get(ES, "ERR_INVALID_PARAM"); value != esData["ERR_INVALID_PARAM"] {
		t.Errorf("language %s get error value: %s", ES, value)
	}
}

func TestLanguageBundle_MustPack(t *testing.T) {
	bundle, err := getBundle()
	if err != nil {
		t.Error(err)
	}
	if value := bundle.MustRender(EN, "SYS_ERR", map[string]string{
		"Error": bundle.Get(EN, "ERR_INVALID_PARAM"),
	}); value != "System error: invalid parameter" {
		t.Errorf("pack message error value: %s", value)
	}
}

func TestLanguageBundle_GetMissKey(t *testing.T) {
	bundle, err := getBundle()
	if err != nil {
		t.Error(err)
	}
	if value := bundle.Get(VI, "ERR_INVALID_KEY"); value != enData["ERR_INVALID_KEY"] {
		t.Errorf("language %s case miss key error value: %s", VI, value)
	}
	if value := bundle.Get(ES, "ERR_INVALID_KEY"); value != enData["ERR_INVALID_KEY"] {
		t.Errorf("language %s case miss key error value: %s", ES, value)
	}
}
