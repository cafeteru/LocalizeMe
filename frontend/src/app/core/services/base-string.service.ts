import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Urls } from '../../shared/constants/urls';
import { HttpClient } from '@angular/common/http';
import { catchError, map, Observable, of } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { BaseString } from '../../types/base-string';
import { Stage } from '../../types/stage';
import { Group } from '../../types/group';

@Injectable({
    providedIn: 'root',
})
export class BaseStringService {
    url = `${environment.urlApi}/${Urls.baseStrings}`;

    constructor(private httpClient: HttpClient) {}

    create(baseString: BaseString): Observable<BaseString> {
        return this.httpClient.post<BaseString>(this.url, baseString, getDefaultHttpOptions());
    }

    delete(baseString: BaseString): Observable<boolean> {
        return this.httpClient
            .delete<boolean>(`${this.url}/${baseString.id}`, getDefaultHttpOptions())
            .pipe(catchError(() => of(false)));
    }

    disable(baseString: BaseString): Observable<BaseString> {
        return this.httpClient.patch<BaseString>(`${this.url}/${baseString.id}`, baseString, getDefaultHttpOptions());
    }

    findAll(): Observable<BaseString[]> {
        return this.httpClient
            .get<BaseString[]>(this.url, getDefaultHttpOptions())
            .pipe(map((baseStrings) => (baseStrings ? baseStrings : [])));
    }

    read(stage: Stage, group: Group, xliff: string): Observable<BaseString[]> {
        return this.httpClient.post<BaseString[]>(
            `${environment.urlApi}/${Urls.xliffs}?stage=${stage.id}&group=${group.id}`,
            xliff,
            getDefaultHttpOptions()
        );
    }

    update(baseString: BaseString): Observable<BaseString> {
        return this.httpClient.put<BaseString>(this.url, baseString, getDefaultHttpOptions());
    }
}
