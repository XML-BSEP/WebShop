import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { RegisteredUser } from '../model/registeredUser';
import { RegistrationService } from '../service/registration/registration.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  constructor(private router: Router,private registrationService : RegistrationService) { }
  public registrationForm: FormGroup;
  private newUser : RegisteredUser;
  ngOnInit(): void {
    this.registrationForm = new FormGroup({
    'name' : new FormControl(null, Validators.required),
    'lastname' : new FormControl(null, Validators.required),
    'username' : new FormControl(null, Validators.required),
    'email' : new FormControl(null, [Validators.required, Validators.email]),
    'password' : new FormControl(null, [Validators.required, Validators.pattern('(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&].{7,}')]),
    'confirmPassword' : new FormControl(null, [Validators.required,    Validators.pattern('(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&].{7,}')]),
  });
  }

  register(){
    var name = this.registrationForm.controls.name.value;
    var lastname = this.registrationForm.controls.lastname.value;
    var username = this.registrationForm.controls.username.value;
    var email = this.registrationForm.controls.email.value;
    var password = this.registrationForm.controls.password.value;
    var confirmPassword = this.registrationForm.controls.confirmPassword.value;
    console.log(password);
    if(password===confirmPassword){
      this.newUser = new RegisteredUser(name, lastname, email, password, confirmPassword, username);

    this.registrationService.register(this.newUser).subscribe(
      res=>{
        alert('Success');
        this.router.navigate(['/regconfirm'], {state: {data: email}});
      },
      error=>{
        alert(error)
      }
        )
    }

  }
}
