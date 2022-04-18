import { ActionReducer, createReducer, on } from '@ngrx/store';

import * as UserActions from '../actions/user.actions';
import { User } from '../../types/user';

export interface UserReducer {
    authorization: string;
    exp: number;
    user: User;
}

export const initialState: UserReducer = {
    exp: 0,
    authorization: '',
    user: {
        id: '',
        admin: false,
        active: false,
        email: '',
        password: '',
    },
};

export function loadUser(state: UserReducer, userReducer: UserReducer): UserReducer {
    return {
        ...state,
        authorization: userReducer.authorization,
        exp: userReducer.exp,
        user: userReducer.user,
    };
}

export function clearUser(state: UserReducer): UserReducer {
    return {
        ...state,
        authorization: initialState.authorization,
        exp: initialState.exp,
        user: initialState.user,
    };
}

export function getUserReducer(): ActionReducer<UserReducer> {
    return createReducer<UserReducer>(
        initialState,
        on(UserActions.loadUser, loadUser),
        on(UserActions.clearUser, clearUser)
    );
}
