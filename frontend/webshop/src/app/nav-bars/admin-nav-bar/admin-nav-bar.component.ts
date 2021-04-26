import { AuthenticationService } from './../../service/authentication/authentication.service';
import { Component, OnInit } from '@angular/core';
import * as mdb from 'mdb-ui-kit';

@Component({
  selector: 'app-admin-nav-bar',
  templateUrl: './admin-nav-bar.component.html',
  styleUrls: ['./admin-nav-bar.component.css']
})
export class AdminNavBarComponent implements OnInit {

  constructor(private authService : AuthenticationService) { }

  ngOnInit(): void {
  }

  logout(){
    this.authService.logout();
  }
}
