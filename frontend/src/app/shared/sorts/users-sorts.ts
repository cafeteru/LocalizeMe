import { checkNotNullParams, sortStrings } from './sort-columns';
import { User } from '../../types/user';

export function sortEmail(a: User, b: User): number {
    const validParams = checkNotNullParams(a.email, b.email);
    return validParams === 0 ? sortStrings(a.email, b.email) : validParams;
}

export function sortIsAdmin(a: User, b: User): number {
    return checkNotNullParams(a.admin, b.admin);
}

export function sortIsActive(a: User, b: User): number {
    return checkNotNullParams(a.active, b.active);
}
