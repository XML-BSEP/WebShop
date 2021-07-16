import { ShowAd } from "./show_ads";

export class DisposableCampaign {
    public id : String;
	public exposureDate : Date;
	public ads : ShowAd[];
	public type : Number;

    constructor() {}
}