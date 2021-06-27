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

}
