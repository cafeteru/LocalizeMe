import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { StageRequest, StageService } from '../../../core/services/stage.service';
import { Stage } from '../../../types/stage';
import { Observable } from 'rxjs';
import { BaseComponent } from '../../../core/base/base.component';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';

@Component({
    selector: 'app-modal-stage',
    templateUrl: './modal-stage.component.html',
    styleUrls: ['./modal-stage.component.scss'],
})
export class ModalStageComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
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
        this.formGroup = new FormGroup({
            Name: new FormControl(this.stage.Name, Validators.required),
            Active: new FormControl(this.stage.Active, Validators.required),
        });
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
            const observable = this.stage.ID ? this.update() : this.create();
            const subscription$ = observable.subscribe({
                next: (data) => {
                    this.isLoading = false;
                    this.close(data);
                    const message = this.stage.ID ? 'Successfully updated stage' : 'Successfully created stage';
                    this.createMessage('success', message);
                },
                error: () => {
                    this.isLoading = false;
                    const message = this.stage.ID
                        ? 'Update not complete. Check the fields.'
                        : 'Create not complete. Check the fields.';
                    this.createMessage('error', message);
                },
            });
            this.subscriptions$.push(subscription$);
        }
    }

    private create(): Observable<Stage> {
        const stageRequest: StageRequest = {
            Name: this.formGroup.controls['Name'].value,
        };
        return this.stageService.create(stageRequest);
    }

    private update(): Observable<Stage> {
        this.stage = {
            ...this.stage,
            Name: this.formGroup.controls['Name'].value,
            Active: this.formGroup.controls['Active'].value,
        };
        return this.stageService.update(this.stage);
    }
}
