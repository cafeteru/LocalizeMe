import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LanguageListComponent } from './language-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('LanguageListComponent', () => {
    let component: LanguageListComponent;
    let fixture: ComponentFixture<LanguageListComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LanguageListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LanguageListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
