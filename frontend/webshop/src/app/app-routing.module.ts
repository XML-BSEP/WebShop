import { ShoppingCartComponent } from './shopping-cart/shopping-cart.component';
import { EditProductComponent } from './product/edit-product/edit-product.component';
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
import { AuthGuard } from './helpers';
import { Role } from './model/role';
import { ShopHomeComponent } from './shop-home/shop-home.component';
import { CreateCampaingComponent } from './create-campaing/create-campaing.component';
import { CreateAd } from './model/create_ad';
import { ChangeCampaignComponent } from './change-campaign/change-campaign.component';
import { CreateAdComponent } from './create-ad/create-ad.component';
import { EditProfileComponent } from './userprofile/edit-profile/edit-profile.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
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
  component: RegistrationConfirmationComponent,

},
{
  path: 'products',
  component: ProductsPageComponent
},
{
  path: 'shopHome',
  component: ShopHomeComponent
},
{
  path:'addProduct',
  component: AddProductComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
},
{
  path:'editProduct',
  component: EditProductComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
},
{
  path:'cart',
  component: ShoppingCartComponent,
  canActivate:[AuthGuard],
  data:{role:[Role.Customer]}
},
{
  path:'forgotPassword',
  component : ForgotPasswordComponent
},
{
  path: 'forbidden',
  component: ForbiddenComponent
},

{
  path: 'createAd',
  component : CreateAdComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
},
{
  path: 'createCampaign',
  component : CreateCampaingComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
}, 
{
  path: 'changeCampaign',
  component : ChangeCampaignComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
},
{
  path: 'settings',
  component : EditProfileComponent,
  canActivate : [AuthGuard],
  data : {roles: [Role.Admin]}
}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
