import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-show-image',
  templateUrl: './show-image.component.html',
  styleUrls: ['./show-image.component.css']
})
export class ShowImageComponent implements OnInit {
  public image;
  constructor( public dialogRef: MatDialogRef<ShowImageComponent>, private dialog : MatDialog,
    @Inject(MAT_DIALOG_DATA) public data: String) {}
  ngOnInit(): void {
    this.image = this.data;
  }

}
