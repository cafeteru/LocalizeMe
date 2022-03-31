import { Stage } from '../../types/stage';
import { sortActive, sortName } from './stages-sorts';

describe('stages-sorts', () => {
    let a: Stage;
    let b: Stage;
    beforeEach(() => {
        a = {
            id: '1',
            active: true,
            name: 'a',
        };
        b = {
            id: '12',
            active: false,
            name: 'b',
        };
    });

    it('check sortName', () => {
        expect(sortName(a, b)).toBeLessThanOrEqual(1);
        expect(sortName(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortName(a, a)).toBe(0);
        b.name = undefined;
        expect(sortName(a, b)).toBe(-1);
        a.name = undefined;
        expect(sortName(a, b)).toBe(1);
    });

    it('check sortActive', () => {
        expect(sortActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortActive(a, a)).toBe(0);
        b.active = undefined;
        expect(sortActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortActive(a, b)).toBe(1);
    });
});
