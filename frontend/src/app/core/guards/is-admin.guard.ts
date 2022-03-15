import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { map, Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Injectable({
    providedIn: 'root',
})
export class IsAdminGuard implements CanActivate {
    constructor(private store: Store<AppState>) {}

    canActivate(): Observable<boolean> {
        return this.store.select('user').pipe(map((user) => user.IsAdmin));
    }
}
