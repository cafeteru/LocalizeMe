import { TestBed } from '@angular/core/testing';

import { LanguageService } from './language.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { createMockLanguage, Language } from '../../types/language';

describe('LanguageService', () => {
    let service: LanguageService;
    let mockHttp: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
        });
        service = TestBed.inject(LanguageService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    it('check create', () => {
        const language = createMockLanguage();
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

    it('check correct findAll', () => {
        const response = [createMockLanguage()];
        service.findAll().subscribe({
            next: (value) => expect(value).toEqual(response),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('GET');
        req.flush(response);
    });

    it('check findAll return null', () => {
        const response = [];
        service.findAll().subscribe({
            next: (value) => expect(value).toEqual(response),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('GET');
        req.flush(null);
    });

    it('check valid delete', () => {
        const language = createMockLanguage();
        service.delete(language).subscribe({
            next: (res) => expect(res).toBeTrue(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${language.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true);
    });

    it('check invalid delete', () => {
        const language = createMockLanguage();
        service.delete(language).subscribe({
            next: (res) => expect(res).toBeFalse(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${language.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true, { status: 400, statusText: 'Bad Request' });
    });

    it('check disable', () => {
        const language = createMockLanguage();
        service.disable(language).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${language.id}`);
        expect(req.request.method).toBe('PATCH');
        req.flush(language);
    });

    it('check update', () => {
        const language = createMockLanguage();
        service.update(language).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('PUT');
        req.flush(language);
    });
});
