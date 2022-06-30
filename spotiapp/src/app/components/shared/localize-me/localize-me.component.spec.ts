import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LocalizeMeComponent } from './localize-me.component';

describe('LocalizeMeComponent', () => {
    let component: LocalizeMeComponent;
    let fixture: ComponentFixture<LocalizeMeComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LocalizeMeComponent],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LocalizeMeComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
