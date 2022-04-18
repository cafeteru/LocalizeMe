import { checkNotNullParams, sortStrings } from './sort-columns';
import { User } from '../../types/user';

export function sortUsersByEmail(a: User, b: User): number {
    const validParams = checkNotNullParams(a.email, b.email);
    return validParams === 0 ? sortStrings(a.email, b.email) : validParams;
}

export function sortUsersByIsAdmin(a: User, b: User): number {
    return checkNotNullParams(a.admin, b.admin);
}

export function sortUsersByIsActive(a: User, b: User): number {
    return checkNotNullParams(a.active, b.active);
}
