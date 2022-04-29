import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalTranslationComponent } from './modal-translation.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { provideMockStore } from '@ngrx/store/testing';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { createMockBaseTranslation } from '../../../types/translation';
import { StageFinderComponent } from '../../stages/stage-finder/stage-finder.component';
import { LanguageFinderComponent } from '../../languages/language-finder/language-finder.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('ModalTranslationComponent', () => {
    let component: ModalTranslationComponent;
    let fixture: ComponentFixture<ModalTranslationComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ModalTranslationComponent, StageFinderComponent, LanguageFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: createMockBaseTranslation(),
                },
                provideMockStore({ initialState: createMockAppState() }),
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ModalTranslationComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
