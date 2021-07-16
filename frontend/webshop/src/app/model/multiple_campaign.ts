import { ShowAd } from "./show_ads";

export class MultipleCampaign {
    public id : String;
	public startDate : Date;
	public endDate : Date;
	public frequency : Number;
	public ads : ShowAd[];
	public type : Number;

    constructor() {}
}