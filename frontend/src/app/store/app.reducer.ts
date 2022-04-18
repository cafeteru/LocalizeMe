import { ActionReducerMap } from '@ngrx/store';
import { getUserReducer, UserReducer } from './reducers/user.reducer';

export interface AppState {
    userInfo: UserReducer;
}

export function getAppReducers(): ActionReducerMap<AppState> {
    return {
        userInfo: getUserReducer(),
    };
}
