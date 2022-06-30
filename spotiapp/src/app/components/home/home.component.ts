import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { BaseComponent } from '../base.component';

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

    constructor(private spotifyService: SpotifyService) {
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
    }

    private getNewReleases(): void {
        const subscription$ = this.spotifyService.getNewReleases().subscribe({
            next: (data: any) => {
                this.nuevasCanciones = data;
                this.loading = false;
            },
            error: (errorServicio) => {
                this.loading = false;
                this.error = true;
                this.mensajeError = errorServicio.error.error.message;
            },
        });
        this.subscriptions$.push(subscription$);
    }
}
