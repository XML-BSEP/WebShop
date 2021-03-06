import { TOUCH_BUFFER_MS } from '@angular/cdk/a11y';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { Authentication } from '../model/authentication';
import { Role } from '../model/role';
import { AuthenticationService } from '../service/authentication/authentication.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public loginForm: FormGroup;
  username;
  password;
  credentials;
  constructor(private router : Router, private authService : AuthenticationService, private toastr : ToastrService) { }

  ngOnInit(): void {

    if(localStorage.getItem('userId')!==null){
      this.router.navigate(['/home']);
    }

    this.loginForm = new FormGroup({
      // 'username' : new FormControl(null, [Validators.required, Validators.email]),
      'username' : new FormControl(null, [Validators.required]),

      'password' : new FormControl(null, [Validators.required])
      // 'password' : new FormControl(null, [Validators.required,Validators.pattern('^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z0-9\d$@$!%*?&].{7,}$')])
    });
  }

  back(){
    this.router.navigate(['/home']);
  }

  login(){
    this.username = this.loginForm.controls.username.value;
    this.password = this.loginForm.controls.password.value;
    console.log(this.username);
    this.credentials = new Authentication(this.username.toLowerCase(), this.password);
    console.log(this.credentials);
    this.authService.login(this.credentials).subscribe(
      result => {
        localStorage.setItem('userId',String(result.id))
        localStorage.setItem('role',String(result.role))
        this.router.navigate(['/home'])


      },
      error=>{
        this.toastr.error(error)
        this.router.navigate(['/login']);

      });
  }


}
