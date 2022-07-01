import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { initialState } from '../../store/reducers/iso-code.reducer';
import { BaseComponent } from '../base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Component({
    selector: 'app-search',
    templateUrl: './search.component.html',
    styles: [],
})
export class SearchComponent extends BaseComponent implements OnInit {
    artists: any[] = [];
    loading: boolean;
    isoCode = initialState.isoCode;
    loadingLocalizeMe: boolean;

    constructor(private spotify: SpotifyService, private store: Store<AppState>) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        const subscription2 = this.store.select('isLoadingReducer').subscribe({
            next: (isLoadingReducer) => {
                this.loadingLocalizeMe = isLoadingReducer.isLoading;
            }
        });
        this.subscriptions$.push(subscription2);
    }

    search(value: string) {
        if (value && value !== '') {
            this.loading = true;
            this.spotify.getArtists(value).subscribe((data: any) => {
                this.artists = data;
                this.loading = false;
            });
        } else {
            this.artists = [];
            this.loading = false;
        }
    }
}
