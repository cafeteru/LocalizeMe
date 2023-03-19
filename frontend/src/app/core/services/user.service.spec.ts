import { TestBed } from '@angular/core/testing';

import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { AppState } from '../../store/app.reducer';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';
import { ResponseLogin } from '../../types/response-login';
import { createMockUser, User } from '../../types/user';
import { UserService } from './user.service';

describe('UserService', () => {
    let service: UserService;
    let mockHttp: HttpTestingController;
    let appState: AppState;

    beforeEach(() => {
        appState = createMockAppState();
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
            providers: [provideMockStore({ initialState: appState })],
        });
        service = TestBed.inject(UserService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    it('check login', () => {
        const loginData: User = {
            email: 'username',
            password: 'password',
            id: undefined,
            admin: undefined,
            active: true,
        };
        const token: ResponseLogin = {
            authorization:
                'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmUiOnRydWU' +
                'sImFkbWluIjp0cnVlLCJlbWFpbCI6ImFkbWluQGVtYWlsLmVzIiwiZXhwIjox' +
                'NjUwMzIxMDE5LCJpZCI6IjYyMjEyYjkyYWI2MzE0MWE2ODQ3MzlmMyJ9.hGw' +
                '3Seg9PqVpLauF9XZiC_XhWNwBVWc-jbW5mCzARi4',
        };
        service.login(loginData).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/login`);
        expect(req.request.method).toBe('POST');
        req.flush(token);
    });

    it('check login not active', () => {
        const loginData: User = {
            email: 'username',
            password: 'password',
            id: undefined,
            admin: undefined,
            active: false,
        };
        const token: ResponseLogin = {
            authorization:
                'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmUiOmZhbHNlLC' +
                'JhZG1pbiI6dHJ1ZSwiZW1haWwiOiJhZG1pbkBlbWFpbC5lcyIsImV4cCI6MTY1' +
                'MDMyMDk2NywiaWQiOiI2MjIxMmI5MmFiNjMxNDFhNjg0NzM5ZjMifQ.aqk_ukbf' +
                'GT72o-DmPyxCid6r4zcjVc4HY9uPAnUJ_cY',
        };
        service.login(loginData).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/login`);
        expect(req.request.method).toBe('POST');
        req.flush(token);
    });

    it('check logout', () => {
        service.logout();
        expect(appState.userInfo.user.email).toEqual('');
    });

    it('check findMe', () => {
        service.findMe().subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/me`);
        expect(req.request.method).toBe('GET');
        req.flush(createMockUser());
    });

    it('check register', () => {
        const loginData: User = {
            email: 'username',
            password: 'password',
            id: undefined,
            admin: false,
            active: true,
        };
        service.register(loginData).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}`);
        expect(req.request.method).toBe('POST');
        req.flush(createMockUser());
    });

    it('check updateMe', () => {
        service.updateMe(createMockUser()).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/me`);
        expect(req.request.method).toBe('PUT');
        req.flush(createMockUser());
    });

    it('check update', () => {
        const user = createMockUser();
        service.update(user).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}`);
        expect(req.request.method).toBe('PUT');
        req.flush(user);
    });

    it('check findAll', () => {
        service.findAll().subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}`);
        expect(req.request.method).toBe('GET');
        req.flush([createMockUser()]);
    });

    it('check disable', () => {
        const user = createMockUser();
        service.disable(user).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.id}`);
        expect(req.request.method).toBe('PATCH');
        req.flush(user);
    });

    it('check valid delete', () => {
        const user = createMockUser();
        service.delete(user).subscribe({
            next: (res) => expect(res).toBeTrue(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true);
    });

    it('check invalid delete', () => {
        const user = createMockUser();
        service.delete(user).subscribe({
            next: (res) => expect(res).toBeFalse(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true, { status: 400, statusText: 'Bad Request' });
    });
});
