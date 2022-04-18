import { Stage } from '../../types/stage';
import { sortStagesByActive, sortStagesByName } from './stages-sorts';

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
        expect(sortStagesByName(a, b)).toBeLessThanOrEqual(1);
        expect(sortStagesByName(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortStagesByName(a, a)).toBe(0);
        b.name = undefined;
        expect(sortStagesByName(a, b)).toBe(-1);
        a.name = undefined;
        expect(sortStagesByName(a, b)).toBe(1);
    });

    it('check sortActive', () => {
        expect(sortStagesByActive(a, b)).toBeGreaterThanOrEqual(-1);
        expect(sortStagesByActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortStagesByActive(a, a)).toBe(0);
        b.active = undefined;
        expect(sortStagesByActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortStagesByActive(a, b)).toBe(1);
    });
});
