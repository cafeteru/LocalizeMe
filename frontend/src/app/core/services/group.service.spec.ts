import { TestBed } from '@angular/core/testing';

import { GroupService } from './group.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { createMockGroup, Group } from '../../types/group';
import { createMockUser } from '../../types/user';

describe('GroupService', () => {
    let service: GroupService;
    let mockHttp: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
        });
        service = TestBed.inject(GroupService);
        mockHttp = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        mockHttp.verify();
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    it('check create', () => {
        const group = createMockGroup();
        const groupDto: Group = {
            name: group.name,
            permissions: [],
            owner: createMockUser(),
            public: true,
            active: true,
            id: undefined,
        };
        service.create(groupDto).subscribe({
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('POST');
        req.flush(group);
    });

    it('check findAll', () => {
        const response = [createMockGroup()];
        service.findAll().subscribe({
            next: (groups) => expect(groups).toEqual(response),
            error: (err) => fail(err),
        });
        const req = mockHttp.expectOne(`${service.url}`);
        expect(req.request.method).toBe('GET');
        req.flush(response);
    });
});
