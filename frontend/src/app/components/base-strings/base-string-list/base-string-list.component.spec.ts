import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BaseStringListComponent } from './base-string-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ModalBaseStringComponent } from '../modal-base-string/modal-base-string.component';

describe('ModalBaseStringComponent', () => {
    let component: BaseStringListComponent;
    let fixture: ComponentFixture<BaseStringListComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [BaseStringListComponent, ModalBaseStringComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
        }).compileComponents();
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
