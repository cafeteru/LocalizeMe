import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Urls } from '../../shared/constants/urls';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { Group, GroupDto } from '../../types/group';

@Injectable({
    providedIn: 'root',
})
export class GroupService {
    url = `${environment.urlApi}/${Urls.groups}`;

    constructor(private httpClient: HttpClient) {}

    create(groupDto: GroupDto): Observable<Group> {
        return this.httpClient.post<Group>(this.url, groupDto, getDefaultHttpOptions());
    }

    disable(group: Group): Observable<Group> {
        return this.httpClient.patch<Group>(`${this.url}/${group.id}`, group, getDefaultHttpOptions());
    }

    findAll(): Observable<Group[]> {
        return this.httpClient
            .get<Group[]>(this.url, getDefaultHttpOptions())
            .pipe(map((groups) => (groups ? groups : [])));
    }

    update(group: Group): Observable<Group> {
        return this.httpClient.put<Group>(this.url, group, getDefaultHttpOptions());
    }
}
