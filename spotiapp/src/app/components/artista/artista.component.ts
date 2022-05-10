import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SpotifyService } from '../../services/spotify.service';
import { initialState } from '../../store/reducers/iso-code.reducer';
import { BaseComponent } from '../base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';

@Component({
    selector: 'app-artista',
    templateUrl: './artista.component.html',
    styles: [],
})
export class ArtistaComponent extends BaseComponent implements OnInit {
    artista: any = {};
    topTracks: any[] = [];
    isoCode = initialState.isoCode;
    loadingArtist: boolean;

    constructor(private router: ActivatedRoute, private spotify: SpotifyService, private store: Store<AppState>) {
        super();
        this.loadingArtist = true;

        this.router.params.subscribe((params) => {
            this.getArtista(params['id']);
            this.getTopTracks(params['id']);
        });
    }

    ngOnInit() {
        super.ngOnInit();
        this.subscriptions$.push(
            this.store.select('isoCodeReducer').subscribe((isoCodeReducer) => {
                this.isoCode = isoCodeReducer.isoCode;
            })
        );
    }

    getArtista(id: string) {
        this.loadingArtist = true;

        this.spotify.getArtista(id).subscribe((artista) => {
            console.log(artista);
            this.artista = artista;

            this.loadingArtist = false;
        });
    }

    getTopTracks(id: string) {
        this.spotify.getTopTracks(id).subscribe((topTracks) => {
            console.log(topTracks);
            this.topTracks = topTracks;
        });
    }
}
