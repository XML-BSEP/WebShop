import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { ItemToCart } from './../model/itemToCart';
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
  constructor(private toastr : ToastrService,private shopService :ShopService, private router: Router) { }

  ngOnInit(): void {
    this.addressGroup = new FormGroup({
      'address' : new FormControl(null, Validators.required),
      'zip' : new FormControl(null, [Validators.required,Validators.pattern(/^-?(0|[1-9]\d*)?$/)]),
      'city' : new FormControl(null, Validators.required),
      'state' : new FormControl(null, Validators.required),
    })
    this.curUsr = JSON.parse(localStorage.getItem('currentUser'))

    this.getCart()
  }
  calculateTotalPrice(){
    for(let i=0; i<this.products.length;i++){
      this.totalPrice += this.products[i].price
    }
  }
  removeFromCart(product){
    var curUsr = JSON.parse(localStorage.getItem('currentUser'))
    let userId = Number(curUsr.id)
    let item = new ItemToCart(userId, product.id)
    console.log(item)
    this.shopService.removeFromCart(item).subscribe(
      res=>{
        this.toastr.success('Success');
        location.reload();
      },
      err=>{
        this.toastr.error(err)
      }
  )
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
    this.myShoppingCart = new ShoppingCart(this.addressGroup.controls.address.value, Number(this.addressGroup.controls.zip.value), this.addressGroup.controls.city.value, this.addressGroup.controls.state.value, this.curUsr.id)
    console.log(this.myShoppingCart)


    this.shopService.placeOrder(this.myShoppingCart).subscribe(
        success => {
          this.toastr.success("Order successfully placed :D")
          this.router.navigate(['/home']);
        },
        error => {
          this.toastr.error(error)
        }

      )
  }
}
