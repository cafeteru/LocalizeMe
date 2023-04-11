import { TestBed } from '@angular/core/testing';

import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';
import { LoadTokenGuard } from './load-token.guard';
import { AppState } from '../../store/app.reducer';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';
import { IToken } from '../../types/itoken';

describe('LoadTokenGuard', () => {
    let guard: LoadTokenGuard;
    let appState: AppState;
    let store: MockStore;

    beforeEach(() => {
        appState = createMockAppState();
        TestBed.configureTestingModule({
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(LoadTokenGuard);
        store = TestBed.inject(MockStore);
    });

    afterEach(() => {
        store.resetSelectors();
        localStorage.clear();
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });

    it('check with authorization', (done) => {
        const today = new Date();
        const mockToken: IToken = {
            exp: new Date(today.setDate(today.getDate() + 1)).getTime() / 1000,
            id: '1234567890',
            email: 'test@example.com',
            active: true,
            admin: false,
        };
        const spy = spyOn(guard as any, 'getIToken').and.returnValue(mockToken);
        localStorage.setItem('authorization', 'test');
        guard.canActivate().subscribe((res) => {
            expect(res).toBeTruthy();
            expect(spy).toHaveBeenCalled();
            done();
        });
    });

    it('check with invalid authorization', (done) => {
        const today = new Date();
        const mockToken: IToken = {
            exp: new Date(today.setDate(today.getDate() - 1000)).getTime() / 1000,
            id: '1234567890',
            email: 'test@example.com',
            active: true,
            admin: false,
        };
        const spy = spyOn(guard as any, 'getIToken').and.returnValue(mockToken);
        localStorage.setItem('authorization', 'test');
        guard.canActivate().subscribe((res) => {
            expect(spy).toHaveBeenCalled();
            expect(res).toBeTruthy();
            done();
        });
    });

    it('should return a valid IToken object', () => {
        const authorization =
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmUiOnRydWU' +
            'sImFkbWluIjp0cnVlLCJlbWFpbCI6ImFkbWluQGVtYWlsLmVzIiwiZXhwIjox' +
            'NjUwMzIxMDE5LCJpZCI6IjYyMjEyYjkyYWI2MzE0MWE2ODQ3MzlmMyJ9.hGw' +
            '3Seg9PqVpLauF9XZiC_XhWNwBVWc-jbW5mCzARi4';
        const expectedToken: IToken = {
            active: true,
            admin: true,
            email: 'admin@email.es',
            exp: 1650321019,
            id: '62212b92ab63141a684739f3',
        };
        expect(guard['getIToken'](authorization)).toEqual(expectedToken);
    });

    it('should return false when there is no authorization token', (done) => {
        guard.canActivate().subscribe((res) => {
            expect(res).toBeFalse();
            done();
        });
    });
});
