import { User } from '../../types/user';
import { sortEmail, sortIsActive, sortIsAdmin } from './users-sorts';

describe('users-sorts', () => {
    let a: User;
    let b: User;
    beforeEach(() => {
        a = {
            ID: '1',
            Admin: true,
            Email: 'a',
            Active: true,
            Password: 'a',
        };
        b = {
            ID: '12',
            Admin: false,
            Email: 'a1',
            Active: false,
            Password: 'a1',
        };
    });

    it('check sortEmail', () => {
        expect(sortEmail(a, b)).toBeLessThanOrEqual(1);
        expect(sortEmail(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortEmail(a, a)).toBe(0);
        b.Email = undefined;
        expect(sortEmail(a, b)).toBe(-1);
        a.Email = undefined;
        expect(sortEmail(a, b)).toBe(1);
    });

    it('check sortIsActive', () => {
        expect(sortIsActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortIsActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortIsActive(a, a)).toBe(0);
        b.Active = undefined;
        expect(sortIsActive(a, b)).toBe(-1);
        a.Active = undefined;
        expect(sortIsActive(a, b)).toBe(1);
    });

    it('check sortIsAdmin', () => {
        expect(sortIsAdmin(a, b)).toBe(-1);
        expect(sortIsAdmin(b, a)).toBe(1);
        expect(sortIsAdmin(a, a)).toBe(0);
        b.Admin = undefined;
        expect(sortIsAdmin(a, b)).toBe(-1);
        a.Admin = undefined;
        expect(sortIsAdmin(a, b)).toBe(1);
    });
});
