import { TestBed } from '@angular/core/testing';

import { BaseStringService } from './base-string.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

describe('BaseStringService', () => {
    let service: BaseStringService;
    let mockHttp: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
        });
        service = TestBed.inject(BaseStringService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });
});
