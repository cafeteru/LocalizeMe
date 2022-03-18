import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { map, Observable, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import { Urls } from '../../shared/constants/urls';

@Injectable({
    providedIn: 'root',
})
export class IsActiveGuard implements CanActivate {
    constructor(private store: Store<AppState>, private router: Router) {}

    canActivate(): Observable<boolean> {
        return this.store.select('user').pipe(
            map((user) => user.IsActive),
            tap((res) => {
                if (!res) {
                    this.router.navigateByUrl(Urls.menu).then();
                }
            })
        );
    }
}
