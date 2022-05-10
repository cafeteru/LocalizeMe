import { ActionReducer, createReducer, on } from '@ngrx/store';

import * as IsoCodeActions from '../actions/iso-code.actions';

export interface IsoCodeReducer {
    isoCode: string;
}

export const initialState: IsoCodeReducer = {
    isoCode: 'esp',
};

export function loadIsoCode(state: IsoCodeReducer, isoCodeReducer: IsoCodeReducer): IsoCodeReducer {
    return {
        ...state,
        isoCode: isoCodeReducer.isoCode,
    };
}

export function clearIsoCode(state: IsoCodeReducer): IsoCodeReducer {
    return {
        ...state,
        isoCode: initialState.isoCode,
    };
}

export function getUserReducer(): ActionReducer<IsoCodeReducer> {
    return createReducer<IsoCodeReducer>(
        initialState,
        on(IsoCodeActions.loadIsoCode, loadIsoCode),
        on(IsoCodeActions.clearIsoCode, clearIsoCode)
    );
}
