import { TestBed } from '@angular/core/testing';

import { CheckTokenGuard } from './check-token.guard';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from '../../app-routing';
import { provideMockStore } from '@ngrx/store/testing';
import { createAppStateMock } from '../../store/mocks/create-app-state-mock';

describe('CheckTokenGuard', () => {
    let guard: CheckTokenGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [RouterTestingModule.withRoutes(routes)],
            providers: [provideMockStore({ initialState: createAppStateMock() })],
        });
        guard = TestBed.inject(CheckTokenGuard);
    });

    afterAll(() => {
        localStorage.clear();
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });

    it('check canActivate valid exp', () => {
        const value = new Date().getTime() / 1_000 + 30;
        localStorage.setItem('exp', value.toString());
        expect(guard.canActivate()).toBeTrue();
    });

    it('check canActivate expired exp', () => {
        const value = new Date().getTime() / 1_000 - 30;
        localStorage.setItem('exp', value.toString());
        expect(guard.canActivate()).toBeFalse();
    });

    it('check canActivate without valid exp', () => {
        localStorage.setItem('exp', '1a');
        expect(guard.canActivate()).toBeFalse();
    });
});
