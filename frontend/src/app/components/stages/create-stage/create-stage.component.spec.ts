import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateStageComponent } from './create-stage.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CoreModule } from '../../../core/core.module';
import { SharedModule } from '../../../shared/shared.module';
import { MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';

describe('CreateStageComponent', () => {
    let component: CreateStageComponent;
    let fixture: ComponentFixture<CreateStageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [CreateStageComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(CreateStageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
