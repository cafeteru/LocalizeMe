import { Component, Inject, OnInit } from '@angular/core';
import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseComponent } from '../../../core/base/base.component';
import { LanguageService } from '../../../core/services/language.service';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { Observable } from 'rxjs';
import { Language, LanguageDto } from '../../../types/language';

@Component({
    selector: 'app-modal-language',
    templateUrl: './modal-language.component.html',
    styleUrls: ['./modal-language.component.scss'],
})
export class ModalLanguageComponent extends BaseComponent implements OnInit {
    formGroup = new UntypedFormGroup({});
    isLoading = false;

    constructor(
        @Inject(MAT_DIALOG_DATA) public language: Language,
        private matDialogRef: MatDialogRef<ModalLanguageComponent>,
        private nzMessageService: NzMessageService,
        private languageService: LanguageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new UntypedFormGroup({
            isoCode: new UntypedFormControl(this.language.isoCode, Validators.required),
            description: new UntypedFormControl(this.language.description),
            active: new UntypedFormControl(this.language.active, Validators.required),
        });
    }

    get titleModal(): string {
        return this.language.id ? 'Update language' : 'Create language';
    }

    get btnModal(): string {
        return this.language.id ? 'Update' : 'Create';
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    close(language?: Language): void {
        this.matDialogRef.close(language);
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const observable = this.language.id ? this.update() : this.create();
            const subscription$ = observable.subscribe({
                next: (data) => {
                    this.isLoading = false;
                    this.close(data);
                    const message = this.language.id
                        ? 'Successfully updated language'
                        : 'Successfully created language';
                    this.createMessage('success', message);
                },
                error: () => {
                    this.isLoading = false;
                    const message = this.language.id
                        ? 'Update not complete. Check the fields.'
                        : 'Create not complete. Check the fields.';
                    this.createMessage('error', message);
                },
            });
            this.subscriptions$.push(subscription$);
        }
    }

    private create(): Observable<Language> {
        const languageDto: LanguageDto = {
            description: this.formGroup.controls['description'].value,
            isoCode: this.formGroup.controls['isoCode'].value,
        };
        return this.languageService.create(languageDto);
    }

    private update(): Observable<Language> {
        this.language = {
            ...this.language,
            description: this.formGroup.controls['description'].value,
            isoCode: this.formGroup.controls['isoCode'].value,
            active: this.formGroup.controls['active'].value,
        };
        return this.languageService.update(this.language);
    }
}
