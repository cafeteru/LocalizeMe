import { TestBed } from '@angular/core/testing';

import { CheckTokenGuard } from './check-token.guard';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from '../../app-routing';

describe('CheckTokenGuard', () => {
    let guard: CheckTokenGuard;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [RouterTestingModule.withRoutes(routes)],
        });
        guard = TestBed.inject(CheckTokenGuard);
    });

    it('should be created', () => {
        expect(guard).toBeTruthy();
    });
});
