import { Image } from "./image"

export class ProductForCart{
  id : Number
  name : String
  price: number
  picture : String
  description : String
  constructor(id : Number, name :String, price : number, picture : String, description : String){
    this.id = id
    this.name = name
    this.price = price
    this.picture = picture
    this.description = description
  }

}
