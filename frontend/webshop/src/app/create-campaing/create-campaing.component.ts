import { DatePipe } from '@angular/common';
import { Component, OnInit, AfterViewInit, ViewChild} from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from 'ngx-toastr';
import { ShowImageComponent } from 'src/app/dialogs/show-image/show-image.component';
import { CreateDisposableCampaign } from 'src/app/model/create_disposable_campaign';
import { CreateMultipleCampaign } from 'src/app/model/create_multiple_campaign';
import { ShowAd } from 'src/app/model/show_ads';
import { AgentService } from 'src/app/service/agent_service';

@Component({
  selector: 'app-create-campaing',
  templateUrl: './create-campaing.component.html',
  styleUrls: ['./create-campaing.component.css']
})
export class CreateCampaingComponent implements OnInit,AfterViewInit {
  isLinear = true;
  campaignChecked = "";
  exposeDateDisposableCampaing  = new Date();
  today = new Date();
  public startTimedisposableCampaignTime : String = '00:00';
  public disposableCampaignTime : string = '00:00';
  public adsForCampaign : ShowAd[];
  exposeStartDateMultipleCapaign = new Date();
  exposeEndDateMultipleCapaign = new Date();
  public ads : ShowAd[];
  public multipleCampaignStartTime : string = '00:00';
  public multipleCampaignEndTime : string = '00:00';
  public sinImputarValue : string = 'postChecked';
  public frequency : Number = 0;
  

  constructor(private agentService : AgentService, private dialog : MatDialog, private datePipe : DatePipe, private toastr : ToastrService) { }

  ngOnInit(): void {
    this.adsForCampaign = [];
    this.agentService.getAllAdsPerAgent().subscribe(
      res => {
        this.ads = [];
        for(let a of res) {
          a.isAdded = false;
          if (a.type === 0) {
            a.mediaType = "Photo"
          } else {
            a.mediaType = "Video"
          }
          this.ads.push(a)
        }
      }
    )
  }

  ngAfterViewInit() {}
 
  showImage(a : ShowAd) {
    const dialogRef = this.dialog.open(ShowImageComponent, {
      width: '35vw',
      height: '90vh',
      data: a.media
    });
  }
  addToCampaign(a : ShowAd) {
    if(a.isAdded) {
      this.adsForCampaign.forEach((element, index)=>{
        if(element.id === a.id) this.adsForCampaign.splice(index,1);
     });
    } else {
      this.adsForCampaign.push(a)
    }

    a.isAdded = !a.isAdded    
  }
  onTimeDisposableCampaignChange(time) {
    this.disposableCampaignTime = time;
    console.log(this.disposableCampaignTime)
    console.log(this.exposeDateDisposableCampaing)
  }

  onStartTimeMulitpleCampaignChange(time) {
    this.disposableCampaignTime = time;
    console.log(this.multipleCampaignStartTime)
    console.log(this.exposeStartDateMultipleCapaign)
  }
  addDisposableCampaign() {

    let campaign = new CreateDisposableCampaign();

  

    campaign.ads = this.adsForCampaign;
    if (campaign.ads.length == 0) {
      this.toastr.error("Choose at leats one ad")
      return
    }
    if (this.sinImputarValue === "storyChecked") {
      campaign.type = 0;
    } else {
    campaign.type = 1;
    }

    let latest_date =this.datePipe.transform(this.exposeDateDisposableCampaing, 'yyyy-MM-dd');
    latest_date = latest_date + " " + this.disposableCampaignTime;
    campaign.exposureDate = new Date(latest_date)
    console.log(campaign);

    this.agentService.createDisposableCampaign(campaign).subscribe(
      res => {
        this.toastr.success("Successfully added")
      }, err => {
        
        this.toastr.error("Error")
      }
    )

  }

  addMultipleCampaign() {
    console.log(this.adsForCampaign)
    let campaign = new CreateMultipleCampaign();
    let latest_date =this.datePipe.transform(this.exposeStartDateMultipleCapaign, 'yyyy-MM-dd');
    latest_date = latest_date + " " + this.multipleCampaignStartTime;
    campaign.startDate = new Date(latest_date)

    let latest_date_end =this.datePipe.transform(this.exposeEndDateMultipleCapaign, 'yyyy-MM-dd');
    latest_date = latest_date + " " + this.multipleCampaignEndTime;
    campaign.endDate = new Date(latest_date_end)

    campaign.frequency = this.frequency;
    if (this.sinImputarValue === "storyChecked") {
      campaign.type = 0;
    } else {
    campaign.type = 1;
    }
    campaign.ads = this.adsForCampaign;
    if (campaign.ads.length == 0) {
      this.toastr.error("Choose at leats one ad")
      return
    }
    


    this.agentService.createMultipleCampaign(campaign).subscribe(
      res => {
        this.toastr.success("Successfully added")
      }, err => {
        
        this.toastr.error("Error")
      }
    )

  }

  onEndTimeTimeMulitpleCampaignChange(time) {
    this.disposableCampaignTime = time;
    console.log(this.multipleCampaignEndTime)
    console.log(this.exposeEndDateMultipleCapaign)
  }

}
