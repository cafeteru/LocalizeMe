import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenuComponent } from './menu.component';
import { RouterTestingModule } from '@angular/router/testing';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { createMockAppState } from '../../store/mocks/create-mock-app-state';

describe('MenuComponent', () => {
    let component: MenuComponent;
    let fixture: ComponentFixture<MenuComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [RouterTestingModule, CoreModule, SharedModule, HttpClientTestingModule],
            declarations: [MenuComponent],
            providers: [provideMockStore({ initialState: createMockAppState() })],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(MenuComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
