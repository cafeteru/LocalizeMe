import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BaseStringListComponent } from './base-string-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ModalBaseStringComponent } from '../modal-base-string/modal-base-string.component';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { createMockAppState } from '../../../store/mocks/create-mock-app-state';

describe('BaseStringListComponent', () => {
    let component: BaseStringListComponent;
    let fixture: ComponentFixture<BaseStringListComponent>;
    let store: MockStore;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [BaseStringListComponent, ModalBaseStringComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
            providers: [provideMockStore({ initialState: createMockAppState() })],
        }).compileComponents();
        store = TestBed.inject(MockStore);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(BaseStringListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
