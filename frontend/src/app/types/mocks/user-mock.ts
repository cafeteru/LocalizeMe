import { User } from '../user';

export function createUserMock(): User {
    return {
        IsActive: true,
        Password: 'Password',
        IsAdmin: false,
        Email: 'Email',
        ID: 'ID',
    };
}
