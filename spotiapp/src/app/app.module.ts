import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { SearchComponent } from './components/search/search.component';
import { ArtistaComponent } from './components/artista/artista.component';
import { NavbarComponent } from './components/shared/navbar/navbar.component';

// Importar rutas
import { ROUTES } from './app.routes';

// Pipes
import { NoimagePipe } from './pipes/noimage.pipe';
import { DomseguroPipe } from './pipes/domseguro.pipe';

import { TarjetasComponent } from './components/tarjetas/tarjetas.component';
import { LoadingComponent } from './components/shared/loading/loading.component';
import { LocalizeMePipe } from './pipes/localize-me.pipe';
import { StoreModule } from '@ngrx/store';
import { getAppReducers } from './store/app.reducer';
import { environment } from '../environments/environment';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';

@NgModule({
    declarations: [
        AppComponent,
        HomeComponent,
        SearchComponent,
        ArtistaComponent,
        NavbarComponent,
        NoimagePipe,
        DomseguroPipe,
        TarjetasComponent,
        LoadingComponent,
        LocalizeMePipe,
    ],
    imports: [
        BrowserModule,
        HttpClientModule,
        RouterModule.forRoot(ROUTES, { useHash: true }),
        StoreModule.forRoot(getAppReducers()),
        StoreDevtoolsModule.instrument({
            maxAge: 25,
            logOnly: environment.production,
        }),
    ],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule {}
