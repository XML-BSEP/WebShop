import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { AuthenticatedUser } from 'src/app/model/authenticatedUser';
import { Authentication } from '../../model/authentication';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private currentUserSubject: BehaviorSubject<AuthenticatedUser>;
  public currentUser: Observable<AuthenticatedUser>;

  constructor(private http: HttpClient, private router : Router) {
      this.currentUserSubject = new BehaviorSubject<AuthenticatedUser>(JSON.parse(localStorage.getItem('currentUser')));
      this.currentUser = this.currentUserSubject.asObservable();
  }

  public get currentUserValue(): AuthenticatedUser {
      return this.currentUserSubject.value;
  }

public getUserValue() : AuthenticatedUser {
   // console.log("Token" + this.currentUserSubject.value.token.accessToken);
    return this.currentUserSubject.value;
}

login(credentials: Authentication){


    return this.http.post<AuthenticatedUser>(`${environment.baseUrl}/${environment.login}`, credentials)
    .pipe(map(response => {
      localStorage.setItem('currentUser', JSON.stringify(response));
      console.log(response);
      this.currentUserSubject.next(response);
      return response;
    }));
}

logout() {
      // remove user from local storage to log user out

      this.http.post(`${environment.baseUrl}/${environment.login}`, null);
      localStorage.removeItem('currentUser');
      localStorage.removeItem('userId');
      localStorage.removeItem('role');
      this.currentUserSubject.next(null);
      this.router.navigate(['/']);
}

refresh(refreshToken : string) {
  return this.http.post<AuthenticatedUser>(`${environment.baseUrl}/${environment.refresh}`, refreshToken)
  .pipe(map(response => {
    localStorage.setItem('currentUser', JSON.stringify(response));
    this.currentUserSubject.next(response);
    return response;
  }))
}
}
