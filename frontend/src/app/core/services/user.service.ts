import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment';
import { ResponseLogin } from '../../types/response-login';
import { catchError, map, Observable, of } from 'rxjs';
import jwt_decode from 'jwt-decode';
import { IToken } from '../../types/itoken';
import { UserReducer } from '../../store/reducers/user.reducer';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import * as userActions from '../../store/actions/user.actions';
import { createMockUser, LoginData, User } from '../../types/user';
import { getDefaultHttpOptions } from './default-http-options';
import { Urls } from '../../shared/constants/urls';

@Injectable({
    providedIn: 'root',
})
export class UserService {
    url = `${environment.urlApi}`;
    urlUsers = `${environment.urlApi}/${Urls.users}`;

    constructor(private httpClient: HttpClient, private store: Store<AppState>) {}

    delete(user: User): Observable<boolean> {
        return this.httpClient.delete<User>(`${this.urlUsers}/${user.id}`, getDefaultHttpOptions()).pipe(
            map(() => true),
            catchError(() => of(false))
        );
    }

    disable(user: User): Observable<User> {
        return this.httpClient.patch<User>(`${this.urlUsers}/${user.id}`, user, getDefaultHttpOptions());
    }

    findAll(): Observable<User[]> {
        return this.httpClient.get<User[]>(this.urlUsers, getDefaultHttpOptions());
    }

    findMe(): Observable<User> {
        return this.httpClient.get<User>(`${this.urlUsers}/me`, getDefaultHttpOptions());
    }

    login(loginData: LoginData): Observable<User> {
        return this.httpClient.post<ResponseLogin>(`${this.url}/login`, loginData).pipe(
            map((responseLogin) => {
                const iToken = jwt_decode<IToken>(responseLogin.authorization);
                if (iToken.active) {
                    const userReducer: UserReducer = {
                        exp: iToken.exp,
                        authorization: responseLogin.authorization,
                        user: {
                            id: iToken.id,
                            email: iToken.email,
                            active: iToken.active,
                            admin: iToken.admin,
                            password: '',
                        },
                    };
                    localStorage.setItem('authorization', responseLogin.authorization);
                    localStorage.setItem('exp', iToken.exp.toString());
                    this.store.dispatch(userActions.loadUser(userReducer));
                }
                return {
                    ...createMockUser(),
                    email: iToken.email,
                    active: iToken.active,
                    admin: iToken.admin,
                };
            })
        );
    }

    logout(): void {
        localStorage.clear();
        this.store.dispatch(userActions.clearUser());
    }

    update(user: User): Observable<User> {
        return this.httpClient.put<User>(`${this.urlUsers}`, user, getDefaultHttpOptions());
    }

    updateMe(user: User): Observable<User> {
        return this.httpClient.put<User>(`${this.urlUsers}/me`, user, getDefaultHttpOptions());
    }

    register(loginData: LoginData): Observable<User> {
        return this.httpClient.post<User>(this.urlUsers, loginData);
    }
}
