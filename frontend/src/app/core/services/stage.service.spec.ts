import { TestBed } from '@angular/core/testing';

import { StageRequest, StageService } from './stage.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { LoginData } from './user.service';
import { ResponseLogin } from '../../types/response-login';
import { Stage } from '../../types/stage';

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

    it('check login', () => {
        const stageRequest: StageRequest = {
            Name: 'name',
        };
        const token: Stage = {
            ID: '',
            Name: 'name',
            Active: true,
        };
        service.create(stageRequest).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('POST');
        req.flush(token);
    });
});
