export class Person{
  public name : String;
  public lastname : String;
  public phone : String;
  public email: String;

}


export class Account{
  public username : string;
  public password : String;
  public question : String;
  public answer : String;
}


export class NewUser{
  public account : Account;
  public person : Person;
}
