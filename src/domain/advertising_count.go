package domain

type AdvertisingCount struct {
	CampaignId string `json:"campaignId" xml:"campaign_id"`
	AgentId string `json:"agentId" xml:"agent_id"`
	AdvertisedCount int `json:"advertisedCount" xml:"advertised_count"`
}
