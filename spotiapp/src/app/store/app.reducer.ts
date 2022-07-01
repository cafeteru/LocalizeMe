import { ActionReducerMap } from '@ngrx/store';
import { getIsLoading, getUserReducer, IsLoadingReducer, IsoCodeReducer } from './reducers/iso-code.reducer';

export interface AppState {
    isoCodeReducer: IsoCodeReducer;
    isLoadingReducer: IsLoadingReducer;
}

export function getAppReducers(): ActionReducerMap<AppState> {
    return {
        isoCodeReducer: getUserReducer(),
        isLoadingReducer: getIsLoading()
    };
}
