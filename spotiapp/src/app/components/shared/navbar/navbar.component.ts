import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../base.component';
import { AppState } from '../../../store/app.reducer';
import { Store } from '@ngrx/store';
import * as isoCodeActions from '../../../store/actions/iso-code.actions';
import { initialState } from '../../../store/reducers/iso-code.reducer';

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css'],
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NavbarComponent extends BaseComponent implements OnInit {
    isoCode = initialState.isoCode;

    constructor(private store: Store<AppState>) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.subscriptions$.push(
            this.store.select('isoCodeReducer').subscribe((isoCodeReducer) => {
                this.isoCode = isoCodeReducer.isoCode;
            })
        );
    }

    get selectedLanguage(): string {
        return 'flag-icon flag-icon-' + this.isoCode;
    }

    useLanguage(isoCode: string): void {
        this.isoCode = isoCode;
        this.store.dispatch(
            isoCodeActions.loadIsoCode({
                isoCode,
            })
        );
    }
}
