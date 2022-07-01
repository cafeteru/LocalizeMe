import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { BaseComponent } from '../base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html',
    styles: [],
})
export class HomeComponent extends BaseComponent implements OnInit {
    newSongs: any[] = [];
    loading: boolean;
    loadingLocalizeMe: boolean;

    error: boolean;
    messageError: string;

    constructor(private spotifyService: SpotifyService, private store: Store<AppState>) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loading = true;
        this.error = false;
        const subscription$ = this.spotifyService.getToken().subscribe({
            next: () => this.getNewReleases(),
        });
        this.subscriptions$.push(subscription$);
        const subscription2 = this.store.select('isLoadingReducer').subscribe({
            next: (isLoadingReducer) => {
                this.loadingLocalizeMe = isLoadingReducer.isLoading;
            }
        });
        this.subscriptions$.push(subscription2);
    }

    private getNewReleases(): void {
        const subscription$ = this.spotifyService.getNewReleases().subscribe({
            next: (data: any) => {
                this.newSongs = data;
                this.loading = false;
            },
            error: (errorService) => {
                this.loading = false;
                this.error = true;
                this.messageError = errorService.error.error.message;
            },
        });
        this.subscriptions$.push(subscription$);
    }
}
