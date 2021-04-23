import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material-module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ErrorInterceptor, JwtInterceptor } from './helpers';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { RegisterComponent } from './register/register.component';
import { RegistrationConfirmationComponent } from './registration-confirmation/registration-confirmation.component';
import { AdminNavBarComponent } from './nav-bars/admin-nav-bar/admin-nav-bar.component';
import { NotUserNavBarComponent } from './nav-bars/not-user-nav-bar/not-user-nav-bar.component';
import { CustomerNavBarComponent } from './nav-bars/customer-nav-bar/customer-nav-bar.component';
import { AddProductComponent } from './product/add-product/add-product.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    RegisterComponent,
    RegistrationConfirmationComponent,
    AdminNavBarComponent,
    NotUserNavBarComponent,
    CustomerNavBarComponent,
    AddProductComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [{provide : HTTP_INTERCEPTORS, useClass : JwtInterceptor, multi : true},
    {provide : HTTP_INTERCEPTORS, useClass : ErrorInterceptor, multi : true}],
  bootstrap: [AppComponent]
})
export class AppModule { }
