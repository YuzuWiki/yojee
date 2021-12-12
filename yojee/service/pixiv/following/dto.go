package following

import (
	"encoding/json"
	"time"
)

type pageDTO struct {
	Ids  []int32       `json:"ids"`
	Tags []interface{} `json:"tags"`
}

type tagDTO struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	ZhTw   string `json:"zh_tw"`
	Romaji string `json:"romaji"`
}

type illustDTO struct {
	ID                      string                 `json:"id"`
	Title                   string                 `json:"title"`
	IllustType              int                    `json:"illustType"`
	XRestrict               int                    `json:"xRestrict"`
	Restrict                int                    `json:"restrict"`
	Sl                      int                    `json:"sl"`
	URL                     string                 `json:"url"`
	Description             string                 `json:"description"`
	Tags                    []string               `json:"tags"`
	UserID                  string                 `json:"userId"`
	UserName                string                 `json:"userName"`
	Width                   int                    `json:"width"`
	Height                  int                    `json:"height"`
	PageCount               int                    `json:"pageCount"`
	IsBookmarkable          bool                   `json:"isBookmarkable"`
	Alt                     string                 `json:"alt"`
	TitleCaptionTranslation map[string]interface{} `json:"titleCaptionTranslation"`
	CreateDate              time.Time              `json:"createDate"`
	UpdateDate              time.Time              `json:"updateDate"`
	IsUnlisted              bool                   `json:"isUnlisted"`
	IsMasked                bool                   `json:"isMasked"`
	Urls                    map[string]string      `json:"urls"`
	ProfileImageURL         string                 `json:"profileImageUrl"`
}

type novelDTO struct {
	ID                      int64                  `json:"id,string"`
	Title                   string                 `json:"title"`
	XRestrict               int                    `json:"xRestrict"`
	Restrict                int                    `json:"restrict"`
	URL                     string                 `json:"url"`
	Tags                    []string               `json:"tags"`
	UserID                  int64                  `json:"userId,string"`
	UserName                string                 `json:"userName"`
	ProfileImageURL         string                 `json:"profileImageUrl"`
	TextCount               int                    `json:"textCount"`
	Description             string                 `json:"description"`
	IsBookmarkable          bool                   `json:"isBookmarkable"`
	BookmarkCount           int                    `json:"bookmarkCount"`
	IsOriginal              bool                   `json:"isOriginal"`
	TitleCaptionTranslation map[string]interface{} `json:"titleCaptionTranslation"`
	IsMasked                bool                   `json:"isMasked"`
	IsUnlisted              bool                   `json:"isUnlisted"`
	CreateDate              time.Time              `json:"createDate"`
	UpdateDate              time.Time              `json:"updateDate"`
}

type novelSeriesDTO struct {
	ID                      int64                        `json:"id,string"`
	Title                   string                       `json:"title"`
	TitleCaptionTranslation interface{}                  `json:"titleCaptionTranslation"`
	Cover                   map[string]map[string]string `json:"cover"`
	Tags                    []string                     `json:"tags"`
	XRestrict               int                          `json:"xRestrict"`
	IsOriginal              bool                         `json:"isOriginal"`
	Genre                   string                       `json:"genre"`
	UserID                  int64                        `json:"userId,string"`
	UserName                string                       `json:"userName"`
	ProfileImageURL         string                       `json:"profileImageUrl"`
	BookmarkCount           int                          `json:"bookmarkCount"`
	IsOneshot               bool                         `json:"isOneshot"`
	Caption                 string                       `json:"caption"`
	IsConcluded             bool                         `json:"isConcluded"`
	EpisodeCount            int                          `json:"episodeCount"`
	PublishedEpisodeCount   int                          `json:"publishedEpisodeCount"`
	LatestPublishDateTime   time.Time                    `json:"latestPublishDateTime"`
	LatestEpisodeID         int64                        `json:"latestEpisodeId,string"`
	IsWatched               bool                         `json:"isWatched"`
	IsNotifying             bool                         `json:"isNotifying"`
	Restrict                int                          `json:"restrict"`
	TextLength              int                          `json:"textLength"`
	PublishedTextLength     int                          `json:"publishedTextLength"`
	CreateDateTime          time.Time                    `json:"createDateTime"`
	UpdateDateTime          time.Time                    `json:"updateDateTime"`
}

type thumbnailDTO struct {
	Illust []illustDTO `json:"illust"`
	Novel  []novelDTO  `json:"novel"`
}

type userDTO struct {
	Partial       int         `json:"partial"`
	Comment       string      `json:"comment"`
	FollowedBack  bool        `json:"followedBack"`
	UserID        int64       `json:"userId,string"`
	Name          string      `json:"name"`
	Image         string      `json:"image"`
	ImageBig      string      `json:"imageBig"`
	Premium       bool        `json:"premium"`
	IsFollowed    bool        `json:"isFollowed"`
	IsMypixiv     bool        `json:"isMypixiv"`
	IsBlocking    bool        `json:"isBlocking"`
	Background    interface{} `json:"background"`
	AcceptRequest bool        `json:"acceptRequest"`
}

type illustSeriesDTO struct {
	ID             int64       `json:"id,string"`
	UserID         int64       `json:"userId,string"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Caption        string      `json:"caption"`
	Total          int         `json:"total"`
	ContentOrder   interface{} `json:"content_order"`
	URL            string      `json:"url"`
	CoverImageSl   int         `json:"coverImageSl"`
	FirstIllustID  int64       `json:"firstIllustId,string"`
	LatestIllustID int64       `json:"latestIllustId,string"`
	WatchCount     interface{} `json:"watchCount"`
	IsWatched      bool        `json:"isWatched"`
	IsNotifying    bool        `json:"isNotifying"`
	CreateDate     time.Time   `json:"createDate"`
	UpdateDate     time.Time   `json:"updateDate"`
}

type tagTranslationDTO map[string]tagDTO

func (dto *tagTranslationDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := &tagTranslationDTO{}
	if err := json.Unmarshal(body, data); err != nil {
		return err
	}

	dto = data
	return nil
}

type FollowLatestDTO struct {
	Page           pageDTO           `json:"page"`
	TagTranslation tagTranslationDTO `json:"tagTranslation"`
	Thumbnails     thumbnailDTO      `json:"thumbnails"`
}

type WatchListDTO struct {
	Page           pageDTO           `json:"page"`
	TagTranslation tagTranslationDTO `json:"tagTranslation"`
	Thumbnails     thumbnailDTO      `json:"thumbnails"`
	IllustSeries   []illustSeriesDTO `json:"illustSeries"`
	NovelSeries    []novelSeriesDTO  `json:"novelSeries"`
	Users          []userDTO         `json:"users"`
}
