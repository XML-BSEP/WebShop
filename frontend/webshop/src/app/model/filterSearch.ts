export class FilterSearch {
    name : String;
    category : String;
    priceRangeStart : number;
    priceRangeEnd : number
    limit : number;
    offset : number;
    order : string;

    constructor(name : String, category : String, priceRangeStart : number, priceRangeEnd : number, limit : number, offset : number, order : string) {
        this.name = name;
        this.category = category;
        this.priceRangeStart = priceRangeStart;
        this.priceRangeEnd = priceRangeEnd;
        this.limit = limit;
        this.offset = offset;
        this.order = order;
        
    }
}