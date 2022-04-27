import { UserServiceMock } from './user.service.mock';
import { createMockUser } from '../../../types/user';

describe('UserServiceMock', () => {
    const userServiceMock = new UserServiceMock();
    const user = createMockUser();

    it('check delete', () => {
        const observable$ = userServiceMock.delete(user);
        observable$.subscribe((res) => expect(res).toBeTrue());
    });

    it('check disable', () => {
        const observable$ = userServiceMock.disable(user);
        observable$.subscribe((res) => expect(res).toEqual(user));
    });

    it('check findMe', () => {
        const observable$ = userServiceMock.findMe();
        observable$.subscribe((res) => expect(res).toBeUndefined());
    });

    it('check login', () => {
        const observable$ = userServiceMock.login(user);
        observable$.subscribe((res) => expect(res).not.toBeUndefined());
    });

    it('check update', () => {
        const observable$ = userServiceMock.update(user);
        observable$.subscribe((res) => expect(res).toEqual(user));
    });

    it('check updateMe', () => {
        const observable$ = userServiceMock.updateMe(user);
        observable$.subscribe((res) => expect(res).toEqual(user));
    });

    it('check register', () => {
        const observable$ = userServiceMock.register(user);
        observable$.subscribe((res) => expect(res).not.toBeUndefined());
    });
});
