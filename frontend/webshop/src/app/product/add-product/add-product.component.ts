import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { Base64 } from 'js-base64';

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
  submitedPictures: String[];

  foods: Food[] = [
    {value: 'steak-0', viewValue: 'Steak'},
    {value: 'pizza-1', viewValue: 'Pizza'},
    {value: 'tacos-2', viewValue: 'Tacos'}
  ];
  currency : Number;
  fileName : String="";
  upload : Boolean = false;
  choose : Boolean = true;

  constructor(private _formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.submitedPictures = new Array();
    this.nameCategoryGroup = this._formBuilder.group({
      productName: ['', Validators.required],
      productCategory: ['', Validators.required]

    });

    this.descriptionPriceGroup = this._formBuilder.group({
      description: ['', Validators.required],
      price: ['', Validators.required],
      currency:['1', Validators.required],
    });

    this.picturesGroup = this._formBuilder.group({
      pictures: ['', Validators.required],
    });

  }
  onFileChanged(event) {
    const file = event.target.files[0]
    this.fileName = file.name;

    let reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = function () {
      this.submitedPictures.append(reader.result)
      console.log(this.submitedPictures);
    };
    reader.onerror = function (error) {
      console.log('Error: ', error);
    };
  }


}
