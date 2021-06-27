import { ProductForCart } from './../../model/productForCart';
import { UserDTO } from './../../model/userDTO';
import { ItemToCart } from './../../model/itemToCart';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Shop } from 'src/app/model/shop';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ShopService {

  constructor(private https : HttpClient) { }

  getAllShops() : Observable<Shop[]> {
    return this.https.get<Shop[]>(`${environment.baseUrl}/${environment.allShops}`)
  }

  addToCart(product : ItemToCart){
    return this.https.post(`${environment.baseUrl}/${environment.addToCart}`,product, {responseType : 'json'});
  }
  getCart(user : UserDTO): Observable<ProductForCart[]>{
    return this.https.post<ProductForCart[]>(`${environment.baseUrl}/${environment.getCart}`,user);
  }
}
