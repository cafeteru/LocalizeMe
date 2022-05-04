import { checkNotNullParams, sortStrings } from './sort-columns';
import { Translation } from '../../types/translation';

export function sortTranslationByVersion(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.version, b.version);
    return validParams === 0 ? a.version - b.version : validParams;
}

export function sortTranslationByContent(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.content, b.content);
    return validParams === 0 ? sortStrings(a.content, b.content) : validParams;
}

export function sortTranslationByStage(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.stage, b.stage);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.stage.name, b.stage.name);
    return validNames === 0 ? sortStrings(a.stage.name, b.stage.name) : validNames;
}

export function sortTranslationByLanguage(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.language, b.language);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.language.isoCode, b.language.isoCode);
    return validNames === 0 ? sortStrings(a.language.isoCode, b.language.isoCode) : validNames;
}

export function sortTranslationByActive(a: Translation, b: Translation): number {
    return checkNotNullParams(a.active, b.active);
}
