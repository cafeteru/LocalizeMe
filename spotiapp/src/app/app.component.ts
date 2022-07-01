import { Component, OnInit } from '@angular/core';
import { BaseComponent } from './components/base.component';
import { Store } from '@ngrx/store';
import { AppState } from './store/app.reducer';
import { LocalizeMeService } from './services/localize-me.service';
import * as isoCodeActions from './store/actions/iso-code.actions';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent extends BaseComponent implements OnInit {
    title = 'app';
    isLoading: boolean;

    constructor(private localizeMeService: LocalizeMeService, private store: Store<AppState>) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.isLoading = true;
        this.store.dispatch(isoCodeActions.loadIsLoading({
            isLoading: this.isLoading
        }));
        localStorage.setItem('isoCode', 'esp');
        const subscription = this.store.select('isoCodeReducer').subscribe({
            next: (isoCodeReducer) => {
                localStorage.setItem('isoCode', isoCodeReducer.isoCode);
            }
        });
        const subscription1 = this.localizeMeService.login().subscribe({
            next: () => this.extracted(),
            error: () => this.extracted()
        });
        this.subscriptions$.push(subscription);
        this.subscriptions$.push(subscription1);
    }

    private extracted(): void {
        this.isLoading = false;
        this.store.dispatch(isoCodeActions.loadIsLoading({
            isLoading: this.isLoading
        }));
    }
}
