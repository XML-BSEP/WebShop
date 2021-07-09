import { ShopService } from './../service/shop/shop.service';
import { ItemToCart } from './../model/itemToCart';
import { ProductServiceService } from './../service/product/product-service.service';
import { Router, ActivatedRouteSnapshot, ActivatedRoute } from '@angular/router';
import { Role } from './../model/role';
import { PageEvent } from '@angular/material/paginator';
import { Product } from './../model/product';
import { FilterSearch } from './../model/filterSearch';
import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { DeletedProduct } from '../model/deletedProduct';
import { AuthenticationService } from '../service/authentication/authentication.service';

@Component({
  selector: 'app-shop-home',
  templateUrl: './shop-home.component.html',
  styleUrls: ['./shop-home.component.css']
})
export class ShopHomeComponent implements OnInit {
  products : Product[]
  itemsCount : number;
  pageSize : number;
  pageSizeOptions = [5, 10, 25]
  filterSearch : FilterSearch
  shopid : number
  shopname : string
  constructor(
    private shopService : ShopService,
    private route : ActivatedRoute,
    private productService : ProductServiceService,
    private router : Router,
    private toastr : ToastrService,
    private prodService : ProductServiceService,
    private authService :AuthenticationService) { }

  ngOnInit(): void {
    this.route.queryParams
    .subscribe(params => {
      this.shopid = params.id;
      this.shopname = params.name
    });
    console.log(this.shopname)
    this.pageSize = 5;
    this.filterSearch = new FilterSearch(Number(this.shopid),"", "", 0, 10000000, this.pageSize, 0, "price asc")

    this.getShopProducts(this.filterSearch)
  }
  public isAdmin() {
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Admin;
  }

  isMine(product : Product) {
    return product.userId === this.authService.currentUserValue.id
  }

  getShopProducts(filterSearch : FilterSearch) {
      this.prodService.getProducts(filterSearch).subscribe(
        data => {
          this.products = []
          console.log("Pre: " + this.products.length)
          this.products = data
          console.log("Posle: " + this.products.length)
          this.itemsCount = data[0].count
        })
  }

  load(event : PageEvent) {
      this.pageSize = event.pageSize
      this.filterSearch.category = ""
      this.filterSearch.name = ""
      this.filterSearch.offset = event.pageIndex * this.pageSize
      console.log("Velicina stranice: " + this.pageSize)
      this.filterSearch.limit = this.pageSize
      this.filterSearch.priceRangeStart = 0
      this.filterSearch.priceRangeEnd = 1000000
      this.filterSearch.order = "price asc"
      this.getShopProducts(this.filterSearch)
      console.log(event.pageSize)
  }
  edit(product){
    console.log(product.currency)
    this.router.navigate(['/editProduct'], {state: {data: product}});

  }
  remove(product){
    console.log(product)
    // this.router.navigate(['/editProduct'], {state: {data: product}});

    var deletedProduct = new DeletedProduct(product.serial.toString(), this.authService.currentUserValue.id)

    console.log(deletedProduct);

    this.productService.deleteProduct(deletedProduct).subscribe(
      res=>{
        this.toastr.success('Success');
        this.router.navigate(['/products'])
      },
      err=>{
        this.toastr.error(err)
      }
        )


  }
  addToCart(item){
    console.log(item)
    var curUsr = JSON.parse(localStorage.getItem('currentUser'))
    let userId = Number(curUsr.id)
    var itemToCart = new ItemToCart(userId, item.productId);

      this.shopService.addToCart(itemToCart).subscribe(
        success => {
          this.toastr.success("Item added to cart")
        },
        error => {
          this.toastr.error(error)
        }

      )
  }
}
