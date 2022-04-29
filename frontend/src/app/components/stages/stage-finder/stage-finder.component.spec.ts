import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StageFinderComponent } from './stage-finder.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('StageFinderComponent', () => {
    let component: StageFinderComponent;
    let fixture: ComponentFixture<StageFinderComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [StageFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(StageFinderComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
