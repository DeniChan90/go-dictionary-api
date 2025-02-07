package constant

import (
	"github.com/DeniChan90/go-dictionary-api/models"
	"github.com/biter777/countries"
	"golang.org/x/text/language"
)

var Languages = []models.Language{
	{
		Code:   language.English,
		Symbol: countries.GB.Emoji(),
		Name:   "English",
	},
	{
		Code:   language.Ukrainian,
		Symbol: countries.Ukraine.Emoji(),
		Name:   "Ukrainan",
	},
	{
		Code:   language.German,
		Symbol: countries.Germany.Emoji(),
		Name:   "German",
	},
	{
		Code:   language.Polish,
		Symbol: countries.Poland.Emoji(),
		Name:   "Polish",
	},
	{
		Code:   language.Spanish,
		Symbol: countries.Spain.Emoji(),
		Name:   "Spanish",
	},
	{
		Code:   language.Swedish,
		Symbol: countries.Sweden.Emoji(),
		Name:   "Swedish",
	},
	{
		Code:   language.French,
		Symbol: countries.France.Emoji(),
		Name:   "French",
	},
	{
		Code:   language.Italian,
		Symbol: countries.Italy.Emoji(),
		Name:   "Italian",
	},
	{
		Code:   language.Portuguese,
		Symbol: countries.Portugal.Emoji(),
		Name:   "Portugese",
	},
	{
		Code:   language.Japanese,
		Symbol: countries.Japan.Emoji(),
		Name:   "Japanese",
	},
	{
		Code:   language.Vietnamese,
		Symbol: countries.Vietnam.Emoji(),
		Name:   "Vietnamese",
	},
}
