import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterComponent } from './register.component';
import { MatDialogRef } from '@angular/material/dialog';
import { SharedModule } from '../../../shared/shared.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CoreModule } from '../../../core/core.module';
import { provideMockStore } from '@ngrx/store/testing';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';

describe('RegisterComponent', () => {
    let component: RegisterComponent;
    let fixture: ComponentFixture<RegisterComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [RegisterComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [
                provideMockStore({ initialState: createMockAppState() }),
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(RegisterComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
