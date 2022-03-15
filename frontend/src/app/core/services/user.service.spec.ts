import { TestBed } from '@angular/core/testing';

import { LoginData, UserService } from './user.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { environment } from '../../../environments/environment';
import { ResponseLogin } from '../../types/response-login';
import { createAppStateMock } from '../mocks/mock-tests';

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
            email: 'username',
            password: 'password',
        };
        const token: ResponseLogin = {
            Authorization:
                'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV' +
                'tYWlsLmVzIiwiZXhwIjoxNjQ2NjU1MjAzLCJpc0FjdGl2ZSI6dHJ1ZSwiaXN' +
                'BZG1pbiI6dHJ1ZX0.Mf2ooBSFbDNZdRBCJCR_R2-59VzDwt6jMYHrW7PHeOk',
        };
        service.login(loginData).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${environment.urlApi}/login`);
        expect(req.request.method).toBe('POST');
        req.flush(token);
    });

    it('check logout', () => {
        service.logout();
        expect(appState.user.Email).toEqual('');
    });
});
