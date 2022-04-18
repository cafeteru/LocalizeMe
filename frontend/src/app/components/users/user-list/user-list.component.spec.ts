import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AppState } from '../../../store/app.reducer';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { createAppStateMock } from '../../../store/mocks/create-app-state-mock';
import { createMockUser } from '../../../types/user';
import { of, throwError } from 'rxjs';
import { UserService } from '../../../core/services/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { nzMessageServiceMock } from '../../../core/mocks/nz-message-service-mock';
import { nzModalServiceMock } from '../../../core/mocks/nz-modal-service-mock';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MatDialogRef } from '@angular/material/dialog';
import { UserServiceMock } from '../../../core/mocks/services/user.service.mock';

describe('UserListComponent', () => {
    let component: UserListComponent;
    let fixture: ComponentFixture<UserListComponent>;
    let appState: AppState;
    let store: MockStore;
    const userServiceMock = new UserServiceMock();

    beforeEach(async () => {
        appState = createAppStateMock();
        await TestBed.configureTestingModule({
            declarations: [UserListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [
                { provide: UserService, useValue: userServiceMock },
                { provide: NzMessageService, useValue: nzMessageServiceMock },
                { provide: NzModalService, useValue: nzModalServiceMock },
                provideMockStore({ initialState: appState }),
            ],
        }).compileComponents();
        store = TestBed.inject(MockStore);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UserListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    afterEach(() => {
        store.resetSelectors();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    it('check loadUsers with error', () => {
        const spy = spyOn(userServiceMock, 'findAll').and.returnValue(throwError('error'));
        component.loadUsers();
        expect(spy).toHaveBeenCalled();
        expect(component.users).toEqual([]);
    });

    it('check loadUsers', () => {
        const spy = spyOn(userServiceMock, 'findAll').and.returnValue(of([createMockUser()]));
        component.loadUsers();
        expect(spy).toHaveBeenCalled();
        expect(component.users).not.toEqual([]);
    });

    it('check disable', () => {
        const user = createMockUser();
        const spy = spyOn(userServiceMock, 'disable').and.returnValue(of(user));
        const spy2 = spyOn(component, 'loadUsers');
        component.disable(user);
        expect(spy).toHaveBeenCalled();
        expect(spy2).toHaveBeenCalled();
    });

    it('check disable with error', () => {
        const user = createMockUser();
        const spy = spyOn(userServiceMock, 'disable').and.returnValue(throwError('error'));
        const spy2 = spyOn(nzMessageServiceMock, 'create');
        component.disable(user);
        expect(spy).toHaveBeenCalled();
        expect(spy2).toHaveBeenCalled();
    });

    it('check delete correct', () => {
        const user = createMockUser();
        const spy = spyOn(userServiceMock, 'delete').and.returnValue(of(true));
        const spy2 = spyOn(component, 'loadUsers');
        const spy3 = spyOn(nzMessageServiceMock, 'create');
        component['delete'](user);
        expect(spy).toHaveBeenCalled();
        expect(spy2).toHaveBeenCalled();
        expect(spy3).toHaveBeenCalled();
    });

    it('check delete with not found user', () => {
        const user = createMockUser();
        const spy = spyOn(userServiceMock, 'delete').and.returnValue(of(false));
        const spy2 = spyOn(component, 'loadUsers');
        const spy3 = spyOn(nzMessageServiceMock, 'create');
        component['delete'](user);
        expect(spy).toHaveBeenCalled();
        expect(spy2).not.toHaveBeenCalled();
        expect(spy3).toHaveBeenCalled();
    });

    it('check delete with server error', () => {
        const user = createMockUser();
        const spy = spyOn(userServiceMock, 'delete').and.returnValue(throwError('error'));
        const spy2 = spyOn(component, 'loadUsers');
        const spy3 = spyOn(nzMessageServiceMock, 'create');
        component['delete'](user);
        expect(spy).toHaveBeenCalled();
        expect(spy2).not.toHaveBeenCalled();
        expect(spy3).toHaveBeenCalled();
    });

    it('check showDeleteModal', () => {
        const user = createMockUser();
        const spy = spyOn(nzModalServiceMock, 'confirm');
        component.showDeleteModal(user);
        expect(spy).toHaveBeenCalled();
    });

    it('check openUpdate', () => {
        const user = createMockUser();
        spyOn(component.dialog, 'open').and.returnValue({
            afterClosed: () => of(undefined),
        } as MatDialogRef<typeof component>);
        const spy = spyOn(component, 'loadUsers');
        component.openUpdate(user);
        expect(spy).not.toHaveBeenCalled();
    });

    it('check openUpdate with update', () => {
        const user = createMockUser();
        spyOn(component.dialog, 'open').and.returnValue({
            afterClosed: () => of(user),
        } as MatDialogRef<typeof component>);
        const spy2 = spyOn(component, 'loadUsers');
        component.openUpdate(user);
        expect(spy2).toHaveBeenCalled();
    });
});
