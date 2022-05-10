import { createAction, props } from '@ngrx/store';
import { IsoCodeReducer } from '../reducers/iso-code.reducer';

export const loadIsoCode = createAction('[UI Component] loadIsoCode', props<IsoCodeReducer>());
export const clearIsoCode = createAction('[UI Component] clearIsoCode');
