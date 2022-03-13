import { AppState } from '../../store/app.reducer';
import * as user from '../../store/reducers/user.reducer';
import { of } from 'rxjs';

export const createAppStateMock = (): AppState => {
    return {
        user: user.initialState,
    };
};

export const matDialogRefMock = {
    close: () => of(undefined),
};
