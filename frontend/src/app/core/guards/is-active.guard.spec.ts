import { fakeAsync, TestBed } from '@angular/core/testing';

import { IsActiveGuard } from './is-active.guard';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from '../../app-routing';
import { AppState } from '../../store/app.reducer';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('IsActiveGuard', () => {
    let guard: IsActiveGuard;
    let appState: AppState;
    let store: MockStore;

    beforeEach(() => {
        appState = createMockAppState();
        TestBed.configureTestingModule({
            imports: [RouterTestingModule.withRoutes(routes), HttpClientTestingModule],
            providers: [provideMockStore({ initialState: appState })],
        });
        guard = TestBed.inject(IsActiveGuard);
        store = TestBed.inject(MockStore);
    });

    afterEach(() => {
        store.resetSelectors();
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });

    it('check when user is active', fakeAsync(() => {
        const newAppState: AppState = {
            ...appState,
            userInfo: {
                ...appState.userInfo,
                user: {
                    ...appState.userInfo.user,
                    active: true,
                },
            },
        };
        store.setState(newAppState);
        guard.canActivate().subscribe((res) => {
            expect(res).toEqual(true);
        });
    }));

    it('check when user is not active', fakeAsync(() => {
        const newAppState: AppState = {
            ...appState,
            userInfo: {
                ...appState.userInfo,
                user: {
                    ...appState.userInfo.user,
                    active: false,
                },
            },
        };
        store.setState(newAppState);
        guard.canActivate().subscribe((res) => {
            expect(res).toEqual(false);
        });
    }));
});
