import { error } from 'selenium-webdriver';
import { ShopService } from './../service/shop/shop.service';
import { Product } from 'src/app/model/product';
import { ProductForCart } from './../model/productForCart';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { ShoppingCart } from '../model/shoppingCart';
import { UserDTO } from '../model/userDTO';

@Component({
  selector: 'app-shopping-cart',
  templateUrl: './shopping-cart.component.html',
  styleUrls: ['./shopping-cart.component.css']
})
export class ShoppingCartComponent implements OnInit {
  addressGroup : FormGroup
  myShoppingCart : ShoppingCart
  product1 : ProductForCart
  product2: ProductForCart
  products : ProductForCart[]
  totalPrice : number = 0;
  curUsr;
  empty : boolean
  constructor(private shopService :ShopService) { }

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

    this.products=[this.product1, this.product2]
    this.getCart()
  }
  calculateTotalPrice(){
    for(let i=0; i<this.products.length;i++){
      this.totalPrice += this.products[i].price
    }
  }
  getCart(){
    var curUsr = JSON.parse(localStorage.getItem('currentUser'))
    let userId = Number(curUsr.id)
    console.log(userId)
    this.shopService.getCart(new UserDTO(userId)).subscribe(
      res => {
        this.products = res
        console.log(res)
        if(res==null){
          this.empty=true;
        }else{
          this.empty=false;
        }
        this.calculateTotalPrice()

      },error=>{
        console.log('error')
      }
      )
  }
  checkout(){
    // this.myShoppingCart = new ShoppingCart(this.products, this.addressGroup.controls.address.value, this.addressGroup.controls.zip.value, this.addressGroup.controls.city.value, this.addressGroup.controls.state.value, this.totalPrice, this.curUsr.id)
    console.log(this.myShoppingCart)
  }
}
