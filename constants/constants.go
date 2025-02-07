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
	{
		Code:   language.Kazakh,
		Symbol: countries.Kazakhstan.Emoji(),
		Name:   "Kazakh",
	},
	{
		Code:   language.Azerbaijani,
		Symbol: countries.Azerbaijan.Emoji(),
		Name:   "Azerbaijan",
	},
	{
		Code:   language.Finnish,
		Symbol: countries.Finland.Emoji(),
		Name:   "Finnish",
	},
	{
		Code:   language.Norwegian,
		Symbol: countries.Norway.Emoji(),
		Name:   "Norwegian",
	},
	{
		Code:   language.Danish,
		Symbol: countries.Denmark.Emoji(),
		Name:   "Danish",
	},
	{
		Code:   language.Lithuanian,
		Symbol: countries.Lithuania.Emoji(),
		Name:   "Lithuanian",
	},
	{
		Code:   language.Latvian,
		Symbol: countries.Latvia.Emoji(),
		Name:   "Latvian",
	},
	{
		Code:   language.Estonian,
		Symbol: countries.Estonia.Emoji(),
		Name:   "Estonian",
	},
}
