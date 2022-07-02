import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseComponent } from '../../../../core/base/base.component';
import { XliffService } from '../../../../core/services/xliff.service';
import { Language } from '../../../../types/language';
import { FormGroupUtil } from '../../../../shared/utils/form-group-util';
import { XliffDto } from '../../../../types/xliff';
import * as FileSaver from 'file-saver';
import { BaseString } from '../../../../types/base-string';
import beautify from 'xml-beautifier';
import { Stage } from '../../../../types/stage';

@Component({
    selector: 'app-create-xliff',
    templateUrl: './create-xliff.component.html',
    styleUrls: ['./create-xliff.component.scss'],
})
export class CreateXliffComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;
    sourceLanguage: Language;
    targetLanguage: Language;

    constructor(
        private matDialogRef: MatDialogRef<CreateXliffComponent>,
        private xliffService: XliffService,
        private nzMessageService: NzMessageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            baseStringIds: new FormControl(undefined),
            sourceLanguageId: new FormControl(undefined, Validators.required),
            stage: new FormControl(undefined, Validators.required),
            targetLanguageId: new FormControl(undefined, Validators.required),
        });
    }

    setSourceLanguage(language: Language) {
        this.sourceLanguage = language;
        this.formGroup.controls['sourceLanguageId'].setValue(language?.id);
    }

    setTargetLanguage(language: Language) {
        this.targetLanguage = language;
        this.formGroup.controls['targetLanguageId'].setValue(language?.id);
    }

    showLanguageError(name: string): boolean {
        return this.formGroup.controls[name].valid;
    }

    setBaseStringIds(baseStrings: BaseString[]) {
        this.formGroup.controls['baseStringIds'].setValue(baseStrings.map((baseString) => baseString.id));
    }

    setStage(stage: Stage) {
        this.formGroup.controls['stage'].setValue(stage);
    }

    showStageError(): boolean {
        return this.formGroup.controls['stage'].valid;
    }

    close(): void {
        this.matDialogRef.close();
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const xliffDto: XliffDto = {
                baseStringIds: this.formGroup.controls['baseStringIds'].value,
                sourceLanguageId: this.formGroup.controls['sourceLanguageId'].value,
                stage: this.formGroup.controls['stage'].value,
                targetLanguageId: this.formGroup.controls['targetLanguageId'].value,
            };
            const subscription$ = this.xliffService.createXliff(xliffDto).subscribe({
                next: (xmlContent) => {
                    const blob = new Blob([beautify(xmlContent)], { type: 'application/xml;charset=utf-8' });
                    FileSaver.saveAs(blob, 'xliff.xml');
                    this.close();
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
