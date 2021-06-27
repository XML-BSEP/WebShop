import { ProductForCart } from './productForCart';
export class ProductInCart{
  product : ProductForCart
  quantity : Number

  constructor(product : ProductForCart, quant : Number){
    this.product = product
    this.quantity = quant
  }
}
