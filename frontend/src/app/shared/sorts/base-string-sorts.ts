import { checkNotNullParams, sortStrings } from './sort-columns';
import { BaseString } from '../../types/base-string';

export function sortBaseStringByIdentifier(a: BaseString, b: BaseString): number {
    const validParams = checkNotNullParams(a.identifier, b.identifier);
    return validParams === 0 ? sortStrings(a.identifier, b.identifier) : validParams;
}

export function sortBaseStringBySourceLanguage(a: BaseString, b: BaseString): number {
    const validParams = checkNotNullParams(a.sourceLanguage, b.sourceLanguage);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.sourceLanguage.isoCode, b.sourceLanguage.isoCode);
    return validNames === 0 ? sortStrings(a.sourceLanguage.isoCode, b.sourceLanguage.isoCode) : validNames;
}

export function sortBaseStringByGroup(a: BaseString, b: BaseString): number {
    const validParams = checkNotNullParams(a.group, b.group);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.group.name, b.group.name);
    return validNames === 0 ? sortStrings(a.group.name, b.group.name) : validNames;
}

export function sortBaseStringByAuthor(a: BaseString, b: BaseString): number {
    const validParams = checkNotNullParams(a.author, b.author);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.author.email, b.author.email);
    return validNames === 0 ? sortStrings(a.author.email, b.author.email) : validNames;
}

export function sortBaseStringByActive(a: BaseString, b: BaseString): number {
    return checkNotNullParams(a.active, b.active);
}
