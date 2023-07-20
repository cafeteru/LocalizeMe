import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { AppState } from '../../store/app.reducer';
import { Store } from '@ngrx/store';
import { LoadTokenGuard } from './load-token.guard';

@Injectable({
    providedIn: 'root',
})
export class LoadTokenGuardMenu extends LoadTokenGuard {
    constructor(protected store: Store<AppState>) {
        super(store);
    }

    canActivate(): Observable<boolean> {
        return super.canActivate().pipe(map((canActivate) => true));
    }
}
