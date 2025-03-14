package constants

// LanguageCode represents the type for language codes
type LanguageCode string

// Supported language codes
const (
	// DefaultLanguage specifies the default language for the application
	DefaultLanguage = EN

	// Language code definitions
	EN LanguageCode = "en" // English
	ZH LanguageCode = "zh" // Chinese
	ES LanguageCode = "es" // Spanish
	HI LanguageCode = "hi" // Hindi
	AR LanguageCode = "ar" // Arabic
	FR LanguageCode = "fr" // French
	DE LanguageCode = "de" // German
	PT LanguageCode = "pt" // Portuguese
	JA LanguageCode = "ja" // Japanese
	KO LanguageCode = "ko" // Korean
	RU LanguageCode = "ru" // Russian
)

// languageDisplayNames maps language codes to their native display names
var languageDisplayNames = map[LanguageCode]string{
	EN: "English",
	ZH: "中文",
	ES: "Español",
	HI: "हिन्दी",
	AR: "العربية",
	FR: "Français",
	DE: "Deutsch",
	PT: "Português",
	JA: "日本語",
	KO: "한국어",
	RU: "Русский",
}

// languagePromptNames maps language codes to their English names for AI prompts
var languagePromptNames = map[LanguageCode]string{
	EN: "English",
	ZH: "Chinese",
	ES: "Spanish",
	HI: "Hindi",
	AR: "Arabic",
	FR: "French",
	DE: "German",
	PT: "Portuguese",
	JA: "Japanese",
	KO: "Korean",
	RU: "Russian",
}

// IsValidLanguage checks if the provided language code is valid
func IsValidLanguage(code string) bool {
	_, ok := languageDisplayNames[LanguageCode(code)]
	return ok
}

// GetLanguageDisplay returns the native display name for a language code
func GetLanguageDisplay(code LanguageCode) string {
	if name, ok := languageDisplayNames[code]; ok {
		return name
	}
	return string(code)
}

// GetLanguagePromptName returns the English name of the language for AI prompts
func GetLanguagePromptName(code LanguageCode) string {
	if name, ok := languagePromptNames[code]; ok {
		return name
	}
	return string(code)
}

// GetSupportedLanguages returns a slice of all supported language codes
func GetSupportedLanguages() []LanguageCode {
	languages := make([]LanguageCode, 0, len(languageDisplayNames))
	for code := range languageDisplayNames {
		languages = append(languages, code)
	}
	return languages
}
