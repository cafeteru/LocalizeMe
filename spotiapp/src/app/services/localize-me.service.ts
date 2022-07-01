import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, of, tap } from 'rxjs';
import { environment } from '../../environments/environment';
import { ResponseLogin } from '../types/response-login';
import { UserDto } from '../types/user';
import { getDefaultHttpOptions } from './default-http-options';

@Injectable({
    providedIn: 'root',
})
export class LocalizeMeService {
    url = `${environment.urlApi}`;
    baseStringUrl = `${environment.urlApi}/baseStrings`;

    constructor(private httpClient: HttpClient) {}

    login(): Observable<any> {
        const userDto: UserDto = {
            email: environment.email,
            password: environment.password,
        };
        return this.httpClient
            .post<ResponseLogin>(`${this.url}/login`, userDto)
            .pipe(
                tap((responseLogin) => localStorage.setItem('localize-me-authorization', responseLogin.authorization))
            );
    }

    findByIdentifierAndLanguage(identifier: string, isoCode: string): Observable<string> {
        return this.httpClient
            .get<string>(
                `${this.baseStringUrl}/content?identifier=${identifier}&isoCode=${isoCode}`,
                getDefaultHttpOptions()
            )
            .pipe(catchError(() => of(identifier)));
    }

    findByIdentifierAndLanguageAndStage(identifier: string, isoCode: string, stageName: string): Observable<string> {
        return this.httpClient
            .get<string>(
                `${this.baseStringUrl}/content/env?identifier=${identifier}&isoCode=${isoCode}&stage=${stageName}`,
                getDefaultHttpOptions()
            )
            .pipe(catchError(() => of(identifier)));
    }
}
