import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Router } from '@angular/router';
import { AuthenticationService } from '../service/authentication/authentication.service';

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
    constructor(private router : Router,
       private authenticationService: AuthenticationService
       ) { }

    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        return next.handle(request).pipe(catchError(err => {
            if ([401, 403].indexOf(err.status) !== -1) {
                // auto logout if 401 response returned from api
                // this.authenticationService.logout();
                let currentUser = this.authenticationService.currentUserValue;

                if (currentUser && currentUser.refresh_token) {
                    this.authenticationService.refresh(currentUser.refresh_token).subscribe(result => {
                        localStorage.setItem('userId',String(result.id))
                        this.router.navigate(['/'])
                    },
                    error=>{
                        this.router.navigate(['/login']);
                    });

                } else {
                    const error = err.error.message || err.statusText;
                    return throwError(error);
                }
            }
        }))
    }
}
