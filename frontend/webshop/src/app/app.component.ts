import { Component } from '@angular/core';
import { AuthenticatedUser } from './model/authenticatedUser';
import { AuthenticationService } from 'src/app/service/authentication/authentication.service';
import { Role } from './model/role';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'webshop';

  user : AuthenticatedUser

  constructor(private authService : AuthenticationService) {}



  public isNotLogged() {
    return !this.authService.getUserValue();
  }

  public isAdmin() {
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Admin;
  }

  
  public isCustomer() {
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Customer;
  }


}
