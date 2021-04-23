import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators, FormControl} from '@angular/forms';

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

  constructor(private _formBuilder: FormBuilder) { }

  ngOnInit() {
    this.firstFormGroup = new FormGroup({
      email: new FormControl('', Validators.required)
    });
    this.secondFormGroup = new FormGroup({
      code: new FormControl('', Validators.required)
    }); 
    this.thirdFormGroup = new FormGroup({
      password: new FormControl(null, Validators.required),
      confirmPassword: new FormControl(null, Validators.required)
    });
  }

  sendMail(){
    var email = this.firstFormGroup.controls.email.value;
    console.log(email)
  }

  verifyCode() {
    var code = this.secondFormGroup.controls.code.value;
    console.log(code)
  }

  
  resetPassword() {
    var password = this.thirdFormGroup.controls.password.value;
    console.log(password)

    var confirmPassword = this.thirdFormGroup.controls.confirmPassword.value;
    console.log(confirmPassword)
  }

}
