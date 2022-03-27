export interface User {
    ID: string;
    Email: string;
    Password: string;
    Admin: boolean;
    Active: boolean;
}

export function createMockUser(): User {
    return {
        Active: true,
        Password: 'Password',
        Admin: false,
        Email: 'Email',
        ID: 'ID',
    };
}
