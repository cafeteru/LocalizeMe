import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Urls } from '../../shared/constants/urls';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { BaseString } from '../../types/base-string';
import { Stage } from '../../types/stage';
import { Group } from '../../types/group';
import { XliffDto } from '../../types/xliff';

@Injectable({
    providedIn: 'root',
})
export class XliffService {
    url = `${environment.urlApi}/${Urls.xliffs}`;

    constructor(private httpClient: HttpClient) {}

    createXliff(xliff: XliffDto): Observable<string> {
        xliff.baseStringIds = xliff.baseStringIds ? xliff.baseStringIds : [];
        return this.httpClient.post(`${this.url}/create`, xliff, {
            responseType: 'text',
            headers: new HttpHeaders({
                'Content-Type': 'application/json',
                authorization: `Bearer ${localStorage.authorization}`,
            }),
        });
    }

    read(stage: Stage, group: Group, xliff: string): Observable<BaseString[]> {
        return this.httpClient.post<BaseString[]>(
            `${this.url}?stage=${stage.id}&group=${group.id}`,
            xliff,
            getDefaultHttpOptions()
        );
    }
}
