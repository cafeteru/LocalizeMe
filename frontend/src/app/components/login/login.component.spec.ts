import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AppState } from '../../store/app.reducer';
import * as user from '../../store/reducers/user.reducer';
import { provideMockStore } from '@ngrx/store/testing';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';

describe('LoginComponent', () => {
    let component: LoginComponent;
    let fixture: ComponentFixture<LoginComponent>;
    const appState: AppState = {
        user: user.initialState,
    };

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LoginComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [provideMockStore({ initialState: appState })],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LoginComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
