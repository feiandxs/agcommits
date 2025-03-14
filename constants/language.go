package constants

// LanguageCode 语言代码类型
type LanguageCode string

// 支持的语言代码
const (
	// DefaultLanguage 默认语言
	DefaultLanguage = EN

	// 语言代码定义
	EN LanguageCode = "en"
	ZH LanguageCode = "zh"
	// 未来可以添加更多语言
	// FR LanguageCode = "fr"
	// JA LanguageCode = "ja"
)

// languageDisplayNames 语言代码到显示名称的映射
var languageDisplayNames = map[LanguageCode]string{
	EN: "English",
	ZH: "中文",
	// 未来可以添加更多语言
	// FR: "Français",
	// JA: "日本語",
}

// languagePromptNames 语言代码到AI提示中使用的名称的映射
var languagePromptNames = map[LanguageCode]string{
	EN: "English",
	ZH: "中文",
	// 未来可以添加更多语言
	// FR: "French",
	// JA: "Japanese",
}

// IsValidLanguage 检查语言代码是否有效
func IsValidLanguage(code string) bool {
	_, ok := languageDisplayNames[LanguageCode(code)]
	return ok
}

// GetLanguageDisplay 获取语言的显示名称
func GetLanguageDisplay(code LanguageCode) string {
	if name, ok := languageDisplayNames[code]; ok {
		return name
	}
	return string(code)
}

// GetLanguagePromptName 获取语言在AI提示中使用的名称
func GetLanguagePromptName(code LanguageCode) string {
	if name, ok := languagePromptNames[code]; ok {
		return name
	}
	return string(code)
}

// GetSupportedLanguages 获取所有支持的语言代码
func GetSupportedLanguages() []LanguageCode {
	languages := make([]LanguageCode, 0, len(languageDisplayNames))
	for code := range languageDisplayNames {
		languages = append(languages, code)
	}
	return languages
}
