import { Observable, of } from 'rxjs';

export class DialogMock {
    open(): any {
        return null;
    }
}

export class DialogRefMock {
    close(): Observable<any> {
        return of(undefined);
    }
}
