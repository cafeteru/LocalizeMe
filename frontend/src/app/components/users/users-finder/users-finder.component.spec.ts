import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UsersFinderComponent } from './users-finder.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { createAppStateMock } from '../../../store/mocks/create-app-state-mock';
import { AppState } from '../../../store/app.reducer';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';

describe('UsersFinderComponent', () => {
    let component: UsersFinderComponent;
    let fixture: ComponentFixture<UsersFinderComponent>;
    let appState: AppState;
    let store: MockStore;

    beforeEach(async () => {
        appState = createAppStateMock();
        await TestBed.configureTestingModule({
            declarations: [UsersFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [provideMockStore({ initialState: appState })],
        }).compileComponents();
        store = TestBed.inject(MockStore);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UsersFinderComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
