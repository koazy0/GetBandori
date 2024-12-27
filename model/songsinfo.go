package model

// Achievement 表示每个成就信息
type Achievement struct {
	MusicID         int    `json:"musicId"`            // 音乐 ID
	AchievementType string `json:"achievementType"`    // 成就类型
	RewardType      string `json:"rewardType"`         // 奖励类型
	RewardID        int    `json:"rewardId,omitempty"` // 奖励 ID（可选）
	Quantity        int    `json:"quantity"`           // 奖励数量
}

// MultiLiveScore 表示多人模式的分数信息
type MultiLiveScore struct {
	MusicID                 int    `json:"musicId"`                 // 音乐 ID
	MusicDifficulty         string `json:"musicDifficulty"`         // 音乐难度
	MultiLiveDifficultyID   int    `json:"multiLiveDifficultyId"`   // 多人模式难度 ID
	ScoreS                  int    `json:"scoreS"`                  // S 分数
	ScoreA                  int    `json:"scoreA"`                  // A 分数
	ScoreB                  int    `json:"scoreB"`                  // B 分数
	ScoreC                  int    `json:"scoreC"`                  // C 分数
	MultiLiveDifficultyType string `json:"multiLiveDifficultyType"` // 多人模式类型
	ScoreSS                 int    `json:"scoreSS"`                 // SS 分数
	ScoreSSS                int    `json:"scoreSSS"`                // SSS 分数
}

// Difficulty 表示每种难度的详细信息
type Difficulty struct {
	PlayLevel         int                       `json:"playLevel"`             // 难度等级
	MultiLiveScoreMap map[string]MultiLiveScore `json:"multiLiveScoreMap"`     // 多人模式分数映射
	NotesQuantity     int                       `json:"notesQuantity"`         // 音符数量
	ScoreC            int                       `json:"scoreC"`                // C 分数阈值
	ScoreB            int                       `json:"scoreB"`                // B 分数阈值
	ScoreA            int                       `json:"scoreA"`                // A 分数阈值
	ScoreS            int                       `json:"scoreS"`                // S 分数阈值
	ScoreSS           int                       `json:"scoreSS"`               // SS 分数阈值
	PublishedAt       []string                  `json:"publishedAt,omitempty"` // 发布时间（可选）
}

// BPM 表示歌曲的 BPM 范围
type BPM struct {
	BPM   float64 `json:"bpm"`   // BPM 值
	Start float64 `json:"start"` // 开始时间（秒）
	End   float64 `json:"end"`   // 结束时间（秒）
}

// SongData 表示主 JSON 数据结构
type SongData struct {
	BgmID        string                `json:"bgmId"`        // 背景音乐 ID
	BgmFile      string                `json:"bgmFile"`      // 背景音乐文件名
	Tag          string                `json:"tag"`          // 标签
	BandID       int                   `json:"bandId"`       // 乐队 ID
	Achievements []Achievement         `json:"achievements"` // 成就列表
	JacketImage  []string              `json:"jacketImage"`  // 封面图像列表
	Seq          int                   `json:"seq"`          // 顺序号
	MusicTitle   []string              `json:"musicTitle"`   // 歌曲标题（多语言）
	Ruby         []string              `json:"ruby"`         // 歌曲标题的拼音或发音
	Phonetic     []string              `json:"phonetic"`     // 歌曲标题的音标
	Lyricist     []string              `json:"lyricist"`     // 作词者
	Composer     []string              `json:"composer"`     // 作曲者
	Arranger     []string              `json:"arranger"`     // 编曲者
	HowToGet     []string              `json:"howToGet"`     // 获取方式
	PublishedAt  []string              `json:"publishedAt"`  // 发布时间
	ClosedAt     []string              `json:"closedAt"`     // 结束时间
	Difficulty   map[string]Difficulty `json:"difficulty"`   // 难度信息
	Length       float64               `json:"length"`       // 歌曲长度（秒）
	Notes        map[string]int        `json:"notes"`        // 音符数量
	BPM          map[string][]BPM      `json:"bpm"`          // BPM 信息
}

// IsValid 函数用于检查SongData是否有效,实现数组接口
func (s SongData) IsValid() bool {
	return s.BgmID != ""
}

type Songdatas []SongData

// FilterEmptyEntries filterEmptyEntries 函数用于过滤掉数组中的空元素
func (S Songdatas) FilterEmptyEntries() Array {
	var filtered Songdatas
	for _, data := range S {
		if data.IsValid() {
			filtered = append(filtered, data)
		}
	}
	return Array(filtered)
}
