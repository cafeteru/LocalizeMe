import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateUserComponent } from './update-user.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { createAppStateMock, matDialogRefMock } from '../../../core/mocks/mock-tests';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CoreModule } from '../../../core/core.module';
import { SharedModule } from '../../../shared/shared.module';
import { provideMockStore } from '@ngrx/store/testing';

describe('UpdateUserComponent', () => {
    let component: UpdateUserComponent;
    let fixture: ComponentFixture<UpdateUserComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [UpdateUserComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [
                provideMockStore({ initialState: createAppStateMock() }),
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: { isAdmin: false },
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UpdateUserComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
