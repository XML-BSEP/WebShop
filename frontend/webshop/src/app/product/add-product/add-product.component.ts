import { ProductServiceService } from 'src/app/service/product/product-service.service';
import { Category } from './../../model/category';
import { NewProduct } from './../../model/newProduct';
import { Image } from './../../model/image';
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { Base64 } from 'js-base64';
import { CategoryService } from 'src/app/service/product/category.service';
import { ToastrService } from 'ngx-toastr';

interface Food {
  value: string;
  viewValue: string;
}
@Component({
  selector: 'app-add-product',
  templateUrl: './add-product.component.html',
  styleUrls: ['./add-product.component.css']
})
export class AddProductComponent implements OnInit {
  nameCategoryGroup: FormGroup;
  descriptionPriceGroup: FormGroup;
  picturesGroup : FormGroup;
  newProduct : NewProduct;
  images: Image[] = new Array();
  isLinear : Boolean = true;
  submitedPictures: String[] = [];

  foods: Food[] = [
    {value: 'steak-0', viewValue: 'Steak'},
    {value: 'pizza-1', viewValue: 'Pizza'},
    {value: 'tacos-2', viewValue: 'Tacos'}
  ];
  currency : Number;
  fileName : String="";
  imgFile : string;
  upload : Boolean = false;
  choose : Boolean = true;
  current=0;
  numberRegEx = /\-?\d*\.?\d{1,2}/;
  allCategories : Category[];
  constructor(private toastr : ToastrService, private productService : ProductServiceService, private categoryService : CategoryService ,private _formBuilder: FormBuilder) { }

  ngOnInit(): void {

    this.nameCategoryGroup = this._formBuilder.group({
      productName: ['', Validators.required],
      productCategory: ['', Validators.required],
    });

    this.descriptionPriceGroup = this._formBuilder.group({
      description: ['', Validators.required],
      'price' : new FormControl(null,[ Validators.required,  Validators.pattern(this.numberRegEx)]),
      'available' : new FormControl(null,[ Validators.required,  Validators.pattern(this.numberRegEx)]),
      currency:['1', Validators.required]
    });

    this.getCategories();
  }
  onFileChanged(e) {
    const reader = new FileReader();
    if(this.current<4){
      if(e.target.files && e.target.files.length) {
        const [file] = e.target.files;
          reader.readAsDataURL(file);

          reader.onload = () => {
            this.imgFile = reader.result as string;
            this.fileName = file.name;
            this.images.push(new Image(this.fileName, this.imgFile));

            this.current++;
            console.log(this.current);
          };
      }

    }

  }

  getCategories() {
    this.categoryService.getAllCategories().subscribe(data => {this.allCategories = data})
}

  removeImg(img){
    this.images = this.images.filter(obj => obj !== img);
    this.current--;
  }

  addProduct(){
    var blobs : String[] = new Array();
    for(let i=0; i<this.images.length;i++){
      blobs.push(this.images[i].file);
    }
    console.log(this.descriptionPriceGroup.controls.currency.value)

    this.newProduct = new NewProduct(this.nameCategoryGroup.controls.productName.value,
                                    this.nameCategoryGroup.controls.productCategory.value,
                                    this.descriptionPriceGroup.controls.price.value,
                                    this.descriptionPriceGroup.controls.description.value,
                                    blobs,
                                    this.descriptionPriceGroup.controls.currency.value,
                                    this.descriptionPriceGroup.controls.available.value);

    console.log(this.newProduct);

    this.productService.addProduct(this.newProduct).subscribe(
      res=>{
        this.toastr.success('Success');
      },
      err=>{
        this.toastr.error(err)
      }
        )

  }
}
