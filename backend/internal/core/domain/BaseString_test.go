package domain

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var language Language
var baseString BaseString

func initValues() {
	language = Language{
		ID:          primitive.ObjectID{},
		IsoCode:     "",
		Description: "",
		Active:      true,
	}
	t1 := Translation{
		Content:  "1",
		Language: &language,
		Version:  1,
		Active:   true,
	}
	t2 := Translation{
		Content:  "2",
		Language: &language,
		Version:  2,
		Active:   true,
	}
	baseString = BaseString{
		Translations: []Translation{t1, t2},
	}
}

func TestBaseString_FindTranslationLastVersionByLanguage(t *testing.T) {
	initValues()
	assert.Equal(t, baseString.FindTranslationLastVersionByLanguage(language), baseString.Translations[1].Content)
}

func TestBaseString_FindTranslationLastVersionByLanguage_InvalidLanguage(t *testing.T) {
	initValues()
	language.Active = false
	baseString.SourceLanguage = &language
	assert.Empty(t, baseString.FindTranslationLastVersionByLanguage(language))
}

func TestBaseString_FindTranslationLastVersionByLanguage_InvalidTranslation(t *testing.T) {
	initValues()
	baseString.Translations[1].Active = false
	assert.Equal(t, baseString.FindTranslationLastVersionByLanguage(language), baseString.Translations[0].Content)
}
