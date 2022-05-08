import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BaseStringTranslationTableComponent } from './base-string-translation-table.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';

describe('BaseStringTranslationTableComponent', () => {
    let component: BaseStringTranslationTableComponent;
    let fixture: ComponentFixture<BaseStringTranslationTableComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [BaseStringTranslationTableComponent],
            imports: [SharedModule, CoreModule],
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
