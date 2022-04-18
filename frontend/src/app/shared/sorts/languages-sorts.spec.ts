import { sortLanguagesByActive, sortLanguagesByDescription, sortLanguagesByIsoCode } from './languages-sorts';
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
        expect(sortLanguagesByIsoCode(a, b)).toBeLessThanOrEqual(1);
        expect(sortLanguagesByIsoCode(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortLanguagesByIsoCode(a, a)).toBe(0);
        b.isoCode = undefined;
        expect(sortLanguagesByIsoCode(a, b)).toBe(-1);
        a.isoCode = undefined;
        expect(sortLanguagesByIsoCode(a, b)).toBe(1);
    });

    it('check sortDescription', () => {
        expect(sortLanguagesByDescription(a, b)).toBeLessThanOrEqual(1);
        expect(sortLanguagesByDescription(b, a)).toBeGreaterThanOrEqual(1);
        expect(sortLanguagesByDescription(a, a)).toBe(0);
        b.description = undefined;
        expect(sortLanguagesByDescription(a, b)).toBe(-1);
        a.description = undefined;
        expect(sortLanguagesByDescription(a, b)).toBe(1);
    });

    it('check sortActive', () => {
        expect(sortLanguagesByActive(b, a)).toBeLessThanOrEqual(1);
        expect(sortLanguagesByActive(a, a)).toBe(0);
        expect(sortLanguagesByActive(a, b)).toBeGreaterThanOrEqual(-1);
        b.active = undefined;
        expect(sortLanguagesByActive(a, b)).toBe(-1);
        a.active = undefined;
        expect(sortLanguagesByActive(a, b)).toBe(1);
    });
});
