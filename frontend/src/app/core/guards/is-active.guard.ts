import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { map, Observable, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import { Urls } from '../../shared/constants/urls';
import { UserService } from '../services/user.service';

@Injectable({
    providedIn: 'root',
})
export class IsActiveGuard implements CanActivate {
    constructor(private store: Store<AppState>, private userService: UserService, private router: Router) {}

    canActivate(): Observable<boolean> {
        return this.store.select('userInfo').pipe(
            map((userReducer) => userReducer.user.active),
            tap((isActive) => {
                if (!isActive) {
                    this.userService.logout();
                    this.router.navigateByUrl(Urls.menu).then();
                }
            })
        );
    }
}
