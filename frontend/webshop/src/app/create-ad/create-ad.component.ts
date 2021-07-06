import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { map, startWith } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { CreateAd } from 'src/app/model/create_ad';
import { AgentService } from 'src/app/service/agent_service';

@Component({
  selector: 'app-create-ad',
  templateUrl: './create-ad.component.html',
  styleUrls: ['./create-ad.component.css']
})
export class CreateAdComponent implements OnInit {

  newAdForm : FormGroup
  imgFile : string;
  isImage: boolean = true;
  isVideo : boolean = true;
  isVideoSelected : boolean = false;
  current : number;
  video : string
  isSelected : boolean = false;
  myControl = new FormControl();
  tagControl = new FormControl();
  hashtags : string[] = new Array();
  hash : string ='';
  linkShop : string = '';
  sinImputarValue : string = 'postChecked';

  ngOnInit(): void {
    this.current=0;
    this.newAdForm = new FormGroup({
      'caption' : new FormControl(null, [Validators.required]),
      'location' : new FormControl(null, [Validators.required]),
      'hash' : new FormControl(null, []),
      'link' : new FormControl(null, [Validators.required]),
     
    });

  }

  constructor(private toastr : ToastrService, private agentService : AgentService) { 
   
  }


  checkradio() {
    console.log(this.newAdForm.controls.postStoryChecked.value)
  }
  checkIfHashtagIsTagged(tag){
    return this.hashtags.some(element => element === tag)
  }
  createAd() {
    
    console.log(this.video)
    var image = ""
    var album : string[] = new Array()
    var video = ""
    if(this.isVideo){
      video = this.video
    }
    
    if(this.isVideo){
      video = this.video
    }
    
    if(this.isImage){
      image = this.imgFile
    }
    var newAd = new CreateAd();

    newAd.description = this.newAdForm.controls.caption.value;
    newAd.hashtags = this.hashtags;
    newAd.link = this.newAdForm.controls.link.value;
    newAd.location = this.newAdForm.controls.location.value;
    if (this.isImage) {
      newAd.media = image;
      newAd.type = 0;
    } else {
      newAd.media = video;
      newAd.type = 1;
    }
    
    

    this.agentService.createAd(newAd).subscribe(
      res => {
        this.toastr.success("Successfully added media!")
      }, error => {
        this.toastr.error(error)
      }      
    )
  }
  hashtag(){
    if(!this.checkIfHashtagIsTagged(this.newAdForm.controls.hash.value)){
      this.hashtags.push(this.newAdForm.controls.hash.value)
      this.newAdForm.setValue({
        hash: ""
      });
    }
  }

  onFileChangedVideo(e){
    this.isImage=false;
    this.isVideo = true;


    const reader = new FileReader();
      if(e.target.files && e.target.files.length) {
        const [file] = e.target.files;
          reader.readAsDataURL(file);

          reader.onload = () => {
            this.video = reader.result as string;
            this.isSelected=true

            this.isVideoSelected = true;
          };


    }

  }

  onFileChanged(e) {
    this.isImage=true;
    this.isVideo = false;
    const reader = new FileReader();
      if(e.target.files && e.target.files.length) {
        const [file] = e.target.files;
          reader.readAsDataURL(file);

          reader.onload = () => {
            this.imgFile = reader.result as string;
            this.isSelected=true

          };


    }

  }

}
