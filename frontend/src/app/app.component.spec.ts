import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { provideMockStore } from '@ngrx/store/testing';
import { CoreModule } from './core/core.module';
import { SharedModule } from './shared/shared.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LoginComponent } from './components/users/login/login.component';
import { createMockAppState } from './store/mocks/create-mock-app-state';

describe('AppComponent', () => {
    const appState = createMockAppState();

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [RouterTestingModule, CoreModule, SharedModule, HttpClientTestingModule],
            declarations: [AppComponent, LoginComponent],
            providers: [provideMockStore({ initialState: appState })],
        }).compileComponents();
    });

    it('should create the app', () => {
        const fixture = TestBed.createComponent(AppComponent);
        const app = fixture.componentInstance;
        expect(app).toBeTruthy();
    });
});
