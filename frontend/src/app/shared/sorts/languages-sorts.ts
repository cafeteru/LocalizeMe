import { checkNotNullParams, sortStrings } from './sort-columns';
import { Stage } from '../../types/stage';
import { Language } from '../../types/language';

export function sortDescription(a: Language, b: Language): number {
    const validParams = checkNotNullParams(a.description, b.description);
    return validParams === 0 ? sortStrings(a.description, b.description) : validParams;
}

export function sortIsoCode(a: Language, b: Language): number {
    const validParams = checkNotNullParams(a.isoCode, b.isoCode);
    return validParams === 0 ? sortStrings(a.isoCode, b.isoCode) : validParams;
}

export function sortActive(a: Language, b: Language): number {
    return checkNotNullParams(a.active, b.active);
}
