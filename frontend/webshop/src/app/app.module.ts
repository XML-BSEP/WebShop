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
import { ProductsPageComponent } from './products-page/products-page/products-page.component';
import { MatInputModule } from '@angular/material/input';
import {MatSelectModule} from '@angular/material/select';
import { NotUserNavBarComponent } from './nav-bars/not-user-nav-bar/not-user-nav-bar.component';
import { CustomerNavBarComponent } from './nav-bars/customer-nav-bar/customer-nav-bar.component';
import { AddProductComponent } from './product/add-product/add-product.component';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { ForbiddenComponent } from './other/forbidden/forbidden.component';
import { ToastrModule } from 'ngx-toastr';
import { EditProductComponent } from './product/edit-product/edit-product.component';
import { ShoppingCartComponent } from './shopping-cart/shopping-cart.component';
import { ShopHomeComponent } from './shop-home/shop-home.component';
import { MatCarouselModule } from '@ngmodule/material-carousel';
import { CreateAdComponent } from './create-ad/create-ad.component';
import { CreateCampaingComponent } from './create-campaing/create-campaing.component';
import { ChangeCampaignComponent } from './change-campaign/change-campaign.component';
import { ShowImageComponent } from './dialogs/show-image/show-image.component';
import { EditProfileComponent } from './userprofile/edit-profile/edit-profile.component';
import { DatePipe } from '@angular/common'


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    RegisterComponent,
    RegistrationConfirmationComponent,
    AdminNavBarComponent,
    ProductsPageComponent,
    NotUserNavBarComponent,
    CustomerNavBarComponent,
    AddProductComponent,
    ForgotPasswordComponent,
    ForbiddenComponent,
    EditProductComponent,
    ShoppingCartComponent,
    ShopHomeComponent,
    CreateAdComponent,
    CreateCampaingComponent,
    ChangeCampaignComponent,
    ShowImageComponent,
    EditProfileComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatInputModule,
    MatSelectModule,
    ToastrModule.forRoot(),
    MatCarouselModule.forRoot()
  ],
  providers: [
    DatePipe,
    {provide : HTTP_INTERCEPTORS, useClass : JwtInterceptor, multi : true},
    {provide : HTTP_INTERCEPTORS, useClass : ErrorInterceptor, multi : true}],
  
  bootstrap: [AppComponent]
})
export class AppModule { }
