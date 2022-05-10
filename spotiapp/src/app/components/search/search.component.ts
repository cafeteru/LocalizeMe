import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { initialState } from '../../store/reducers/iso-code.reducer';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import { BaseComponent } from '../base.component';

@Component({
    selector: 'app-search',
    templateUrl: './search.component.html',
    styles: [],
})
export class SearchComponent extends BaseComponent implements OnInit {
    artistas: any[] = [];
    loading: boolean;
    isoCode = initialState.isoCode;

    constructor(private spotify: SpotifyService, private store: Store<AppState>) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.subscriptions$.push(
            this.store.select('isoCodeReducer').subscribe((isoCodeReducer) => {
                this.isoCode = isoCodeReducer.isoCode;
            })
        );
    }

    buscar(termino: string) {
        if (termino && termino !== '') {
            this.loading = true;
            this.spotify.getArtistas(termino).subscribe((data: any) => {
                this.artistas = data;
                this.loading = false;
            });
        } else {
            this.artistas = [];
            this.loading = false;
        }
    }
}
