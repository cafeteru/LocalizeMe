import { createAction, props } from '@ngrx/store';
import { UserReducer } from '../reducers/user.reducer';

export const loadUser = createAction('[UI Component] loadUser', props<UserReducer>());
export const clearUser = createAction('[UI Component] clearUser');
