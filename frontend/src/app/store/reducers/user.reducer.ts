import { ActionReducer, createReducer, on } from '@ngrx/store';

import * as UserActions from '../actions/user.actions';

export interface UserReducer {
    ID: string;
    Email: string;
    Exp: number;
    Active: boolean;
    Admin: boolean;
    Authorization: string;
}

export const initialState: UserReducer = {
    ID: '',
    Email: '',
    Exp: 0,
    Active: false,
    Admin: false,
    Authorization: '',
};

export function loadUser(state: UserReducer, userReducer: UserReducer): UserReducer {
    return {
        ...state,
        ID: userReducer.ID,
        Email: userReducer.Email,
        Exp: userReducer.Exp,
        Active: userReducer.Active,
        Admin: userReducer.Admin,
        Authorization: userReducer.Authorization,
    };
}

export function clearUser(state: UserReducer): UserReducer {
    return {
        ...state,
        ID: initialState.ID,
        Email: initialState.Email,
        Exp: initialState.Exp,
        Active: initialState.Active,
        Admin: initialState.Admin,
        Authorization: initialState.Authorization,
    };
}

export function getUserReducer(): ActionReducer<UserReducer> {
    return createReducer<UserReducer>(
        initialState,
        on(UserActions.loadUser, loadUser),
        on(UserActions.clearUser, clearUser)
    );
}
