import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TranslationListComponent } from './translation-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('TranslationListComponent', () => {
    let component: TranslationListComponent;
    let fixture: ComponentFixture<TranslationListComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [TranslationListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(TranslationListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
