import { StatisticsReport } from "./statisticsreport";

export class CampaignsReport {
    public timestamp : Date;
    public report_id : String;
    public statistic_report: StatisticsReport[];

    constructor() {}
}