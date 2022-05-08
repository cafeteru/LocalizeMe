import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../../core/base/base.component';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Group } from '../../../../types/group';
import { Stage } from '../../../../types/stage';
import { MatDialogRef } from '@angular/material/dialog';
import { NzUploadChangeParam } from 'ng-zorro-antd/upload';
import { FormGroupUtil } from '../../../../shared/utils/form-group-util';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseString } from '../../../../types/base-string';
import { XliffService } from '../../../../core/services/xliff.service';

@Component({
    selector: 'app-read-xliff',
    templateUrl: './read-xliff.component.html',
    styleUrls: ['./read-xliff.component.scss'],
})
export class ReadXliffComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<ReadXliffComponent>,
        private xliffService: XliffService,
        private nzMessageService: NzMessageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            group: new FormControl(undefined, Validators.required),
            stage: new FormControl(undefined, Validators.required),
            xliff: new FormControl(undefined, Validators.required),
        });
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    setGroup(group: Group) {
        this.formGroup.controls['group'].setValue(group);
    }

    setStage(stage: Stage) {
        this.formGroup.controls['stage'].setValue(stage);
    }

    showStageError(): boolean {
        return this.formGroup.controls['stage'].valid;
    }

    close(baseStrings?: BaseString[]): void {
        this.matDialogRef.close(baseStrings);
    }

    readFile({ file }: NzUploadChangeParam) {
        if (file.status !== 'uploading') {
            let reader = new FileReader();
            reader.onload = (e) => this.formGroup.controls['xliff'].setValue(e.target.result);
            reader.readAsText(file.originFileObj);
        }
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const stage = this.formGroup.controls['stage'].value as Stage;
            const group = this.formGroup.controls['group'].value as Group;
            const xliff = this.formGroup.controls['xliff'].value as string;
            const subscription$ = this.xliffService.read(stage, group, xliff).subscribe({
                next: (basesStrings) => {
                    this.close(basesStrings);
                    this.createMessage('success', 'Successfully loaded strings');
                },
                error: () => {
                    this.isLoading = false;
                    this.createMessage('error', 'Read not complete. Check the fields.');
                },
                complete: () => (this.isLoading = false),
            });
            this.subscriptions$.push(subscription$);
        }
    }
}
