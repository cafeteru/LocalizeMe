import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserListComponent } from './user-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AppState } from '../../../store/app.reducer';
import { provideMockStore } from '@ngrx/store/testing';

describe('UserListComponent', () => {
    let component: UserListComponent;
    let fixture: ComponentFixture<UserListComponent>;
    let appState: AppState;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [UserListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [provideMockStore({ initialState: appState })],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(UserListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
