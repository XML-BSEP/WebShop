import { DeletedProduct } from './../../model/deletedProduct';
import { ToastrModule, ToastrService } from 'ngx-toastr';
import { Category } from './../../model/category';
import { Router } from '@angular/router';
import { Role } from './../../model/role';
import { AuthenticationService } from './../../service/authentication/authentication.service';
import { ThrowStmt } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { FilterSearch } from 'src/app/model/filterSearch';
import { Product } from 'src/app/model/product';
import { ProductServiceService } from 'src/app/service/product/product-service.service';

@Component({
  selector: 'app-products-page',
  templateUrl: './products-page.component.html',
  styleUrls: ['./products-page.component.css']
})
export class ProductsPageComponent implements OnInit {

  products : Product[]
  itemsCount : number;
  pageSize : number;
  pageSizeOptions = [5, 10, 25]
  filterSearch : FilterSearch
  constructor(private productService : ProductServiceService, private router : Router,  private toastr : ToastrService, private prodService : ProductServiceService, private authService :AuthenticationService) { }

  ngOnInit(): void {
    // this.pageSize = 5;
    // this.filterSearch = new FilterSearch(Number("", "", 0, 10000000, this.pageSize, 0, "price asc")

    // this.getProducts(this.filterSearch)
  }
  public isAdmin() {
    return this.authService.getUserValue() && this.authService.getUserValue().role === Role.Admin;
  }

  getProducts(filterSearch : FilterSearch) {
      this.prodService.getProducts(filterSearch).subscribe(
        data => {
          this.products = []
          console.log("Pre: " + this.products.length)
          this.products = data
          console.log("Posle: " + this.products.length)
          this.itemsCount = data[0].count
        })
  }

  isMine() {
    
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
      this.getProducts(this.filterSearch)
      console.log(event.pageSize)
  }
  edit(product){
    console.log(product.currency)
    this.router.navigate(['/editProduct'], {state: {data: product}});

  }
  remove(product){
    console.log(product)
    // this.router.navigate(['/editProduct'], {state: {data: product}});

    var deletedProduct = new DeletedProduct(product.serial.toString())

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
}
