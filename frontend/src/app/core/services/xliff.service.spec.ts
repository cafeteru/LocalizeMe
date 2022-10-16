import { TestBed } from '@angular/core/testing';

import { UserService } from './user.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { ResponseLogin } from '../../types/response-login';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';
import { createMockUser, User } from '../../types/user';
import { AppState } from '../../store/app.reducer';
import { XliffService } from './xliff.service';

describe('XliffService', () => {
    let service: XliffService;
    let mockHttp: HttpTestingController;
    let appState: AppState;

    beforeEach(() => {
        appState = createMockAppState();
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
            providers: [provideMockStore({ initialState: appState })],
        });
        service = TestBed.inject(XliffService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
