import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';
import { catchError, Observable, of } from 'rxjs';
import { getDefaultHttpOptions } from './default-http-options';
import { Urls } from '../../shared/constants/urls';
import { Language, LanguageDto } from '../../types/language';

@Injectable({
    providedIn: 'root',
})
export class LanguageService {
    url = `${environment.urlApi}/${Urls.languages}`;

    constructor(private httpClient: HttpClient) {}

    create(languageDto: LanguageDto): Observable<Language> {
        return this.httpClient.post<Language>(this.url, languageDto, getDefaultHttpOptions());
    }

    delete(language: Language): Observable<boolean> {
        return this.httpClient
            .delete<boolean>(`${this.url}/${language.id}`, getDefaultHttpOptions())
            .pipe(catchError(() => of(false)));
    }

    disable(language: Language): Observable<Language> {
        return this.httpClient.patch<Language>(`${this.url}/${language.id}`, language, getDefaultHttpOptions());
    }

    findAll(): Observable<Language[]> {
        return this.httpClient.get<Language[]>(this.url, getDefaultHttpOptions()).pipe(catchError(() => of([])));
    }

    update(language: Language): Observable<Language> {
        return this.httpClient.put<Language>(this.url, language, getDefaultHttpOptions());
    }
}
