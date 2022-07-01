import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SpotifyService } from '../../services/spotify.service';
import { BaseComponent } from '../base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Component({
    selector: 'app-artista',
    templateUrl: './artista.component.html',
    styles: [],
})
export class ArtistaComponent extends BaseComponent implements OnInit {
    artist: any = {};
    topTracks: any[] = [];
    loadingArtist: boolean;
    loadingLocalizeMe: boolean;

    constructor(private router: ActivatedRoute, private spotify: SpotifyService, private store: Store<AppState>) {
        super();
        this.loadingArtist = true;

        this.router.params.subscribe((params) => {
            this.getArtista(params['id']);
            this.getTopTracks(params['id']);
        });
        const subscription2 = this.store.select('isLoadingReducer').subscribe({
            next: (isLoadingReducer) => {
                this.loadingLocalizeMe = isLoadingReducer.isLoading;
            }
        });
        this.subscriptions$.push(subscription2);
    }

    ngOnInit() {
        super.ngOnInit();
    }

    getArtista(id: string) {
        this.loadingArtist = true;

        this.spotify.getArtist(id).subscribe((artista) => {
            this.artist = artista;
            this.loadingArtist = false;
        });
    }

    getTopTracks(id: string) {
        this.spotify.getTopTracks(id).subscribe((topTracks) => {
            this.topTracks = topTracks;
        });
    }
}
