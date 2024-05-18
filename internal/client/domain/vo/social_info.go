package vo

type SocialInfo struct {
	Instagram string `json:"social_info_instagram"`
	Linkedin  string `json:"social_info_linkedin"`
	Facebook  string `json:"social_info_facebook"`
}

func NewSocialInfo(instagram string, linkedin string, facebook string) (*SocialInfo, error) {
	return &SocialInfo{
		Instagram: instagram,
		Linkedin:  linkedin,
		Facebook:  facebook,
	}, nil
}
