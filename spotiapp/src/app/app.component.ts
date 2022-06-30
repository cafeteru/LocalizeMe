import { Component, OnInit } from '@angular/core';
import { BaseComponent } from './components/base.component';
import { Store } from '@ngrx/store';
import { AppState } from './store/app.reducer';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent extends BaseComponent implements OnInit {
    title = 'app';

    constructor(private store: Store<AppState>) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        localStorage.setItem('isoCode', 'esp');
        const subscription = this.store.select('isoCodeReducer').subscribe({
            next: (isoCodeReducer) => {
                localStorage.setItem('isoCode', isoCodeReducer.isoCode);
            },
        });
        this.subscriptions$.push(subscription);
    }
}
