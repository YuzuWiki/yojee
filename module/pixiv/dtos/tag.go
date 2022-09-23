package dtos

import (
	"encoding/json"
)

// tagName, string alise
type tagName string

// TageDTO  https://www.pixiv.net/ajax/search/tags/%E4%BA%8C%E6%AC%A1%E5%89%B5%E4%BD%9C?lang=zh
type TageDTO struct {
	Tag  string `json:"tag"`
	Word string `json:"word"`
	// tag digests , if exist
	Digest tagDigestDTO `json:"pixpedia"`
	// translation, if exist ( tag_name: translation )
	Transl tagTranslDTO `json:"tagTranslation"`
}

type tagDigestDTO struct {
	Id       int64  `json:"id,string"`
	Abstract string `json:"abstract"`
	Image    string `json:"image"`
	// parent tag
	Parent string `json:"parentTag"`
}

func (dto *tagDigestDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := tagDigestDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type multilingualDTO struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

type tagTranslDTO map[tagName]multilingualDTO

func (dto *tagTranslDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[tagName]multilingualDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}
