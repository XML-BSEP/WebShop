import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject } from 'rxjs';
import { AuthenticatedUser } from '../model/authenticatedUser';
import { AuthenticationService } from '../service/authentication/authentication.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  userLoggedIn : Boolean;
  hasCA: boolean = true;
  currentUserSubject;
  constructor(private router : Router, private authenticationService : AuthenticationService) { }

  ngOnInit(): void {
    console.log(localStorage.getItem('userId'));
    if(localStorage.getItem('userId')!==null){
      this.userLoggedIn = true;

    }else{
      this.userLoggedIn=false;
    }

  }
  isAdminLoggedIn() : boolean {
    this.currentUserSubject = new BehaviorSubject<AuthenticatedUser>(JSON.parse(localStorage.getItem('currentUser')));
    if(this.currentUserSubject.value.role == "Admin") {
      return true;
    }
    return false;
  }


  login(){
    this.router.navigate(['/login']);
  }

  logout(){
    this.authenticationService.logout();
    this.router.navigate(['/login']);
  }

}
