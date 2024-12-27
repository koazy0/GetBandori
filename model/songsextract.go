package model

// SongExtractModel 拿来存数据库的
// SongData有太多用不到的信息
type SongExtractModel struct {
	MusicTitle     string           `json:"musicTitle"`
	Lyricist       string           `json:"lyricist"` // 作词者
	Composer       string           `json:"composer"` // 作曲者
	Arranger       string           `json:"arranger"` // 编曲者
	Length         float64          `json:"length"`
	EachDifficulty []EachDifficulty `json:"eachDifficulty"` //每个等级一个数组
}

type EachDifficulty struct {
	PlayLevel        int    `json:"playLevel"`     // 难度等级
	NotesQuantity    int    `json:"notesQuantity"` // 音符数量
	Notes            int    `json:"notes"`
	BPM              BPM    `json:"bpm"` // BPM 信息
	ScreenShotUrl    string `json:"screenShotUrl"`
	ScreenShotBase64 string `json:"screenShotBase64" gorm:"type:text;size:65535"`
}
