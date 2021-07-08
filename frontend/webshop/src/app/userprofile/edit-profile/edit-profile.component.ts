import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { SaveToken } from 'src/app/model/save_token';
import { StatisticsReport } from 'src/app/model/statisticsreport';
import { AgentService } from 'src/app/service/agent_service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {

  constructor(private agentService : AgentService, private toastr : ToastrService) { }
  public generateToken : boolean = false;
  public tokenGenerated : boolean = false;
  public statistics : StatisticsReport[] = [];
  public statisticsReport : boolean = false;
  public personalToken : String = "";
  ngOnInit(): void {

  }
  goToGenerateToken() {
    this.tokenGenerated = true;
    this.generateToken = true;
  }

  saveToken() {
    if (this.personalToken === "") {
      this.toastr.info("Please enter your token")
    } else {
      let token = new SaveToken();
      token.token = this.personalToken
      this.agentService.saveToken(token).subscribe(
        res => {
          this.toastr.success("Successfully saved!")
        }, err=> {
          this.toastr.error("Error")
        }
      )
    }

  }

  generateStatisticsReport() {
    this.tokenGenerated = false;
    this.generateToken = false;
    this.statisticsReport = true;
    this.agentService.getStatisticsReport().subscribe(
      res => {
        this.statistics = res;
      }
    )
  }

}
