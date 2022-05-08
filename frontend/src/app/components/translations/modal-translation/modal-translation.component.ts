import { Component, Inject, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Translation } from '../../../types/translation';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { Language } from '../../../types/language';
import { Stage } from '../../../types/stage';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { User } from '../../../types/user';

@Component({
    selector: 'app-modal-translation',
    templateUrl: './modal-translation.component.html',
    styleUrls: ['./modal-translation.component.scss'],
})
export class ModalTranslationComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    private author: User;

    constructor(
        @Inject(MAT_DIALOG_DATA) public translation: Translation,
        private matDialogRef: MatDialogRef<ModalTranslationComponent>,
        private store: Store<AppState>
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            content: new FormControl(this.translation.content, Validators.required),
            version: new FormControl(this.translation.version, Validators.required),
            language: new FormControl(this.translation.language, Validators.required),
            stage: new FormControl(this.translation.stage, Validators.required),
            active: new FormControl(this.translation.active, Validators.required),
        });
        const subscription$ = this.store
            .select('userInfo')
            .subscribe((userReducer) => (this.author = userReducer.user));
        this.subscriptions$.push(subscription$);
    }

    get titleModal(): string {
        return this.translation.author ? 'Update translation' : 'Create translation';
    }

    get btnModal(): string {
        return this.translation.author ? 'Update' : 'Create';
    }

    setLanguage(language: Language) {
        this.formGroup.controls['language'].setValue(language);
    }

    showLanguageError(): boolean {
        return this.formGroup.controls['language'].valid;
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

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.translation = {
                ...this.translation,
                active: this.formGroup.controls['active'].value,
                author: this.author,
                language: this.formGroup.controls['language'].value,
                stage: this.formGroup.controls['stage'].value,
                content: this.formGroup.controls['content'].value,
                version: this.formGroup.controls['version'].value,
            };
            this.matDialogRef.close(this.translation);
        }
    }
}
