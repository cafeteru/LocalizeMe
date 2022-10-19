import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';
import { StageService } from '../../../core/services/stage.service';
import { Stage, StageDto } from '../../../types/stage';
import { Observable } from 'rxjs';
import { BaseComponent } from '../../../core/base/base.component';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';

@Component({
    selector: 'app-modal-stage',
    templateUrl: './modal-stage.component.html',
    styleUrls: ['./modal-stage.component.scss'],
})
export class ModalStageComponent extends BaseComponent implements OnInit {
    formGroup = new UntypedFormGroup({});
    isLoading = false;

    constructor(
        @Inject(MAT_DIALOG_DATA) public stage: Stage,
        private matDialogRef: MatDialogRef<ModalStageComponent>,
        private nzMessageService: NzMessageService,
        private stageService: StageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new UntypedFormGroup({
            name: new UntypedFormControl(this.stage.name, Validators.required),
            active: new UntypedFormControl(this.stage.active, Validators.required),
        });
    }

    get titleModal(): string {
        return this.stage.id ? 'Update stage' : 'Create stage';
    }

    get btnModal(): string {
        return this.stage.id ? 'Update' : 'Create';
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    close(stage?: Stage): void {
        this.matDialogRef.close(stage);
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const observable = this.stage.id ? this.update() : this.create();
            const subscription$ = observable.subscribe({
                next: (data) => {
                    this.isLoading = false;
                    this.close(data);
                    const message = this.stage.id ? 'Successfully updated stage' : 'Successfully created stage';
                    this.createMessage('success', message);
                },
                error: () => {
                    this.isLoading = false;
                    const message = this.stage.id
                        ? 'Update not complete. Check the fields.'
                        : 'Create not complete. Check the fields.';
                    this.createMessage('error', message);
                },
            });
            this.subscriptions$.push(subscription$);
        }
    }

    private create(): Observable<Stage> {
        const stageDto: StageDto = {
            name: this.formGroup.controls['name'].value,
        };
        return this.stageService.create(stageDto);
    }

    private update(): Observable<Stage> {
        this.stage = {
            ...this.stage,
            name: this.formGroup.controls['name'].value,
            active: this.formGroup.controls['active'].value,
        };
        return this.stageService.update(this.stage);
    }
}
