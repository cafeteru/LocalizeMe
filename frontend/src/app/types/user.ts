export interface UserDto {
    email: string;
    password: string;
}

export interface User {
    id: string;
    email: string;
    password: string;
    admin: boolean;
    active: boolean;
}

export function createMockUser(): User {
    return {
        active: true,
        password: 'Password',
        admin: false,
        email: 'Email',
        id: 'ID',
    };
}
