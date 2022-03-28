import { matDialogRefMock } from './mock-tests';
import { Observable } from 'rxjs';

describe('MockTests', () => {
    it('should be created', () => {
        const a$ = matDialogRefMock.close();
        a$.subscribe((res) => expect(res).toBeUndefined());
    });
});
