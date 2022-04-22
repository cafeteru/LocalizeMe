import { checkNotNullParams, sortStrings } from './sort-columns';
import { Group } from '../../types/group';

export function sortGroupByName(a: Group, b: Group): number {
    const validParams = checkNotNullParams(a.name, b.name);
    return validParams === 0 ? sortStrings(a.name, b.name) : validParams;
}

export function sortGroupByOwnerEmail(a: Group, b: Group): number {
    const validParams = checkNotNullParams(a.owner, b.owner);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.owner.email, b.owner.email);
    return validNames === 0 ? sortStrings(a.owner.email, b.owner.email) : validNames;
}

export function sortGroupByPublic(a: Group, b: Group): number {
    return checkNotNullParams(a.public, b.public);
}

export function sortGroupByActive(a: Group, b: Group): number {
    return checkNotNullParams(a.active, b.active);
}
