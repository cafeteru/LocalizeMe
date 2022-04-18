import { clearUser, getUserReducer, initialState, loadUser, UserReducer } from './user.reducer';
import { createMockUser } from '../../types/user';

describe('UserReducer', () => {
    const userReducer: UserReducer = { ...initialState };

    it('check loadUser', () => {
        const temp: UserReducer = {
            exp: 1,
            authorization: 'authorization',
            user: createMockUser(),
        };
        const reduce = loadUser(userReducer, temp);
        expect(reduce).toEqual(temp);
        expect(reduce).not.toEqual(userReducer);
    });

    it('check clearUser', () => {
        const reduce = clearUser(userReducer);
        expect(reduce).toEqual(userReducer);
    });

    it('check getLoadedReducer', () => {
        const loadedActionReducer = getUserReducer();
        expect(loadedActionReducer).not.toBeNull();
    });
});
