import { Image } from './image';
export class NewProduct{
  name : String;
  category: string;
  price : Number;
  description : String;
  images : String[];
  available : Number;
  serial : String;
  constructor(name : String, category: string, price : Number, description : String, images : String[], a : Number, s : String){
    this.name = name;
    this.price = price;
    this.category = category;
    this.description = description;
    this.images = images;
    this.available = a;
    this.serial = s;
  }
}
