import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { catchError, Observable, of, switchMap, tap } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { UserDto } from '../types/user';
import { ResponseLogin } from '../types/response-login';

@Injectable({
    providedIn: 'root',
})
export class LocalizeMeService {
    url = `${environment.urlApi}`;
    baseStringUrl = `${environment.urlApi}/baseStrings`;

    constructor(private httpClient: HttpClient) {}

    login(): Observable<any> {
        const userDto: UserDto = {
            email: 'uo239795@uniovi.es',
            password: '123456',
        };
        return this.httpClient
            .post<ResponseLogin>(`${this.url}/login`, userDto)
            .pipe(
                tap((responseLogin) => localStorage.setItem('localize-me-authorization', responseLogin.authorization))
            );
    }

    findByIdentifierAndLanguage(identifier: string, isoCode: string): Observable<string> {
        return this.login().pipe(
            switchMap(() =>
                this.httpClient
                    .get<string>(
                        `${this.baseStringUrl}/content?identifier=${identifier}&isoCode=${isoCode}`,
                        getDefaultHttpOptions()
                    )
                    .pipe(catchError(() => of(identifier)))
            )
        );
    }
}
