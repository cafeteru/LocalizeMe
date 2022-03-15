import { TestBed } from '@angular/core/testing';

import { CheckTokenGuard } from './check-token.guard';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from '../../store/reducers/user.reducer';

describe('CheckTokenGuard', () => {
    let guard: CheckTokenGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideMockStore({ initialState })],
        });
        guard = TestBed.inject(CheckTokenGuard);
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });
});
