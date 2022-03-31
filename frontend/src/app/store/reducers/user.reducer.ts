import { ActionReducer, createReducer, on } from '@ngrx/store';

import * as UserActions from '../actions/user.actions';

export interface UserReducer {
    id: string;
    email: string;
    exp: number;
    active: boolean;
    admin: boolean;
    authorization: string;
}

export const initialState: UserReducer = {
    id: '',
    email: '',
    exp: 0,
    active: false,
    admin: false,
    authorization: '',
};

export function loadUser(state: UserReducer, userReducer: UserReducer): UserReducer {
    return {
        ...state,
        id: userReducer.id,
        email: userReducer.email,
        exp: userReducer.exp,
        active: userReducer.active,
        admin: userReducer.admin,
        authorization: userReducer.authorization,
    };
}

export function clearUser(state: UserReducer): UserReducer {
    return {
        ...state,
        id: initialState.id,
        email: initialState.email,
        exp: initialState.exp,
        active: initialState.active,
        admin: initialState.admin,
        authorization: initialState.authorization,
    };
}

export function getUserReducer(): ActionReducer<UserReducer> {
    return createReducer<UserReducer>(
        initialState,
        on(UserActions.loadUser, loadUser),
        on(UserActions.clearUser, clearUser)
    );
}
