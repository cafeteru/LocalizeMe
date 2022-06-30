import { Component, Input, OnChanges, OnInit } from '@angular/core';
import { LocalizeMeService } from '../../../services/localize-me.service';
import { BaseComponent } from '../../base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { tap } from 'rxjs/operators';

@Component({
    selector: 'app-localize-me',
    templateUrl: './localize-me.component.html',
    styleUrls: ['./localize-me.component.css'],
})
export class LocalizeMeComponent extends BaseComponent implements OnInit, OnChanges {
    identifier = '';
    isoCode = localStorage['isoCode'];
    loading = false;
    value = '';

    constructor(private localizeMeService: LocalizeMeService, private store: Store<AppState>) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.store
            .select('isoCodeReducer')
            .pipe(
                tap((isoCodeReducer) => {
                    this.isoCode = isoCodeReducer.isoCode;
                }),
                tap(() => {
                    this.loadString();
                })
            )
            .subscribe();
        this.subscriptions$.push(subscription);
    }

    ngOnChanges(): void {
        this.loadString();
    }

    @Input()
    public set setIdentifier(identifier: string) {
        this.identifier = identifier;
        this.loadString();
    }

    loadString(): void {
        if (this.identifier) {
            const subscription = this.localizeMeService
                .findByIdentifierAndLanguage(this.identifier, this.isoCode)
                .pipe(tap(() => (this.loading = true)))
                .subscribe({
                    next: (res) => {
                        this.loading = false;
                        this.value = res;
                    },
                    error: () => {
                        this.loading = false;
                        this.value = this.identifier;
                    },
                });
            this.subscriptions$.push(subscription);
        }
    }
}
