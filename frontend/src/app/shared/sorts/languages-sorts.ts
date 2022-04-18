import { checkNotNullParams, sortStrings } from './sort-columns';
import { Language } from '../../types/language';

export function sortLanguagesByDescription(a: Language, b: Language): number {
    const validParams = checkNotNullParams(a.description, b.description);
    return validParams === 0 ? sortStrings(a.description, b.description) : validParams;
}

export function sortLanguagesByIsoCode(a: Language, b: Language): number {
    const validParams = checkNotNullParams(a.isoCode, b.isoCode);
    return validParams === 0 ? sortStrings(a.isoCode, b.isoCode) : validParams;
}

export function sortLanguagesByActive(a: Language, b: Language): number {
    return checkNotNullParams(a.active, b.active);
}
