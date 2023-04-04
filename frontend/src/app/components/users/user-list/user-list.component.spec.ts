import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AppState } from '../../../store/app.reducer';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';
import { createMockUser, User } from '../../../types/user';
import { of, throwError } from 'rxjs';
import { UserService } from '../../../core/services/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { mockNzMessageService } from '../../../core/mocks/mock-nz-message-service';
import { mockNzModalService } from '../../../core/mocks/mock-nz-modal-service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MatDialogRef } from '@angular/material/dialog';
import { MockUserService } from '../../../core/mocks/services/mockUserService';

describe('UserListComponent', () => {
    let component: UserListComponent;
    let fixture: ComponentFixture<UserListComponent>;
    let appState: AppState;
    let store: MockStore;
    const mockUserService = new MockUserService();
    let user: User;

    beforeEach(async () => {
        appState = createMockAppState();
        await TestBed.configureTestingModule({
            declarations: [UserListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [
                { provide: UserService, useValue: mockUserService },
                { provide: NzMessageService, useValue: mockNzMessageService },
                { provide: NzModalService, useValue: mockNzModalService },
                provideMockStore({ initialState: appState }),
            ],
        }).compileComponents();
        store = TestBed.inject(MockStore);
        fixture = TestBed.createComponent(UserListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    beforeEach(() => {
        user = createMockUser();
    });

    afterEach(() => {
        store.resetSelectors();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    describe('loadUsers', () => {
        it('check with error', () => {
            const spy = spyOn(mockUserService, 'findAll').and.returnValue(throwError(() => new Error('error')));
            component.loadUsers();
            expect(spy).toHaveBeenCalled();
            expect(component.users).toEqual([]);
        });

        it('check without errors', () => {
            const spy = spyOn(mockUserService, 'findAll').and.returnValue(of([createMockUser()]));
            component.loadUsers();
            expect(spy).toHaveBeenCalled();
            expect(component.users).not.toEqual([]);
        });
    });

    describe('disable', () => {
        it('check without errors', () => {
            const spy = spyOn(mockUserService, 'disable').and.returnValue(of(user));
            const spyLoadUsers = spyOn(component, 'loadUsers');
            component.disable(user);
            expect(spy).toHaveBeenCalled();
            expect(spyLoadUsers).toHaveBeenCalled();
        });

        it('check with error', () => {
            const spy = spyOn(mockUserService, 'disable').and.returnValue(throwError(() => new Error('error')));
            const spyMessageService = spyOn(mockNzMessageService, 'create');
            component.disable(user);
            expect(spy).toHaveBeenCalled();
            expect(spyMessageService).toHaveBeenCalled();
        });
    });

    describe('callToDelete', () => {
        it('should call to delete', () => {
            const spy = spyOn(component as any, 'delete');
            const callToDelete = component['callToDelete'](user);
            callToDelete();
            expect(spy).toHaveBeenCalled();
        });
    });

    describe('delete', () => {
        it('check correct', () => {
            const spy = spyOn(mockUserService, 'delete').and.returnValue(of(true));
            const spy2 = spyOn(component, 'loadUsers');
            const spy3 = spyOn(mockNzMessageService, 'create');
            component['delete'](user);
            expect(spy).toHaveBeenCalled();
            expect(spy2).toHaveBeenCalled();
            expect(spy3).toHaveBeenCalled();
        });

        it('check with not found user', () => {
            const spy = spyOn(mockUserService, 'delete').and.returnValue(of(false));
            const spy2 = spyOn(component, 'loadUsers');
            const spy3 = spyOn(mockNzMessageService, 'create');
            component['delete'](user);
            expect(spy).toHaveBeenCalled();
            expect(spy2).not.toHaveBeenCalled();
            expect(spy3).toHaveBeenCalled();
        });

        it('check with server error', () => {
            const spy = spyOn(mockUserService, 'delete').and.returnValue(throwError(() => new Error('error')));
            const spy2 = spyOn(component, 'loadUsers');
            const spy3 = spyOn(mockNzMessageService, 'create');
            component['delete'](user);
            expect(spy).toHaveBeenCalled();
            expect(spy2).not.toHaveBeenCalled();
            expect(spy3).toHaveBeenCalled();
        });
    });

    describe('showDeleteModal', () => {
        it('check it', () => {
            const spy = spyOn(mockNzModalService, 'confirm');
            component.showDeleteModal(user);
            expect(spy).toHaveBeenCalled();
        });
    });

    describe('openCreate', () => {
        it('check with create', () => {
            const modalSpy = spyOn(component.dialog, 'open').and.returnValue({
                afterClosed: () => of(undefined),
            } as MatDialogRef<typeof component>);
            const spy = spyOn(component, 'loadUsers');
            component.openUpdate(user);
            expect(spy).not.toHaveBeenCalled();
            expect(modalSpy).toHaveBeenCalled();
        });

        it('check with update', () => {
            const modalSpy = spyOn(component.dialog, 'open').and.returnValue({
                afterClosed: () => of(user),
            } as MatDialogRef<typeof component>);
            const spy2 = spyOn(component, 'loadUsers');
            component.openUpdate(user);
            expect(spy2).toHaveBeenCalled();
            expect(modalSpy).toHaveBeenCalled();
        });
    });
});
