package dtos

import "encoding/json"

// TopProfileDTO return user's  profile (top)
type TopProfileDTO struct {
	Illusts   illustMapDTO `json:"illusts"`
	Manga     mangaMapDTO  `json:"manga"`
	Novels    novelMapDTO  `json:"novels"`
	ExtraData extraDataDTO `json:"extra_data"`
}

type extraDataDTO struct {
	Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Canonical   string `json:"canonical"`
		Ogp         struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"ogp"`
		Twitter struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Card        string `json:"card"`
		} `json:"twitter"`
		DescriptionHeader string `json:"descriptionHeader"`
	} `json:"meta"`
}

// AllProfileDTO return user's  profile (all)
type AllProfileDTO struct {
	Illusts IllustMapDTO `json:"illusts"`
	Manga   MangaMapDTO  `json:"manga"`
	Novel   NovelMapDTO  `json:"novels"`
}

type IllustMapDTO map[string]struct{}

func (dto *IllustMapDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]struct{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type MangaMapDTO map[string]struct{}

func (dto *MangaMapDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]struct{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type NovelMapDTO map[string]struct{}

func (dto *NovelMapDTO) UnmarshalJSON(body []byte) error {
	if len(body) < 5 {
		return nil
	}

	data := map[string]struct{}{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}
