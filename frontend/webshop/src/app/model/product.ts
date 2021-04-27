export class Product {
    name : String;
    price : number;
    image : String;
    currency : String;
    category : String;
    available : number;
    description : String;
    count : number;

    constructor(name : String, price : number, image : String, currency : String, category : String, available : number, description : String, count : number) {
        this.name = name;
        this.price = price;
        this.image = image;
        this.currency = currency;
        this.category = category;
        this.available = available;
        this.description = description;
        this.count = count;
    }
}