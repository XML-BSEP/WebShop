export class Product {
    name : String;
    price : number;
    image : String;
    currency : String;
    category : String;
    available : number;
    description : String;

    constructor(name : String, price : number, image : String, currency : String, category : String, available : number, description : String) {
        this.name = name;
        this.price = price;
        this.image = image;
        this.currency = currency;
        this.category = category;
        this.available = available;
        this.description = description;
    }
}