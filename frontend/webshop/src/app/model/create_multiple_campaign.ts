import { ShowAd } from "./show_ads";

export class CreateMultipleCampaign {
    public startDate : Date;
    public endDate : Date;
    public frequency : Number;
    public type : Number;
    public ads : ShowAd[];
    public isAdded : boolean;

    constructor() {}
}