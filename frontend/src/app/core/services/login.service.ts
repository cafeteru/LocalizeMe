import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { ResponseLogin } from '../../types/response-login';
import { map, Observable } from 'rxjs';
import jwt_decode from 'jwt-decode';
import { IToken } from '../../types/itoken';
import { UserReducer } from '../../store/reducers/user.reducer';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import * as userActions from '../../store/actions/user.actions';

export interface LoginData {
    email: string;
    password: string;
}

@Injectable({
    providedIn: 'root',
})
export class LoginService {
    private url = `${environment.urlApi}/login`;

    constructor(private httpClient: HttpClient, private store: Store<AppState>) {}

    login(loginData: LoginData): Observable<void> {
        return this.httpClient.post<ResponseLogin>(this.url, loginData).pipe(
            map((res) => {
                const iToken = jwt_decode<IToken>(res.Authorization);
                const userReducer: UserReducer = {
                    Email: iToken.Email,
                    Exp: iToken.exp,
                    IsActive: iToken.IsActive,
                    IsAdmin: iToken.IsAdmin,
                    Authorization: res.Authorization,
                };
                this.store.dispatch(userActions.loadUser(userReducer));
            })
        );
    }

    logout(): void {
        this.store.dispatch(userActions.clearUser());
    }
}
