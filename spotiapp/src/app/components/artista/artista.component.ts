import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SpotifyService } from '../../services/spotify.service';
import { BaseComponent } from '../base.component';

@Component({
    selector: 'app-artista',
    templateUrl: './artista.component.html',
    styles: [],
})
export class ArtistaComponent extends BaseComponent implements OnInit {
    artista: any = {};
    topTracks: any[] = [];
    loadingArtist: boolean;

    constructor(private router: ActivatedRoute, private spotify: SpotifyService) {
        super();
        this.loadingArtist = true;

        this.router.params.subscribe((params) => {
            this.getArtista(params['id']);
            this.getTopTracks(params['id']);
        });
    }

    ngOnInit() {
        super.ngOnInit();
    }

    getArtista(id: string) {
        this.loadingArtist = true;

        this.spotify.getArtist(id).subscribe((artista) => {
            this.artista = artista;
            this.loadingArtist = false;
        });
    }

    getTopTracks(id: string) {
        this.spotify.getTopTracks(id).subscribe((topTracks) => {
            this.topTracks = topTracks;
        });
    }
}
