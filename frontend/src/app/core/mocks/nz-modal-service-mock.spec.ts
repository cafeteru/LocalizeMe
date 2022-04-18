import { nzModalServiceMock } from './nz-modal-service-mock';

describe('NzModalServiceMock', () => {
    it('check confirm', () => {
        const res = nzModalServiceMock.confirm();
        expect(res).toBeUndefined();
    });
});
