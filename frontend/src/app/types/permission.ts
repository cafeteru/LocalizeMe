import { User } from './user';

export interface Permission {
    user: User;
    canWrite: boolean;
}
