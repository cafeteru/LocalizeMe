import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReadXliffComponent } from './read-xliff.component';
import { SharedModule } from '../../../../shared/shared.module';
import { CoreModule } from '../../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../../core/mocks/mock-tests';
import { GroupFinderComponent } from '../../../groups/group-finder/group-finder.component';
import { StageFinderComponent } from '../../../stages/stage-finder/stage-finder.component';

describe('ReadXliffComponent', () => {
    let component: ReadXliffComponent;
    let fixture: ComponentFixture<ReadXliffComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ReadXliffComponent, GroupFinderComponent, StageFinderComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule, BrowserAnimationsModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ReadXliffComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
