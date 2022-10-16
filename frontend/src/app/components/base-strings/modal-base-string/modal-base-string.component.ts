import { Component, Inject, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BaseComponent } from '../../../core/base/base.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseString } from '../../../types/base-string';
import { BaseStringService } from '../../../core/services/base-string.service';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { Observable } from 'rxjs';
import { Language } from '../../../types/language';
import { Group } from '../../../types/group';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { User } from '../../../types/user';
import { Translation } from '../../../types/translation';

@Component({
    selector: 'app-modal-base-string',
    templateUrl: './modal-base-string.component.html',
    styleUrls: ['./modal-base-string.component.scss'],
})
export class ModalBaseStringComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;
    languages: Language[] = [];
    groups: Group[] = [];
    author: User;

    constructor(
        @Inject(MAT_DIALOG_DATA) public baseString: BaseString,
        private matDialogRef: MatDialogRef<ModalBaseStringComponent>,
        private nzMessageService: NzMessageService,
        private baseStringService: BaseStringService,
        private store: Store<AppState>
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        const { translations, identifier, active, group, sourceLanguage } = this.baseString;
        this.formGroup = new FormGroup({
            identifier: new FormControl(identifier, Validators.required),
            active: new FormControl(active, Validators.required),
            group: new FormControl(group),
            sourceLanguage: new FormControl(sourceLanguage, Validators.required),
            translations: new FormControl(translations ? translations : []),
        });
        const subscription$ = this.store
            .select('userInfo')
            .subscribe((userReducer) => (this.author = userReducer.user));
        this.subscriptions$.push(subscription$);
    }

    get titleModal(): string {
        return this.baseString.id ? 'Update baseString' : 'Create baseString';
    }

    get btnModal(): string {
        return this.baseString.id ? 'Update' : 'Create';
    }

    setGroup(group: Group) {
        this.formGroup.controls['group'].setValue(group);
    }

    setLanguage(sourceLanguage: Language) {
        this.formGroup.controls['sourceLanguage'].setValue(sourceLanguage);
    }

    showLanguageError(): boolean {
        return this.formGroup.controls['sourceLanguage'].valid;
    }

    setTranslations($event: Translation[]) {
        this.formGroup.controls['translations'].setValue($event);
    }

    close(baseString?: BaseString): void {
        this.matDialogRef.close(baseString);
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const observable = this.baseString.id ? this.update() : this.create();
            const subscription$ = observable.subscribe({
                next: (baseString) => {
                    this.isLoading = false;
                    this.close(baseString);
                    const message = this.baseString.id ? 'Successfully updated string' : 'Successfully created string';
                    this.createMessage('success', message);
                },
                error: () => {
                    this.isLoading = false;
                    const message = this.baseString.id
                        ? 'Update not complete. Check the fields.'
                        : 'Create not complete. Check the fields.';
                    this.createMessage('error', message);
                },
            });
            this.subscriptions$.push(subscription$);
        }
    }

    private create(): Observable<BaseString> {
        const baseString: BaseString = {
            active: true,
            author: this.author,
            group: this.formGroup.controls['group'].value,
            id: undefined,
            identifier: this.formGroup.controls['identifier'].value,
            sourceLanguage: this.formGroup.controls['sourceLanguage'].value,
            translations: this.formGroup.controls['translations'].value,
        };
        return this.baseStringService.create(baseString);
    }

    private update(): Observable<BaseString> {
        const controls = this.formGroup.controls;
        this.baseString = {
            ...this.baseString,
            active: controls['active'].value,
            group: controls['group'].value,
            identifier: controls['identifier'].value,
            sourceLanguage: controls['sourceLanguage'].value,
            translations: controls['translations'].value,
        };
        return this.baseStringService.update(this.baseString);
    }
}
