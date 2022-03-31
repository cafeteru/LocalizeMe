import { User } from '../../types/user';
import { sortEmail, sortIsActive, sortIsAdmin } from './users-sorts';

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
        expect(sortEmail(a, b)).toBeLessThanOrEqual(1);
        expect(sortEmail(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortEmail(a, a)).toBe(0);
        b.email = undefined;
        expect(sortEmail(a, b)).toBe(-1);
        a.email = undefined;
        expect(sortEmail(a, b)).toBe(1);
    });

    it('check sortIsActive', () => {
        expect(sortIsActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortIsActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortIsActive(a, a)).toBe(0);
        b.active = undefined;
        expect(sortIsActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortIsActive(a, b)).toBe(1);
    });

    it('check sortIsAdmin', () => {
        expect(sortIsAdmin(a, b)).toBe(-1);
        expect(sortIsAdmin(b, a)).toBe(1);
        expect(sortIsAdmin(a, a)).toBe(0);
        b.admin = undefined;
        expect(sortIsAdmin(a, b)).toBe(-1);
        a.admin = undefined;
        expect(sortIsAdmin(a, b)).toBe(1);
    });
});
