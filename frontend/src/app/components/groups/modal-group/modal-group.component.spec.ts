import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalGroupComponent } from './modal-group.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { createAppStateMock } from '../../../store/mocks/create-app-state-mock';
import { createMockGroup } from '../../../types/group';

describe('ModalGroupComponent', () => {
    let component: ModalGroupComponent;
    let fixture: ComponentFixture<ModalGroupComponent>;
    let store: MockStore;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ModalGroupComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: createMockGroup(),
                },
                provideMockStore({ initialState: createAppStateMock() }),
            ],
        }).compileComponents();
        store = TestBed.inject(MockStore);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ModalGroupComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
