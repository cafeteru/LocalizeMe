import { TestBed } from '@angular/core/testing';

import { IsActiveGuard } from './is-active.guard';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';

describe('IsActiveGuard', () => {
    let guard: IsActiveGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(IsActiveGuard);
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });
});
