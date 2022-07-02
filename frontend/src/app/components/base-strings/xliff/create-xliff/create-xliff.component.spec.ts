import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateXliffComponent } from './create-xliff.component';
import { SharedModule } from '../../../../shared/shared.module';
import { CoreModule } from '../../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import {
    CreateXliffListBaseStringsComponent
} from '../create-xliff-list-base-strings/create-xliff-list-base-strings.component';
import { LanguageFinderComponent } from '../../../languages/language-finder/language-finder.component';
import { MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../../core/mocks/mock-tests';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { StageFinderComponent } from '../../../stages/stage-finder/stage-finder.component';

describe('CreateXliffComponent', () => {
    let component: CreateXliffComponent;
    let fixture: ComponentFixture<CreateXliffComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [CreateXliffComponent, CreateXliffListBaseStringsComponent, LanguageFinderComponent, StageFinderComponent],
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
        fixture = TestBed.createComponent(CreateXliffComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
