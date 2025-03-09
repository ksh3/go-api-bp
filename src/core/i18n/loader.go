package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Translator struct {
	translations map[string]map[string]string
}

func NewTranslator() *Translator {
	return &Translator{translations: make(map[string]map[string]string)}
}

func (t *Translator) Translate(lang, key string) string {
	if values, exists := t.translations[lang]; exists {
		if translation, ok := values[key]; ok {
			return translation
		}
	}
	return key
}

func (t *Translator) LoadTranslations(corePath string, featurePath string) error {
	if err := t.loadTranslationFiles(corePath); err != nil {
		fmt.Printf("Failed to load core translations: %v\n", err)
	}

	// NOTE: find feature directories.
	features, err := os.ReadDir(featurePath)
	if err != nil {
		return fmt.Errorf("failed to read features: %w", err)
	}

	for _, feature := range features {
		if feature.IsDir() {
			featureI18nPath := filepath.Join(featurePath, feature.Name(), "i18n")
			if err := t.loadTranslationFiles(featureI18nPath); err != nil {
				fmt.Printf("Skipping %s: %v\n", featureI18nPath, err)
			}
		}
	}
	return nil
}

func (t *Translator) loadTranslationFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read translation files: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			lang := file.Name()[:len(file.Name())-5] // "en.json" -> "en"
			filePath := filepath.Join(dir, file.Name())

			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Failed to read %s: %v\n", filePath, err)
				continue
			}

			var translations map[string]string
			if err := json.Unmarshal(data, &translations); err != nil {
				fmt.Printf("Invalid JSON in %s: %v\n", filePath, err)
				continue
			}

			// 🔥 既存データとマージ
			if _, exists := t.translations[lang]; !exists {
				t.translations[lang] = make(map[string]string)
			}
			for k, v := range translations {
				t.translations[lang][k] = v
			}
		}
	}
	return nil
}
