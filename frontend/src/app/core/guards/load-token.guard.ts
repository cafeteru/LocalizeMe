import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { map, Observable, of, tap } from 'rxjs';
import { AppState } from '../../store/app.reducer';
import { Store } from '@ngrx/store';
import jwt_decode from 'jwt-decode';
import { IToken } from '../../types/itoken';
import { UserReducer } from '../../store/reducers/user.reducer';
import * as userActions from '../../store/actions/user.actions';

@Injectable({
    providedIn: 'root',
})
export class LoadTokenGuard implements CanActivate {
    constructor(private store: Store<AppState>) {}

    canActivate(): Observable<boolean> {
        const authorization = localStorage.authorization;
        if (!authorization) {
            this.clearUser();
            return of(true);
        }
        return this.store.select('userInfo').pipe(
            tap(() => this.loadUser(authorization)),
            map(() => true)
        );
    }

    private clearUser(): void {
        localStorage.clear();
        this.store.dispatch(userActions.clearUser());
    }

    private loadUser(authorization: string): void {
        const iToken = jwt_decode<IToken>(authorization);
        const reducer: UserReducer = {
            exp: iToken.exp,
            authorization: authorization,
            user: {
                id: iToken.id,
                email: iToken.email,
                active: iToken.active,
                admin: iToken.admin,
                password: '',
            },
        };
        this.store.dispatch(userActions.loadUser(reducer));
    }
}
