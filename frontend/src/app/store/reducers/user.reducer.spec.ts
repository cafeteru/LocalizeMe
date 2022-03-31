import { clearUser, getUserReducer, initialState, loadUser, UserReducer } from './user.reducer';

describe('UserReducer', () => {
    const userReducer: UserReducer = {
        id: initialState.id,
        email: initialState.email,
        exp: initialState.exp,
        active: initialState.active,
        admin: initialState.admin,
        authorization: initialState.authorization,
    };

    it('check loadUser', () => {
        const temp: UserReducer = {
            id: '1',
            email: 'email@email.es',
            exp: 1,
            active: true,
            admin: true,
            authorization: 'Authorization',
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
