package constants

import (
	"fmt"
)

// LanguageCode represents the type for language codes
type LanguageCode string

// LanguageInfo represents all information about a language
type LanguageInfo struct {
	Index       int    // Display index for selection
	NativeName  string // Name in native language
	EnglishName string // Name in English
	Description string // Optional description or region info
}

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

// languageInfo stores all language related information
var languageInfo = map[LanguageCode]LanguageInfo{
	EN: {Index: 1, NativeName: "English", EnglishName: "English", Description: "Default"},
	ZH: {Index: 2, NativeName: "中文", EnglishName: "Chinese", Description: "Simplified Chinese"},
	ES: {Index: 3, NativeName: "Español", EnglishName: "Spanish", Description: "Spanish"},
	HI: {Index: 4, NativeName: "हिन्दी", EnglishName: "Hindi", Description: "Hindi"},
	AR: {Index: 5, NativeName: "العربية", EnglishName: "Arabic", Description: "Arabic"},
	FR: {Index: 6, NativeName: "Français", EnglishName: "French", Description: "French"},
	DE: {Index: 7, NativeName: "Deutsch", EnglishName: "German", Description: "German"},
	PT: {Index: 8, NativeName: "Português", EnglishName: "Portuguese", Description: "Portuguese"},
	JA: {Index: 9, NativeName: "日本語", EnglishName: "Japanese", Description: "Japanese"},
	KO: {Index: 10, NativeName: "한국어", EnglishName: "Korean", Description: "Korean"},
	RU: {Index: 11, NativeName: "Русский", EnglishName: "Russian", Description: "Russian"},
}

// GetLanguageByIndex returns the language code for a given index
func GetLanguageByIndex(index int) (LanguageCode, bool) {
	for code, info := range languageInfo {
		if info.Index == index {
			return code, true
		}
	}
	return "", false
}

// GetLanguageInfo returns the complete language information for a given code
func GetLanguageInfo(code LanguageCode) (LanguageInfo, bool) {
	info, ok := languageInfo[code]
	return info, ok
}

// GetFormattedLanguageList returns a formatted list of languages for display
func GetFormattedLanguageList() []string {
	result := make([]string, len(languageInfo))
	for code, info := range languageInfo {
		result[info.Index-1] = fmt.Sprintf("%2d. %-10s (%s)",
			info.Index,
			info.EnglishName,
			info.NativeName)
	}
	return result
}

// IsValidLanguage checks if the provided language code is valid
func IsValidLanguage(code string) bool {
	_, ok := languageInfo[LanguageCode(code)]
	return ok
}

// GetLanguageDisplay returns the native display name for a language code
func GetLanguageDisplay(code LanguageCode) string {
	if info, ok := languageInfo[code]; ok {
		return info.NativeName
	}
	return string(code)
}

// GetLanguagePromptName returns the English name of the language for AI prompts
func GetLanguagePromptName(code LanguageCode) string {
	if info, ok := languageInfo[code]; ok {
		return info.EnglishName
	}
	return string(code)
}

// GetSupportedLanguages returns a slice of all supported language codes
func GetSupportedLanguages() []LanguageCode {
	languages := make([]LanguageCode, 0, len(languageInfo))
	for code := range languageInfo {
		languages = append(languages, code)
	}
	return languages
}
