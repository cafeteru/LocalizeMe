import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { map, Observable, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import { Urls } from '../../shared/constants/urls';

@Injectable({
    providedIn: 'root',
})
export class IsAdminGuard implements CanActivate {
    constructor(private store: Store<AppState>, private router: Router) {}

    canActivate(): Observable<boolean> {
        return this.store.select('userInfo').pipe(
            map((userReducer) => userReducer.user.admin),
            tap((isAdmin) => {
                if (!isAdmin) {
                    this.router.navigateByUrl(Urls.menu).then();
                }
            })
        );
    }
}
