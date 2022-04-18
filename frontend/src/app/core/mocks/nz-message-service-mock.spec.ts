import { nzMessageServiceMock } from './nz-message-service-mock';

describe('NzMessageServiceMock', () => {
    it('check create', () => {
        const res = nzMessageServiceMock.create('');
        expect(res).toBeUndefined();
    });
});
