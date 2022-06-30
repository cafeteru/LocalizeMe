import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../base.component';
import { AppState } from '../../../store/app.reducer';
import { Store } from '@ngrx/store';
import * as isoCodeActions from '../../../store/actions/iso-code.actions';
import { initialState } from '../../../store/reducers/iso-code.reducer';
import { LocalizeMeService } from '../../../services/localize-me.service';

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css'],
})
export class NavbarComponent extends BaseComponent implements OnInit {
    isoCode = initialState.isoCode;

    constructor(private localizeMeService: LocalizeMeService, private store: Store<AppState>) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        const subscription1 = this.localizeMeService.login().subscribe();
        this.subscriptions$.push(subscription1);
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
