import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupFinderComponent } from './group-finder.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('GroupFinderComponent', () => {
    let component: GroupFinderComponent;
    let fixture: ComponentFixture<GroupFinderComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [GroupFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(GroupFinderComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
