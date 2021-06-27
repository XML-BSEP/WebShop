import { Product } from 'src/app/model/product';
import { CategoryService } from './../../service/product/category.service';
import { ProductServiceService } from './../../service/product/product-service.service';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { Category } from './../../model/category';
import { Image } from './../../model/image';
import { NewProduct } from './../../model/newProduct';
import { FormGroup, FormBuilder, Validators, FormControl } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-edit-product',
  templateUrl: './edit-product.component.html',
  styleUrls: ['./edit-product.component.css']
})
export class EditProductComponent implements OnInit {
  nameCategoryGroup: FormGroup;
  descriptionPriceGroup: FormGroup;
  picturesGroup : FormGroup;
  newProduct : NewProduct;
  images: Image[] = new Array();
  isLinear : Boolean = true;
  submitedPictures: String[] = [];


  fileName : String="";
  imgFile : string;
  upload : Boolean = false;
  choose : Boolean = true;
  current=0;
  cat : Number;
  numberRegEx = /\-?\d*\.?\d{1,2}/;

  allCategories : Category[];
  productToEdit : Product;
  files : String[] = new Array();
  constructor(private router: Router, private toastr : ToastrService, private productService : ProductServiceService, private categoryService : CategoryService ,private _formBuilder: FormBuilder) { }

  ngOnInit(): void {
    if(history.state.data === undefined){
      this.router.navigate(['/home'])
    }else{
      this.productToEdit = history.state.data;
      console.log(this.productToEdit)

    }
    this.nameCategoryGroup = this._formBuilder.group({
      'productName' : new FormControl(this.productToEdit.name, Validators.required),
      'productCategory' : new FormControl(this.productToEdit.category, Validators.required),
    });

    this.descriptionPriceGroup = this._formBuilder.group({
      'description' : new FormControl(this.productToEdit.description, Validators.required),
      'price' : new FormControl(this.productToEdit.price,[ Validators.required,  Validators.pattern(this.numberRegEx)]),
      'available' : new FormControl(this.productToEdit.available,[ Validators.required,  Validators.pattern('^[0-9]+$')]),
    });
    this.selectImages();
    this.getCategories();
  }

  selectImages(){
    for(let i=0; i<this.productToEdit.image.length; i++){
      let slices = this.productToEdit.image[i].split(':')
      let second = slices[2].split('/')
      let name = second[3]
      this.images.push(new Image(name,this.productToEdit.image[i]))

      this.current++;
    }
  }


toDataUrl(url, callback) {
  var xhr = new XMLHttpRequest();
  xhr.onload = function() {
      var reader = new FileReader();
      reader.onloadend = function() {
          callback(reader.result);
          return reader.result as string;
      }
      reader.readAsDataURL(xhr.response);
  };
  xhr.open('GET', url);
  xhr.responseType = 'blob';
  xhr.send();
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
            console.log(this.images);
          };
      }

    }

  }

  getCategories() {
    this.categoryService.getAllCategories().subscribe(data => {
      this.allCategories = data;
    })
}

  removeImg(img){
    this.images = this.images.filter(obj => obj !== img);
    this.current--;
  }

  confirmEdit(){
    var blobs : String[] = new Array();
    for(let i=0; i<this.images.length;i++){
      blobs.push(this.images[i].file);
    }
    console.log(blobs)

    this.newProduct = new NewProduct(this.nameCategoryGroup.controls.productName.value,
                                    this.nameCategoryGroup.controls.productCategory.value,
                                    this.descriptionPriceGroup.controls.price.value.toString(),
                                    this.descriptionPriceGroup.controls.description.value,
                                    blobs,
                                    this.descriptionPriceGroup.controls.available.value.toString(),
                                    this.productToEdit.serial.toString());

    console.log(this.newProduct);

    this.productService.editProduct(this.newProduct).subscribe(
      res=>{
        this.toastr.success('Success');
        this.router.navigate(['/products'])
      },
      err=>{
        this.toastr.error(err)
      }
        )

  }
}
