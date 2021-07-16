import { AdvertisingCount } from "./advertisingcount";
import { ClickEvent } from "./clickevent";

export class StatisticsReport {
    public CampaignId : String;
    public campaignId : String;
	public description : String;
	public advertisedLinks : String[];
	public clicks : ClickEvent[];
	public advertisingCount : AdvertisingCount;
	public numOfLikes : Number;
	public numOfDislikes : Number;
	public numOfComments : Number;

    constructor() {}
}