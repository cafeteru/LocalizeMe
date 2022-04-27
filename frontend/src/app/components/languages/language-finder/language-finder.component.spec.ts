import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LanguageFinderComponent } from './language-finder.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('LanguageFinderComponent', () => {
    let component: LanguageFinderComponent;
    let fixture: ComponentFixture<LanguageFinderComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LanguageFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LanguageFinderComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
