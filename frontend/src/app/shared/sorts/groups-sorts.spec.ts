import { Group } from '../../types/group';
import { createMockUser } from '../../types/user';
import { sortGroupByActive, sortGroupByName, sortGroupByOwnerEmail, sortGroupByPublic } from './groups-sorts';

describe('groups-sorts', () => {
    let a: Group;
    let b: Group;

    beforeEach(() => {
        a = {
            id: '1',
            active: true,
            name: 'a',
            owner: createMockUser(),
            permissions: [],
            public: true,
        };
        b = {
            id: '2',
            active: false,
            name: 'b',
            owner: createMockUser(),
            permissions: [],
            public: false,
        };
        b.owner.email = 'zzzz@email.es';
    });

    describe('sortGroupByName', () => {
        it('check with null name', () => {
            b.name = undefined;
            expect(sortGroupByName(a, b)).toBe(-1);
            a.name = undefined;
            expect(sortGroupByName(a, b)).toBe(1);
        });

        it('check with valid values', () => {
            expect(sortGroupByName(a, b)).toBeLessThanOrEqual(1);
            expect(sortGroupByName(b, a)).toBeGreaterThanOrEqual(1);
            expect(sortGroupByName(a, a)).toBe(0);
        });
    });

    describe('sortGroupByOwnerEmail', () => {
        it('check with null owner', () => {
            a.owner = undefined;
            expect(sortGroupByOwnerEmail(a, b)).toBeLessThanOrEqual(1);
        });

        it('check with null owner email', () => {
            b.owner.email = undefined;
            expect(sortGroupByOwnerEmail(a, b)).toBe(-1);
            a.owner.email = undefined;
            expect(sortGroupByOwnerEmail(a, b)).toBe(1);
        });

        it('check with valid values', () => {
            expect(sortGroupByOwnerEmail(a, b)).toBeLessThanOrEqual(1);
            expect(sortGroupByOwnerEmail(b, a)).toBeGreaterThanOrEqual(1);
            expect(sortGroupByOwnerEmail(a, a)).toBe(0);
        });
    });

    describe('sortGroupByActive', () => {
        it('check with undefined active', () => {
            b.active = undefined;
            expect(sortGroupByActive(a, b)).toBe(-1);
            a.active = undefined;
            expect(sortGroupByActive(a, b)).toBe(1);
        });

        it('check with valid values', () => {
            expect(sortGroupByActive(b, a)).toBeLessThanOrEqual(1);
            expect(sortGroupByActive(a, a)).toBe(0);
            expect(sortGroupByActive(a, b)).toBeGreaterThanOrEqual(-1);
        });
    });

    describe('sortGroupByPublic', () => {
        it('check with undefined public', () => {
            b.public = undefined;
            expect(sortGroupByPublic(a, b)).toBe(-1);
            a.public = undefined;
            expect(sortGroupByPublic(a, b)).toBe(1);
        });

        it('check with valid values', () => {
            expect(sortGroupByPublic(b, a)).toBeLessThanOrEqual(1);
            expect(sortGroupByPublic(a, a)).toBe(0);
            expect(sortGroupByPublic(a, b)).toBeGreaterThanOrEqual(-1);
        });
    });
});
