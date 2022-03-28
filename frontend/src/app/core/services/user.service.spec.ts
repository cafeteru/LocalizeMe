import { TestBed } from '@angular/core/testing';

import { LoginData, UserService } from './user.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { ResponseLogin } from '../../types/response-login';
import { createAppStateMock } from '../../store/mocks/create-app-state-mock';
import { createMockUser } from '../../types/user';

describe('UserService', () => {
    let service: UserService;
    let mockHttp: HttpTestingController;
    const appState = createAppStateMock();

    beforeEach(() => {
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
        const loginData: LoginData = {
            Email: 'username',
            Password: 'password',
        };
        const token: ResponseLogin = {
            Authorization:
                'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY3RpdmUiOnRydWUsIk' +
                'FkbWluIjp0cnVlLCJFbWFpbCI6ImFkbWluQGVtYWlsLmVzIiwiSUQiOiI2MjI' +
                'xMmI5MmFiNjMxNDFhNjg0NzM5ZjMiLCJleHAiOjE2NDg1MzM5MTJ9.mWhtyw3' +
                'B8HdW9j1wpW0Plx-pI4OSUQmRx1parMW1gko',
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
        expect(appState.user.Email).toEqual('');
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
        const loginData: LoginData = {
            Email: 'username',
            Password: 'password',
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
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.ID}`);
        expect(req.request.method).toBe('PATCH');
        req.flush(user);
    });

    it('check valid delete', () => {
        const user = createMockUser();
        service.delete(user).subscribe({
            next: (res) => expect(res).toBeTrue(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.ID}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true);
    });

    it('check invalid delete', () => {
        const user = createMockUser();
        service.delete(user).subscribe({
            next: (res) => expect(res).toBeFalse(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.urlUsers}/${user.ID}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true, { status: 400, statusText: 'Bad Request' });
    });
});
