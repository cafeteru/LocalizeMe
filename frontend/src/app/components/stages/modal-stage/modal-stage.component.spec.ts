import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalStageComponent } from './modal-stage.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CoreModule } from '../../../core/core.module';
import { SharedModule } from '../../../shared/shared.module';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { matDialogRefMock } from '../../../core/mocks/mock-tests';
import { createMockStage } from '../../../types/stage';

describe('CreateStageComponent', () => {
    let component: ModalStageComponent;
    let fixture: ComponentFixture<ModalStageComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [ModalStageComponent],
            imports: [HttpClientTestingModule, CoreModule, SharedModule],
            providers: [
                {
                    provide: MatDialogRef,
                    useValue: matDialogRefMock,
                },
                {
                    provide: MAT_DIALOG_DATA,
                    useValue: createMockStage(),
                },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(ModalStageComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
