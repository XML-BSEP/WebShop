import { Component, OnInit } from '@angular/core';
import { Product } from 'src/app/model/product';
import { ProductServiceService } from 'src/app/service/product/product-service.service';

@Component({
  selector: 'app-products-page',
  templateUrl: './products-page.component.html',
  styleUrls: ['./products-page.component.css']
})
export class ProductsPageComponent implements OnInit {

  products : Product[]
  constructor(private prodService : ProductServiceService) { }

  ngOnInit(): void {
    this.getProducts()
  }

  getProducts() {
      this.prodService.getProducts().subscribe(data => {this.products = data})
  }

}
