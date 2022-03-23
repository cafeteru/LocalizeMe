import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseComponent } from '../../../core/base/base.component';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { StageRequest, StageService } from '../../../core/services/stage.service';

@Component({
    selector: 'app-create-stage',
    templateUrl: './create-stage.component.html',
    styleUrls: ['./create-stage.component.scss'],
})
export class CreateStageComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<CreateStageComponent>,
        private stageService: StageService,
        private message: NzMessageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            name: new FormControl('', Validators.required),
        });
    }

    createMessage(type: string, message: string): void {
        this.message.create(type, message);
    }

    close(): void {
        this.matDialogRef.close();
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const stageRequest: StageRequest = {
                Name: this.formGroup.controls['name'].value,
            };
            const subscription = this.stageService.create(stageRequest).subscribe({
                next: () => {
                    this.isLoading = false;
                    this.close();
                    this.createMessage('success', 'Successfully created.');
                },
                error: () => {
                    this.isLoading = false;
                    this.createMessage('error', 'Create not complete. Check the fields.');
                },
            });
            this.subscriptions.push(subscription);
        }
    }
}
