import { Image } from './image';
export class NewProduct{
  name : String;
  category: string;
  price : Number;
  description : String;
  images : String[];
  currency : Number;
  available : Number;
  constructor(name : String, category: string, price : Number, description : String, images : String[], cur : Number, a : Number){
    this.name = name;
    this.price = price;
    this.category = category;
    this.description = description;
    this.images = images;
    this.currency = cur;
    this.available = a;
  }
}
