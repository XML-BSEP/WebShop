import { Component } from '@angular/core';
import { AuthenticatedUser } from './model/authenticatedUser';
import { AuthenticationService } from 'src/app/service/authentication/authentication.service';
import { Role } from './model/role';
import { AuthGuard } from './helpers';
import { Router } from '@angular/router';
import { ThrowStmt } from '@angular/compiler';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'webshop';

  user : AuthenticatedUser
  router : Router
  showHeader : boolean
  constructor(private authService : AuthenticationService, private _router : Router) {
    this.router = _router
  }



  public isNotLogged() {
    return !this.authService.getUserValue();
  }

  public isAdmin() {
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Admin;
  }

  
  public isCustomer() {
    if (this.router.url === '/forbidden') {
      return false;
    } 
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Customer;
  }

  
}
