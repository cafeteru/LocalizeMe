import { TestBed } from '@angular/core/testing';

import { LanguageService } from './language.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { createMockLanguage, Language } from '../../types/language';

describe('LanguageService', () => {
    let service: LanguageService;
    let mockHttp: HttpTestingController;
    let language: Language;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
        });
        service = TestBed.inject(LanguageService);
        mockHttp = TestBed.inject(HttpTestingController);
        language = createMockLanguage();
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    describe('create', () => {
        it('check valid', () => {
            const groupDto: Language = {
                description: language.description,
                isoCode: language.isoCode,
                id: undefined,
                active: true,
            };
            service.create(groupDto).subscribe({
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}`);
            expect(req.request.method).toBe('POST');
            req.flush(language);
        });
    });

    describe('findAll', () => {
        it('check correct', () => {
            const response = [language];
            service.findAll().subscribe({
                next: (value) => expect(value).toEqual(response),
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}`);
            expect(req.request.method).toBe('GET');
            req.flush(response);
        });

        it('should return an empty array if the request fails', () => {
            service.findAll().subscribe({
                next: (value) => expect(value).toEqual([]),
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}`);
            expect(req.request.method).toBe('GET');
            req.flush('Invalid request', { status: 400, statusText: 'Bad Request' });
        });
    });

    describe('delete', () => {
        it('check valid', () => {
            service.delete(language).subscribe({
                next: (res) => expect(res).toBeTrue(),
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}/${language.id}`);
            expect(req.request.method).toBe('DELETE');
            req.flush(true);
        });

        it('check invalid', () => {
            service.delete(language).subscribe({
                next: (res) => expect(res).toBeFalse(),
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}/${language.id}`);
            expect(req.request.method).toBe('DELETE');
            req.flush(true, { status: 400, statusText: 'Bad Request' });
        });
    });

    describe('disable', () => {
        it('check valid', () => {
            service.disable(language).subscribe({
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}/${language.id}`);
            expect(req.request.method).toBe('PATCH');
            req.flush(language);
        });
    });

    describe('update', () => {
        it('check valid', () => {
            service.update(language).subscribe({
                error: (err) => fail(err),
            });
            const req = mockHttp.expectOne(`${service.url}`);
            expect(req.request.method).toBe('PUT');
            req.flush(language);
        });
    });
});
