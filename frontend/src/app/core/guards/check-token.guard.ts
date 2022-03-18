import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Urls } from '../../shared/constants/urls';

@Injectable({
    providedIn: 'root',
})
export class CheckTokenGuard implements CanActivate {
    constructor(private router: Router) {}

    canActivate(): boolean {
        const authorization = localStorage.Authorization;
        const exp = localStorage.Exp;
        if (!authorization || !exp || isNaN(exp)) {
            return false;
        }
        const value = Date.now() < Number(exp) * 1_000;
        if (!value) {
            this.router.navigateByUrl(Urls.menu).then();
        }
        return value;
    }
}
