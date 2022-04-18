import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalLanguageComponent } from './modal-language.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { createMockLanguage } from '../../../types/language';

describe('ModalLanguageComponent', () => {
    let component: ModalLanguageComponent;
    let fixture: ComponentFixture<ModalLanguageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ModalLanguageComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: createMockLanguage(),
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ModalLanguageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
