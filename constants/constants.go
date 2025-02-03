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
	},
	{
		Code:   language.Ukrainian,
		Symbol: countries.Ukraine.Emoji(),
	},
	{
		Code:   language.Spanish,
		Symbol: countries.Spain.Emoji(),
	},
	{
		Code:   language.French,
		Symbol: countries.France.Emoji(),
	},
	{
		Code:   language.Vietnamese,
		Symbol: countries.Vietnam.Emoji(),
	},
}
