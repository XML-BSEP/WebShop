import { Product } from 'src/app/model/product';
import { ProductInCart } from './productInCart';
export class ShoppingCart{
  products : ProductInCart[]
  address : string
  zip : Number
  city : string
  state : string
  totalPrice: number
  userId : Number;
  constructor(products : ProductInCart[], address : string, zip :Number, city : string, state : string, totalPrice : number, userId : Number){
    this.products = products
    this.address = address
    this.zip = zip
    this.city = city
    this.state = state
    this.totalPrice = totalPrice
    this.userId = userId
  }
}
