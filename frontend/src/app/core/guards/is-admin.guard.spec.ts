import { TestBed } from '@angular/core/testing';

import { IsAdminGuard } from './is-admin.guard';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';

describe('IsAdminGuard', () => {
    let guard: IsAdminGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(IsAdminGuard);
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });
});
