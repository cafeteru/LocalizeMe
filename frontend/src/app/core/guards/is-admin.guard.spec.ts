import { fakeAsync, TestBed } from '@angular/core/testing';

import { IsAdminGuard } from './is-admin.guard';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from '../../app-routing';
import { AppState } from '../../store/app.reducer';
import { createAppStateMock } from '../../store/mocks/create-app-state-mock';

describe('IsAdminGuard', () => {
    let guard: IsAdminGuard;
    let appState: AppState;
    let store: MockStore;

    beforeEach(() => {
        appState = createAppStateMock();
        TestBed.configureTestingModule({
            imports: [RouterTestingModule.withRoutes(routes)],
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(IsAdminGuard);
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
                    admin: true,
                },
            },
        };
        store.setState(newAppState);
        guard.canActivate().subscribe((res) => {
            expect(res).toBeTrue();
        });
    }));

    it('check when user is not active', fakeAsync(() => {
        const newAppState: AppState = {
            ...appState,
            userInfo: {
                ...appState.userInfo,
                user: {
                    ...appState.userInfo.user,
                    admin: false,
                },
            },
        };
        store.setState(newAppState);
        guard.canActivate().subscribe((res) => {
            expect(res).toEqual(false);
        });
    }));
});
