import { Product } from 'src/app/model/product';
import { ProductForCart } from './../model/productForCart';
import { ProductInCart } from './../model/productInCart';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { ShoppingCart } from '../model/shoppingCart';

@Component({
  selector: 'app-shopping-cart',
  templateUrl: './shopping-cart.component.html',
  styleUrls: ['./shopping-cart.component.css']
})
export class ShoppingCartComponent implements OnInit {
  addressGroup : FormGroup
  myShoppingCart : ShoppingCart
  products : ProductInCart[]
  product1 : ProductForCart
  product2: ProductForCart
  totalPrice : number = 0;
  curUsr;
  constructor() { }

  ngOnInit(): void {
    this.addressGroup = new FormGroup({
      'address' : new FormControl(null, Validators.required),
      'zip' : new FormControl(null, [Validators.required,Validators.pattern(/^-?(0|[1-9]\d*)?$/)]),
      'city' : new FormControl(null, Validators.required),
      'state' : new FormControl(null, Validators.required),
    })

    this.curUsr = JSON.parse(localStorage.getItem('currentUser'))
    this.product1 = new ProductForCart(1, "Product 1", 660, 'https://i.imgur.com/o2fKskJ.jpg', 'nzm')
    this.product2 = new ProductForCart(2, "Product 2", 666, 'https://i.imgur.com/GQnIUfs.jpg', 'nzm stvarno')

    this.products=[new ProductInCart(this.product1, 1), new ProductInCart(this.product2, 2)]
    this.calculateTotalPrice()


  }
  calculateTotalPrice(){
    for(let i=0; i<this.products.length;i++){
      this.totalPrice += this.products[i].product.price
    }
  }
  checkout(){
    this.myShoppingCart = new ShoppingCart(this.products, this.addressGroup.controls.address.value, this.addressGroup.controls.zip.value, this.addressGroup.controls.city.value, this.addressGroup.controls.state.value, this.totalPrice, this.curUsr.id)
    console.log(this.myShoppingCart)
  }
}
