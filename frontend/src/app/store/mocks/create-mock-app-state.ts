import { AppState } from '../app.reducer';
import { initialState } from '../reducers/user.reducer';

export const createMockAppState = (): AppState => {
    return {
        userInfo: initialState,
    };
};
