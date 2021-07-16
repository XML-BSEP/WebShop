import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeCampaignComponent } from './change-campaign.component';

describe('ChangeCampaignComponent', () => {
  let component: ChangeCampaignComponent;
  let fixture: ComponentFixture<ChangeCampaignComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangeCampaignComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangeCampaignComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
