package dtos

type UserInfoDTO struct {
	UserID     int64       `json:"userId,string"`
	Name       string      `json:"name"`
	Avatar     string      `json:"imageBig"`
	IsFollowed bool        `json:"isFollowed"`
	Following  int32       `json:"following"`
	Region     userDataDTO `json:"region"`
	Gender     userDataDTO `json:"gender"`
}

type userDataDTO struct {
	Name         string `json:"region"`
	PrivacyLevel string `json:"privacyLevel"`
}
