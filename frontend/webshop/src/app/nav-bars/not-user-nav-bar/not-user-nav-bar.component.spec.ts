import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotUserNavBarComponent } from './not-user-nav-bar.component';

describe('NotUserNavBarComponent', () => {
  let component: NotUserNavBarComponent;
  let fixture: ComponentFixture<NotUserNavBarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NotUserNavBarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NotUserNavBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
