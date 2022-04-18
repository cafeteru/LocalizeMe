import { TestBed } from '@angular/core/testing';

import { StageService } from './stage.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { createMockStage, Stage, StageDto } from '../../types/stage';

describe('StageService', () => {
    let service: StageService;
    let mockHttp: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
        });
        service = TestBed.inject(StageService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    it('check create', () => {
        const stage = createMockStage();
        const stageRequest: StageDto = {
            name: stage.name,
        };
        service.create(stageRequest).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('POST');
        req.flush(stage);
    });

    it('check findAll', () => {
        const response = [createMockStage()];
        service.findAll().subscribe({
            next: (value) => expect(value).toEqual(response),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('GET');
        req.flush(response);
    });

    it('check null findAll', () => {
        service.findAll().subscribe({
            next: (value) => expect(value).toEqual([]),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('GET');
        req.flush(null);
    });

    it('check disable', () => {
        const stage = createMockStage();
        service.disable(stage).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${stage.id}`);
        expect(req.request.method).toBe('PATCH');
        req.flush(stage);
    });

    it('check valid delete', () => {
        const stage = createMockStage();
        service.delete(stage).subscribe({
            next: (res) => expect(res).toBeTrue(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${stage.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true);
    });

    it('check invalid delete', () => {
        const stage = createMockStage();
        service.delete(stage).subscribe({
            next: (res) => expect(res).toBeFalse(),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}/${stage.id}`);
        expect(req.request.method).toBe('DELETE');
        req.flush(true, { status: 400, statusText: 'Bad Request' });
    });

    it('check update', () => {
        const stage = createMockStage();
        service.update(stage).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('PUT');
        req.flush(stage);
    });
});
