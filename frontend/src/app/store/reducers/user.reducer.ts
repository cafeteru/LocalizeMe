import { ActionReducer, createReducer, on } from '@ngrx/store';

import * as UserActions from '../actions/user.actions';

export interface UserReducer {
    Email: string;
    Exp: number;
    IsActive: boolean;
    IsAdmin: boolean;
    Authorization: string;
}

export const initialState: UserReducer = {
    Email: '',
    Exp: 0,
    IsActive: false,
    IsAdmin: false,
    Authorization: '',
};

export function loadUser(state: UserReducer, userReducer: UserReducer): UserReducer {
    return {
        ...state,
        Email: userReducer.Email,
        Exp: userReducer.Exp,
        IsActive: userReducer.IsActive,
        IsAdmin: userReducer.IsAdmin,
        Authorization: userReducer.Authorization,
    };
}

export function clearUser(state: UserReducer): UserReducer {
    return {
        ...state,
        Email: initialState.Email,
        Exp: initialState.Exp,
        IsActive: initialState.IsActive,
        IsAdmin: initialState.IsAdmin,
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
