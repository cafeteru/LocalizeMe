import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { CoreModule } from '../../../core/core.module';
import { SharedModule } from '../../../shared/shared.module';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';
import { DialogMock, DialogRefMock } from '../../../core/mocks/mock-dialog';
import { UserServiceMock } from '../../../core/mocks/services/user.service.mock';
import { UserService } from '../../../core/services/user.service';
import { of, throwError } from 'rxjs';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { createMockUser, User } from '../../../types/user';

describe('LoginComponent', () => {
    let component: LoginComponent;
    let fixture: ComponentFixture<LoginComponent>;
    const dialogRefMock = new DialogRefMock();
    const dialogMock = new DialogMock();
    const userServiceMock = new UserServiceMock();

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LoginComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule, BrowserAnimationsModule],
            providers: [
                provideMockStore({initialState: createMockAppState()}),
                {
                    provide: MatDialogRef,
                    useValue: dialogRefMock,
                },
                {
                    provide: MatDialog,
                    useValue: dialogMock,
                },
                {provide: UserService, useValue: userServiceMock},
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LoginComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    it('check close', () => {
        spyOn(dialogRefMock, 'close');
        component.close();
        expect(dialogRefMock.close).toHaveBeenCalled();
    });

    it('check openRegister', () => {
        spyOn(dialogRefMock, 'close');
        spyOn(dialogMock, 'open');
        component.openRegister();
        expect(dialogRefMock.close).toHaveBeenCalled();
        expect(dialogMock.open).toHaveBeenCalled();
    });

    it('check login with error', () => {
        const spy = spyOn(dialogRefMock, 'close');
        const spyService = spyOn(userServiceMock, 'login').and.returnValue(
          throwError(() => new Error(`Error`))
        );
        const spyLogout = spyOn(userServiceMock, 'logout');
        component.login();
        expect(spyService).toHaveBeenCalled();
        expect(spy).not.toHaveBeenCalled();
        expect(spyLogout).not.toHaveBeenCalled();
    });

    it('check login with active user', () => {
        const spy = spyOn(dialogRefMock, 'close');
        const spyService = spyOn(userServiceMock, 'login').and.returnValue(
          of(createMockUser())
        );
        const spyLogout = spyOn(userServiceMock, 'logout');
        component.login();
        expect(spyService).toHaveBeenCalled();
        expect(spy).toHaveBeenCalled();
        expect(spyLogout).not.toHaveBeenCalled();
    });

    it('check login with no active user', () => {
        const spy = spyOn(dialogRefMock, 'close');
        const user: User = {...createMockUser(), active: false};
        const spyService = spyOn(userServiceMock, 'login').and.returnValue(
          of(user)
        );
        const spyLogout = spyOn(userServiceMock, 'logout');
        component.login();
        expect(spyService).toHaveBeenCalled();
        expect(spy).toHaveBeenCalled();
        expect(spyLogout).toHaveBeenCalled();
    });
});
