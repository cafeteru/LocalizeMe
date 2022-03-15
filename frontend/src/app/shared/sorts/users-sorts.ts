import { checkNotNullParams, sortStrings } from './sort-columns';
import { User } from '../../types/user';

export function sortEmail(a: User, b: User): number {
    const validParams = checkNotNullParams(a.Email, b.Email);
    return validParams === 0 ? sortStrings(a.Email, b.Email) : validParams;
}

export function sortIsAdmin(a: User, b: User): number {
    return checkNotNullParams(a.IsAdmin, b.IsAdmin);
}

export function sortIsActive(a: User, b: User): number {
    return checkNotNullParams(a.IsActive, b.IsActive);
}
