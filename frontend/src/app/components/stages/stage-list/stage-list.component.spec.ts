import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StageListComponent } from './stage-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('StageListComponent', () => {
    let component: StageListComponent;
    let fixture: ComponentFixture<StageListComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [StageListComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(StageListComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
