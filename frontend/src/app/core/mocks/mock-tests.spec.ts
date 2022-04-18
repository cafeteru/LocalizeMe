import { matDialogRefMock } from './mock-tests';

describe('MockTests', () => {
    it('check close', () => {
        const observable$ = matDialogRefMock.close();
        observable$.subscribe((res) => expect(res).toBeUndefined());
    });
});
