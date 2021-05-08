export class Product {
    name : String;
    price : number;
    image : String[];
    currency : String;
    category : String;
    available : number;
    description : String;
    count : number;
    serial : Number;
    constructor(name : String, price : number, image : String[], currency : String, category : String, available : number, description : String, count : number, serial : Number) {
        this.name = name;
        this.price = price;
        this.image = image;
        this.currency = currency;
        this.category = category;
        this.available = available;
        this.description = description;
        this.count = count;
        this.serial = serial;
    }
}
