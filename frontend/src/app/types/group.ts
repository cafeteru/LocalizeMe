import { createMockUser, User } from './user';
import { Permission, PermissionDto } from './permission';

export interface GroupDto {
    name: string;
    owner: User;
    permissions: PermissionDto[];
}

export interface Group {
    id: string;
    name: string;
    owner: User;
    permissions: Permission[];
    active: boolean;
}

export function createMockGroup(): Group {
    return {
        id: '1',
        name: 'group',
        owner: createMockUser(),
        active: true,
        permissions: [],
    };
}
