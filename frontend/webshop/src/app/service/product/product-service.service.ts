import { NewProduct } from './../../model/newProduct';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Product } from 'src/app/model/product';
import { FilterSearch } from 'src/app/model/filterSearch';

@Injectable({
  providedIn: 'root'
})
export class ProductServiceService {

  constructor(private http: HttpClient) { }

  getProducts(filterSearch : FilterSearch) : Observable<Product[]> {
    return this.http.post<Product[]>(`${environment.baseUrl}/${environment.filterSearch}`, filterSearch)
  }

  addProduct(data : NewProduct){
    return this.http.post(`${environment.baseUrl}/${environment.addProduct}`,data, {responseType : 'text'});
  }

}
