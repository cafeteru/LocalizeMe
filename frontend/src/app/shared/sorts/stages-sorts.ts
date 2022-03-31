import { checkNotNullParams, sortStrings } from './sort-columns';
import { Stage } from '../../types/stage';

export function sortName(a: Stage, b: Stage): number {
    const validParams = checkNotNullParams(a.name, b.name);
    return validParams === 0 ? sortStrings(a.name, b.name) : validParams;
}

export function sortActive(a: Stage, b: Stage): number {
    return checkNotNullParams(a.active, b.active);
}
