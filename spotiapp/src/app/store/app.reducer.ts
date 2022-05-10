import { ActionReducerMap } from '@ngrx/store';
import { getUserReducer, IsoCodeReducer } from './reducers/iso-code.reducer';

export interface AppState {
    isoCodeReducer: IsoCodeReducer;
}

export function getAppReducers(): ActionReducerMap<AppState> {
    return {
        isoCodeReducer: getUserReducer(),
    };
}
