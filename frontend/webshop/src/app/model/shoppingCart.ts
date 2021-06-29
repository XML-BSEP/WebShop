import { Product } from 'src/app/model/product';
export class ShoppingCart{
  address : string
  zip : Number
  city : string
  state : string
  totalPrice: number
  userId : Number;
  constructor(address : string, zip :Number, city : string, state : string, totalPrice : number, userId : Number){
    this.address = address
    this.zip = zip
    this.city = city
    this.state = state
    this.totalPrice = totalPrice
    this.userId = userId
  }
}
