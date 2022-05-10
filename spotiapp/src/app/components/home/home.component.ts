import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { BaseComponent } from '../base.component';
import { LocalizeMeService } from '../../services/localize-me.service';

@Component({
    selector: 'app-home',
    templateUrl: './home.component.html',
    styles: [],
})
export class HomeComponent extends BaseComponent implements OnInit {
    nuevasCanciones: any[] = [];
    loading: boolean;

    error: boolean;
    mensajeError: string;

    constructor(private spotifyService: SpotifyService, private localizeMeService: LocalizeMeService) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loading = true;
        this.error = false;
        this.getNewReleases();
    }

    private getNewReleases(): void {
        const subscription$ = this.spotifyService.getNewReleases().subscribe(
            (data: any) => {
                this.nuevasCanciones = data;
                this.loading = false;
            },
            (errorServicio) => {
                this.loading = false;
                this.error = true;
                this.mensajeError = errorServicio.error.error.message;
            }
        );
        this.subscriptions$.push(subscription$);
    }
}
