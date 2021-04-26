import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators, FormControl} from '@angular/forms';
import { ResetPasswordService} from '../service/reset-password/reset-password.service'
import { ResetMail } from '../model/resetMail';
import { ResetPass } from '../model/resetPass';
import { Router } from '@angular/router';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent implements OnInit {

  isLinear = true;
  firstFormGroup: FormGroup;
  secondFormGroup: FormGroup;
  thirdFormGroup : FormGroup;
  resetMail : ResetMail;
  code : string;
  resetPass : ResetPass;
  email : string;

  constructor(private _formBuilder: FormBuilder, private resetPasswordSerivce : ResetPasswordService, private router : Router) { }

  ngOnInit() {
    this.firstFormGroup = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email])
    });
    this.secondFormGroup = new FormGroup({
      code: new FormControl('', Validators.required)
    }); 
    this.thirdFormGroup = new FormGroup({
      password: new FormControl(null, [Validators.required, Validators.pattern('(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&].{7,}')]),
      confirmPassword: new FormControl(null, [Validators.required, Validators.pattern('(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&].{7,}')])
    });
  }

  sendMail(){
    this.email = this.firstFormGroup.controls.email.value;
    console.log(this.email)
    this.resetMail = new ResetMail(this.email)
    
    this.resetPasswordSerivce.resetPasswordMail(this.resetMail).subscribe(
      res=>{
        alert('Check your mail');
      },
      error=>{
        alert("Fail - email is not in use!");
      }
      )

  }

  verifyCode() {
    this.code = this.secondFormGroup.controls.code.value;
    console.log(this.code) 
  
  }

  
  resetPassword() {
    var password = this.thirdFormGroup.controls.password.value;
    console.log(password)

    var confirmPassword = this.thirdFormGroup.controls.confirmPassword.value;
    console.log(confirmPassword)

    this.resetPass = new ResetPass(this.email, password, confirmPassword, this.code)

    this.resetPasswordSerivce.resetPassword(this.resetPass).subscribe(
      res=>{
        alert("Successfully changed password");
        this.router.navigate(['/login']);
      },
      error=>{
        alert("Check your code or passwords");
      }
      )

  }

}
