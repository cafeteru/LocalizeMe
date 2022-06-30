import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../../services/spotify.service';
import { initialState } from '../../store/reducers/iso-code.reducer';
import { BaseComponent } from '../base.component';

@Component({
    selector: 'app-search',
    templateUrl: './search.component.html',
    styles: [],
})
export class SearchComponent extends BaseComponent implements OnInit {
    artists: any[] = [];
    loading: boolean;
    isoCode = initialState.isoCode;

    constructor(private spotify: SpotifyService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
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
