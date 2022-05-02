import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BaseStringTranslationTableComponent } from './base-string-translation-table.component';

describe('GbaseStringTranslationTableComponent', () => {
    let component: BaseStringTranslationTableComponent;
    let fixture: ComponentFixture<BaseStringTranslationTableComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [BaseStringTranslationTableComponent],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(BaseStringTranslationTableComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
