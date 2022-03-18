import { AppState } from '../app.reducer';
import { initialState } from '../reducers/user.reducer';

export const createAppStateMock = (): AppState => {
    return {
        user: initialState,
    };
};
