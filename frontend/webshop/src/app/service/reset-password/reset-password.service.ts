import { Injectable } from '@angular/core';
import { environment } from './../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { ResetMail } from 'src/app/model/resetMail';
import { ResetPass } from 'src/app/model/resetPass';

@Injectable({
  providedIn: 'root'
})
export class ResetPasswordService {

  constructor(private https : HttpClient) { }

  resetPasswordMail(data : ResetMail){
    return this.https.post(`${environment.baseUrl}/${environment.resetPasswordMail}`,data, {responseType : 'text'});
  }

  
  resetPassword(data : ResetPass){
    return this.https.post(`${environment.baseUrl}/${environment.resetPassword}`,data, {responseType : 'text'});
  }
}
