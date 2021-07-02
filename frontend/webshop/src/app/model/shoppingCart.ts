import { Product } from 'src/app/model/product';
export class ShoppingCart{
  address : string
  zip : Number
  city : string
  state : string
  userId : Number;
  constructor(address : string, zip :Number, city : string, state : string, userId : Number){
    this.address = address
    this.zip = zip
    this.city = city
    this.state = state
    this.userId = userId
  }
}
