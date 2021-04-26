import { AddProductComponent } from './product/add-product/add-product.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { ProductsPageComponent } from './products-page/products-page/products-page.component';
import { RegisterComponent } from './register/register.component';
import { RegistrationConfirmationComponent } from './registration-confirmation/registration-confirmation.component';
import {ForgotPasswordComponent } from './forgot-password/forgot-password.component'
import { ForbiddenComponent } from './other/forbidden/forbidden.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
},
{
  path: 'home',
  component: HomeComponent
},
{
  path: 'login',
  component: LoginComponent
},
{
  path:'registration',
  component: RegisterComponent
},
{
  path:'regconfirm',
  component: RegistrationConfirmationComponent
},
{
  path: 'products',
  component: ProductsPageComponent
}, 
{
  path:'addProduct',
  component: AddProductComponent
},
{
  path:'forgotPassword',
  component : ForgotPasswordComponent
},
{
  path: 'forbidden',
  component: ForbiddenComponent
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
