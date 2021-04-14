import { timestamp } from "rxjs/operators";

export class RegisteredUser{
  public name : String;
  public surname : String;
  public email : String;
  public question : String;
  public answer : String;
  public password : String;
  public confirmedpassword : String;
  public username : String;

  constructor(name : String, surname : String, email : String, question : String, answer : String, password : String, confirmedpassword : String, username : String){
    this.name = name;
    this.surname = surname;
    this.email = email;
    this.question = question;
    this.answer = answer;
    this.password = password;
    this.confirmedpassword = confirmedpassword;
    this.username = username;
  }
}
