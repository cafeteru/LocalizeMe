import { TestBed } from '@angular/core/testing';

import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';
import { LoadTokenGuard } from './load-token.guard';

describe('LoadTokenGuard', () => {
    let guard: LoadTokenGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(LoadTokenGuard);
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });
});
