import { DatePipe } from '@angular/common';
import { ThrowStmt } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from 'ngx-toastr';
import { ShowImageComponent } from 'src/app/dialogs/show-image/show-image.component';
import { DisposableCampaign } from 'src/app/model/disposable_campaign';
import { MultipleCampaign } from 'src/app/model/multiple_campaign';
import { AgentService } from 'src/app/service/agent_service';

@Component({
  selector: 'app-change-campaign',
  templateUrl: './change-campaign.component.html',
  styleUrls: ['./change-campaign.component.css']
})
export class ChangeCampaignComponent implements OnInit {

  public isDisposableShowing : Boolean = false;
  public isMultipleShowing : Boolean = false;
  public disposableCampaigns : DisposableCampaign[];
  public multipleCampaigns : MultipleCampaign[];
  public multipleStartTime : String;
  public multipleEndTime : String;
  public today : Date = new Date();


  constructor(private agentService : AgentService, public toastr : ToastrService, private dialog : MatDialog, private datePipe : DatePipe) { }

  ngOnInit(): void {
  }

  clickShowDisposable() {
    this.isDisposableShowing = true;
    this.isMultipleShowing = false;

    this.agentService.getAllDisposableCampaigns().subscribe(
      res => {
        this.disposableCampaigns = res;   
        console.log(this.disposableCampaigns)
      }, 
      error => {
        this.toastr.error(error);
      }
      
    );
  }

  clikcShowMultiple() {
    this.isDisposableShowing = false;
    this.isMultipleShowing = true;

    this.agentService.getAllMultipleCampaigns().subscribe(
      res => {
        this.multipleCampaigns = res;
        console.log(this.multipleCampaigns)
      },
      err => {
        this.toastr.error(err);
      }
    );
  }

  changeCampaing(c) {
    console.log(c);
  }

  showImageDisposable(a : DisposableCampaign) {
    console.log(a)
    for(var i = 0; i < a.ads.length; i++) {
      const dialogRef = this.dialog.open(ShowImageComponent, {
        width: '35vw',
        height: '90vh',
        data : a.ads[i].media
      });
    }

  }

  showImageMultiple(a : MultipleCampaign) {
   console.log(a)
    for(var i = 0; i < a.ads.length; i++) {
      const dialogRef = this.dialog.open(ShowImageComponent, {
        width: '35vw',
        height: '90vh',
        data : a.ads[i].media
      });
    }

  }

  deleteDisposableCampaign(d) {
    this.agentService.deleteDisposableCampaign(d).subscribe(
      res => {
        this.toastr.success("Deleted <3")
        location.reload();
      },
      err => {
        this.toastr.error(err);
      }
      
    );
  }

  deleteMultipleCampaign(d) {
    this.agentService.deleteMultipleCampaign(d).subscribe(
      res => {
        this.toastr.success("Deleted <3")
        location.reload();
      },
      err => {
        this.toastr.error(err);
      }
      
    );
  }



  changeMultipleCampaign(d : MultipleCampaign, startTime : any, endTime : any) {
   /* let latest_date =this.datePipe.transform(this.exposeDateDisposableCampaing, 'yyyy-MM-dd');
    latest_date = latest_date + " " + this.disposableCampaignTime;*/

    console.log(startTime)
    console.log(endTime)
    let startDate = this.datePipe.transform(d.startDate, 'yyyy-MM-dd');
    startDate = startDate + " " + startTime;

    let endDate = this.datePipe.transform(d.endDate, 'yyyy-MM-dd');
    endDate = endDate + " " + endTime;

    
    d.startDate = new Date(startDate);
    d.endDate = new Date(endDate);

    console.log(d.startDate)
    console.log(d.endDate)
    
    this.agentService.updateMultipleCampaign(d).subscribe(
      res => {
        this.toastr.success("Updated <3");
        location.reload();
      }, 
      err => {
        this.toastr.error(err);
      }
    );
    


    
  }

 
  

}
