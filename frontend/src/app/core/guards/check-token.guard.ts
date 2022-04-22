import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Urls } from '../../shared/constants/urls';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Injectable({
    providedIn: 'root',
})
export class CheckTokenGuard implements CanActivate {
    constructor(private store: Store<AppState>, private router: Router) {}

    canActivate(): boolean {
        const exp = localStorage.exp;
        if (!exp || isNaN(exp)) {
            return false;
        }
        const validToken = checkToken(exp);
        if (!validToken) {
            this.router.navigateByUrl(Urls.menu).then();
        }
        return validToken;
    }
}

export function checkToken(exp: string): boolean {
    const time = Number(exp) * 1_000;
    return Date.now() < time;
}
