import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-registration-confirmation',
  templateUrl: './registration-confirmation.component.html',
  styleUrls: ['./registration-confirmation.component.css']
})
export class RegistrationConfirmationComponent implements OnInit {

  constructor() { }
  public codeForm : FormGroup;
  ngOnInit(): void {
    this.codeForm = new FormGroup({
      //maybe add pattern for code validation on frontend
      'code' : new FormControl(null, [Validators.required])
    });

  }
  confirm(){

  }
  resend(){

  }
}
