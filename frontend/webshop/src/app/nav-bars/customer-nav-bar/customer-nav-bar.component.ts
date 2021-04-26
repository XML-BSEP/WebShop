import { Component, OnInit } from '@angular/core';
import { AuthenticationService } from './../../service/authentication/authentication.service';

@Component({
  selector: 'app-customer-nav-bar',
  templateUrl: './customer-nav-bar.component.html',
  styleUrls: ['./customer-nav-bar.component.css']
})
export class CustomerNavBarComponent implements OnInit {

  constructor(private authService : AuthenticationService) { }

  ngOnInit(): void {
  }
  logout(){
    this.authService.logout();
  }
}
