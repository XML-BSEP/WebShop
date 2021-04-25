import { Image } from './image';
export class NewProduct{
  name : String;
  category: String;
  price : Number;
  description : String;
  images : String[];
  currency : Number;
  constructor(name : String, category: String, price : Number, description : String, images : String[], cur : Number){
    this.name = name;
    this.price = price;
    this.category = category;
    this.description = description;
    this.images = images;
    this.currency = cur;
  }
}
