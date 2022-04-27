import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalBaseStringComponent } from './modal-base-string.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { provideMockStore } from '@ngrx/store/testing';
import { createAppStateMock } from '../../../store/mocks/create-app-state-mock';
import { createMockBaseString } from '../../../types/base-string';
import { GroupFinderComponent } from '../../groups/group-finder/group-finder.component';
import { LanguageFinderComponent } from '../../languages/language-finder/language-finder.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('ModalBaseStringComponent', () => {
    let component: ModalBaseStringComponent;
    let fixture: ComponentFixture<ModalBaseStringComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ModalBaseStringComponent, GroupFinderComponent, LanguageFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: createMockBaseString(),
                },
                provideMockStore({ initialState: createAppStateMock() }),
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ModalBaseStringComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
