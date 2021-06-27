export class FilterSearch {
    userId : number
    name : String;
    category : String;
    priceRangeStart : number;
    priceRangeEnd : number
    limit : number;
    offset : number;
    order : string;

    constructor(userId :number, name : String, category : String, priceRangeStart : number, priceRangeEnd : number, limit : number, offset : number, order : string) {
        this.userId = userId
        this.name = name;
        this.category = category;
        this.priceRangeStart = priceRangeStart;
        this.priceRangeEnd = priceRangeEnd;
        this.limit = limit;
        this.offset = offset;
        this.order = order;

    }
}
