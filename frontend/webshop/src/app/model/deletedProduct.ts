export class DeletedProduct{
  serial : String;
  public userId : Number;
  constructor(serial : String, userId : Number){
    this.serial = serial;
    this.userId = userId;
  }
}
