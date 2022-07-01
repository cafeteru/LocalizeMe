import { createAction, props } from '@ngrx/store';
import { IsLoadingReducer, IsoCodeReducer } from '../reducers/iso-code.reducer';

export const loadIsoCode = createAction('[UI Component] loadIsoCode', props<IsoCodeReducer>());
export const loadIsLoading = createAction('[UI Component] loadIsLoading', props<IsLoadingReducer>());
export const clearIsoCode = createAction('[UI Component] clearIsoCode');
