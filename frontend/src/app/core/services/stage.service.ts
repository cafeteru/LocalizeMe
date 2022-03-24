import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { Stage } from '../../types/stage';
import { getDefaultHttpOptions } from './default-http-options';

export interface StageRequest {
    Name: string;
}

@Injectable({
    providedIn: 'root',
})
export class StageService {
    url = `${environment.urlApi}/stages`;

    constructor(private httpClient: HttpClient) {}

    create(stageRequest: StageRequest): Observable<Stage> {
        return this.httpClient.post<Stage>(this.url, stageRequest, getDefaultHttpOptions());
    }

    findAll(): Observable<Stage[]> {
        return this.httpClient
            .get<Stage[]>(this.url, getDefaultHttpOptions())
            .pipe(map((stages) => (stages ? stages : [])));
    }
}