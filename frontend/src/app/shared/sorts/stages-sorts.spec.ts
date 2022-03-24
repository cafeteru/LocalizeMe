import { Stage } from '../../types/stage';
import { sortActive, sortName } from './stages-sorts';

describe('stages-sorts', () => {
    let a: Stage;
    let b: Stage;
    beforeEach(() => {
        a = {
            ID: '1',
            Active: true,
            Name: 'a',
        };
        b = {
            ID: '12',
            Active: false,
            Name: 'b',
        };
    });

    it('check sortName', () => {
        expect(sortName(a, b)).toBeLessThanOrEqual(1);
        expect(sortName(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortName(a, a)).toBe(0);
        b.Name = undefined;
        expect(sortName(a, b)).toBe(-1);
        a.Name = undefined;
        expect(sortName(a, b)).toBe(1);
    });

    it('check sortActive', () => {
        expect(sortActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortActive(a, a)).toBe(0);
        b.Active = undefined;
        expect(sortActive(a, b)).toBe(-1);
        a.Active = undefined;
        expect(sortActive(a, b)).toBe(1);
    });
});
