import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Authentication } from '../model/authentication';
import { Role } from '../model/role';
import { AuthenticationService } from '../service/authentication/authentication.service';

@Component({
  selector: 'app-login',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
  public loginForm: FormGroup;
  username;
  password;
  credentials;
  constructor(private router : Router, private authService : AuthenticationService) { }

  ngOnInit(): void {
    this.loginForm = new FormGroup({
      'username' : new FormControl(null, [Validators.required, Validators.email]),
      'password' : new FormControl(null, Validators.required)
    });
  }
}