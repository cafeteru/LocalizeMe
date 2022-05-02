import { checkNotNullParams, sortStrings } from './sort-columns';
import { Permission } from '../../types/permission';

export function sortPermissionsByUserEmail(a: Permission, b: Permission): number {
    const validParams = checkNotNullParams(a.user.email, b.user.email);
    return validParams === 0 ? sortStrings(a.user.email, b.user.email) : validParams;
}

export function sortPermissionsByCanWrite(a: Permission, b: Permission): number {
    return checkNotNullParams(a.canWrite, b.canWrite);
}
