import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { map, Observable, of, tap } from 'rxjs';
import { AppState } from '../../store/app.reducer';
import { Store } from '@ngrx/store';
import jwt_decode from 'jwt-decode';
import { IToken } from '../../types/itoken';
import { UserReducer } from '../../store/reducers/user.reducer';
import * as userActions from '../../store/actions/user.actions';
import { checkToken } from './check-token.guard';

@Injectable({
    providedIn: 'root',
})
export class LoadTokenGuard implements CanActivate {
    constructor(private store: Store<AppState>) {}

    canActivate(): Observable<boolean> {
        const { authorization } = localStorage;
        if (authorization) {
            return this.store.select('userInfo').pipe(
                tap(() => this.loadUser(authorization)),
                map(() => true)
            );
        }
        this.clearUser();
        return of(false);
    }

    private clearUser(): void {
        localStorage.clear();
        this.store.dispatch(userActions.clearUser());
    }

    private loadUser(authorization: string): void {
        const { exp, id, email, active, admin } = this.getIToken(authorization);
        if (checkToken(exp)) {
            const reducer: UserReducer = {
                exp,
                authorization,
                user: {
                    id: id,
                    email: email,
                    active: active,
                    admin: admin,
                    password: '',
                },
            };
            this.store.dispatch(userActions.loadUser(reducer));
        } else {
            this.clearUser();
        }
    }

    private getIToken(authorization: string): IToken {
        return jwt_decode<IToken>(authorization);
    }
}
