import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateXliffListBaseStringsComponent } from './create-xliff-list-base-strings.component';
import { SharedModule } from '../../../../shared/shared.module';
import { CoreModule } from '../../../../core/core.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('CreateXliffListBaseStringsComponent', () => {
    let component: CreateXliffListBaseStringsComponent;
    let fixture: ComponentFixture<CreateXliffListBaseStringsComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [CreateXliffListBaseStringsComponent],
            imports: [SharedModule, CoreModule, HttpClientTestingModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(CreateXliffListBaseStringsComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
