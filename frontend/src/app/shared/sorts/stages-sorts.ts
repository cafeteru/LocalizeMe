import { checkNotNullParams, sortStrings } from './sort-columns';
import { Stage } from '../../types/stage';

export function sortName(a: Stage, b: Stage): number {
    const validParams = checkNotNullParams(a.Name, b.Name);
    return validParams === 0 ? sortStrings(a.Name, b.Name) : validParams;
}

export function sortActive(a: Stage, b: Stage): number {
    return checkNotNullParams(a.Active, b.Active);
}
