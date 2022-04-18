import { User } from '../../types/user';
import { sortUsersByEmail, sortUsersByIsActive, sortUsersByIsAdmin } from './users-sorts';

describe('users-sorts', () => {
    let a: User;
    let b: User;
    beforeEach(() => {
        a = {
            id: '1',
            admin: true,
            email: 'a',
            active: true,
            password: 'a',
        };
        b = {
            id: '12',
            admin: false,
            email: 'a1',
            active: false,
            password: 'a1',
        };
    });

    it('check sortEmail', () => {
        expect(sortUsersByEmail(a, b)).toBeLessThanOrEqual(1);
        expect(sortUsersByEmail(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortUsersByEmail(a, a)).toBe(0);
        b.email = undefined;
        expect(sortUsersByEmail(a, b)).toBe(-1);
        a.email = undefined;
        expect(sortUsersByEmail(a, b)).toBe(1);
    });

    it('check sortIsActive', () => {
        expect(sortUsersByIsActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortUsersByIsActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortUsersByIsActive(a, a)).toBe(0);
        b.active = undefined;
        expect(sortUsersByIsActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortUsersByIsActive(a, b)).toBe(1);
    });

    it('check sortIsAdmin', () => {
        expect(sortUsersByIsAdmin(a, b)).toBe(-1);
        expect(sortUsersByIsAdmin(b, a)).toBe(1);
        expect(sortUsersByIsAdmin(a, a)).toBe(0);
        b.admin = undefined;
        expect(sortUsersByIsAdmin(a, b)).toBe(-1);
        a.admin = undefined;
        expect(sortUsersByIsAdmin(a, b)).toBe(1);
    });
});
