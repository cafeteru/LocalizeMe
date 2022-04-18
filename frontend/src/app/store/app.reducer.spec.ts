import { getAppReducers } from './app.reducer';

describe('AppReducer', () => {
    it('check getAppReducers', () => {
        const appReducers = getAppReducers();
        expect(appReducers).not.toBeNull();
        expect(appReducers.userInfo).not.toBeUndefined();
    });
});
