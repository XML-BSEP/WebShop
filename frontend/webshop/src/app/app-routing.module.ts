import { AddProductComponent } from './product/add-product/add-product.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { RegistrationConfirmationComponent } from './registration-confirmation/registration-confirmation.component';

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
  path:'addProduct',
  component: AddProductComponent
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
