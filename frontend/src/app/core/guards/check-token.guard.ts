import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { AppState } from '../../store/app.reducer';
import { Store } from '@ngrx/store';

@Injectable({
    providedIn: 'root',
})
export class CheckTokenGuard implements CanActivate {
    constructor() {}

    canActivate(): boolean {
        const authorization = localStorage.Authorization;
        const exp = localStorage.Exp;
        if (!authorization || !exp || isNaN(exp)) {
            return false;
        }
        return Date.now() < Number(exp) * 1_000;
    }
}
