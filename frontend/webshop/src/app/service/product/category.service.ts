import { Observable } from 'rxjs';
import { environment } from './../../../environments/environment';
import { Category } from './../../model/category';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor(private http: HttpClient) { }

  getAllCategories() : Observable<Category[]> {
    return this.http.get<Category[]>(`${environment.baseUrl}/${environment.categories}`)
  }
}
