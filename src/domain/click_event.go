package domain

type ClickEvent struct {
	CampaignId         string `json:"campaignId" xml:"campaign_id"`
	InfluencerId       string `json:"influencerId" xml:"influencer_id"`
	NumOfClicks        int    `json:"clicks" xml:"num_of_clicks"`
	InfluencerUsername string `json:"influencerUsername" xml:"influencer_username"`
}
