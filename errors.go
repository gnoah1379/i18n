package i18n

import "errors"

var ErrFileExtensionNotSupport = errors.New("file extension not support")
var ErrInvalidLanguage = errors.New("invalid language")
var ErrInvalidLanguageData = errors.New("invalid language data")
var ErrRootBundleDoesNotInit = errors.New("root bundle does not init")
