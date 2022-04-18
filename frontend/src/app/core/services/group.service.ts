import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Urls } from '../../shared/constants/urls';
import { HttpClient } from '@angular/common/http';
import { Language, LanguageDto } from '../../types/language';
import { Observable } from 'rxjs';
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
}
