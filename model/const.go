package model

// 乐团
const (
	PPP     = "Poppin'Party"
	Roselia = "Roselia"
	PP      = "Pastel＊Palettes"
)

// 人物

type SongsInformation struct {
	ID int

	Title         string
	TitleAlias    string
	Alias         string
	Kind          string
	Lyrics        string
	LyricsAlias   string
	Composer      string
	ComposerAlias string
	Arranger      string
	ArrangerAlias string
}

const (
	UrlBase = "https://bestdori.com/api/songs/"
	Xpath   = "/html/body/div/div[4]/div[2]/div[1]/div[6]/table/tbody/tr[1]/td[2]/div/div"
	MaxID   = 2000
)
