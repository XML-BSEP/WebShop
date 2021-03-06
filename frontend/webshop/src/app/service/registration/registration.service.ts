import { environment } from './../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegisteredUser } from 'src/app/model/registeredUser';
import { ConfirmRegistration} from 'src/app/model/confirmRegistration'

@Injectable({
  providedIn: 'root'
})
export class RegistrationService {

  constructor(private https : HttpClient) { }

  register(data : RegisteredUser){
    return this.https.post(`${environment.baseUrl}/${environment.registration}`,data, {responseType : 'json'});
  }

  confAcc(data : ConfirmRegistration){
    return this.https.post(`${environment.baseUrl}/${environment.confirmAccount}`,data, {responseType : 'json'});
  }

  resend(data : Object) {
    return this.https.post(`${environment.baseUrl}/${environment.resendRegCode}`, data, {responseType : 'json' })
  }
}
