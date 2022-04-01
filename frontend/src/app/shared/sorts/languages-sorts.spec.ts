import { sortActive, sortDescription, sortIsoCode } from './languages-sorts';
import { Language } from '../../types/language';

describe('languages-sorts', () => {
    let a: Language;
    let b: Language;

    beforeEach(() => {
        a = {
            id: '1',
            active: true,
            isoCode: 'eng',
            description: 'english',
        };
        b = {
            id: '2',
            active: false,
            isoCode: 'esp',
            description: 'spanish',
        };
    });

    it('check sortIsoCode', () => {
        expect(sortIsoCode(a, b)).toBeLessThanOrEqual(1);
        expect(sortIsoCode(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortIsoCode(a, a)).toBe(0);
        b.isoCode = undefined;
        expect(sortIsoCode(a, b)).toBe(-1);
        a.isoCode = undefined;
        expect(sortIsoCode(a, b)).toBe(1);
    });

    it('check sortDescription', () => {
        expect(sortDescription(a, b)).toBeLessThanOrEqual(1);
        expect(sortDescription(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortDescription(a, a)).toBe(0);
        b.description = undefined;
        expect(sortDescription(a, b)).toBe(-1);
        a.description = undefined;
        expect(sortDescription(a, b)).toBe(1);
    });

    it('check sortActive', () => {
        expect(sortActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortActive(a, a)).toBe(0);
        expect(sortActive(a, b)).toBeGreaterThanOrEqual(-1);
        b.active = undefined;
        expect(sortActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortActive(a, b)).toBe(1);
    });
});
