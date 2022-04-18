import { createMockUser, User } from './user';

export interface PermissionDto {
    email: string;
    canWriteGroup: boolean;
}

export interface Permission {
    id: string;
    user: User;
    canWriteGroup: boolean;
}

export function createMockPermission(): Permission {
    return {
        id: '1',
        user: createMockUser(),
        canWriteGroup: true,
    };
}
