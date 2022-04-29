import { fakeAsync, TestBed } from '@angular/core/testing';

import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';
import { LoadTokenGuard } from './load-token.guard';
import { AppState } from '../../store/app.reducer';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';
import { of } from 'rxjs';

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

    it('check with authorization', fakeAsync(() => {
        const authorization =
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmUiOnRydWU' +
            'sImFkbWluIjp0cnVlLCJlbWFpbCI6ImFkbWluQGVtYWlsLmVzIiwiZXhwIjox' +
            'NjUwMzIxMDE5LCJpZCI6IjYyMjEyYjkyYWI2MzE0MWE2ODQ3MzlmMyJ9.hGw' +
            '3Seg9PqVpLauF9XZiC_XhWNwBVWc-jbW5mCzARi4';
        localStorage.setItem('authorization', authorization);
        guard.canActivate().subscribe((res) => {
            expect(res).toBeTrue();
        });
    }));
});
