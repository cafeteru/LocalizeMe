import { clearUser, getUserReducer, initialState, loadUser, UserReducer } from './user.reducer';

describe('UserReducer', () => {
    const userReducer: UserReducer = {
        ID: initialState.ID,
        Email: initialState.Email,
        Exp: initialState.Exp,
        Active: initialState.Active,
        Admin: initialState.Admin,
        Authorization: initialState.Authorization,
    };

    it('check loadUser', () => {
        const temp: UserReducer = {
            ID: '1',
            Email: 'email@email.es',
            Exp: 1,
            Active: true,
            Admin: true,
            Authorization: 'Authorization',
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
