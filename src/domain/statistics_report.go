package domain


type StatisticsReport struct {
	CampaignId string `json:"campaignId" xml:"campaign_id"`
	Description string `json:"description" xml:"description"`
	AdvertisedLinks []string `json:"advertisedLinks" xml:"advertised_links"`
	Clicks []ClickEvent `json:"clicks" xml:"clicks"`
	AdvertisingCount AdvertisingCount `json:"advertisingCount" xml:"advertising_count"`
	NumOfLikes int `json:"numOfLikes" xml:"num_of_likes"`
	NumOfDislikes int `json:"numOfDislikes" xml:"num_of_dislikes"`
	NumOfComments int `json:"numOfComments" xml:"num_of_comments"`
	CampaignType string `xml:"campaign_type" json:"campaignType"`
	CampaignPeriod string `xml:"campaign_period" json:"campaignPeriod"`
	AdvertisementFrequency string `xml:"advertisement_frequency" json:"advertisementFrequency"`
}
