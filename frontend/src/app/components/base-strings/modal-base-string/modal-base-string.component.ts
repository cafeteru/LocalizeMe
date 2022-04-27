import { Component, Inject, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BaseComponent } from '../../../core/base/base.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseString } from '../../../types/base-string';
import { BaseStringService } from '../../../core/services/base-string.service';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { Observable, of } from 'rxjs';
import { Language } from '../../../types/language';
import { Group } from '../../../types/group';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { User } from '../../../types/user';

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
        this.formGroup = new FormGroup({
            identifier: new FormControl(this.baseString.identifier, Validators.required),
            active: new FormControl(this.baseString.active, Validators.required),
            group: new FormControl(this.baseString.group, Validators.required),
            sourceLanguage: new FormControl(this.baseString.sourceLanguage, Validators.required),
        });
        const subscription$ = this.store
            .select('userInfo')
            .subscribe((userReducer) => (this.author = userReducer.user));
        this.subscriptions$.push(subscription$);
    }

    get titleModal(): string {
        return this.baseString.id ? 'Update string' : 'Create string';
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
                next: (data) => {
                    this.isLoading = false;
                    this.close(data);
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
            identifier: this.formGroup.controls['identifier'].value,
            active: true,
            group: this.formGroup.controls['group'].value,
            sourceLanguage: this.formGroup.controls['sourceLanguage'].value,
            author: this.author,
            id: undefined,
            translations: [],
        };
        return this.baseStringService.create(baseString);
    }

    private update(): Observable<BaseString> {
        return of(undefined);
        // this.baseString = {
        //     ...this.baseString,
        //     description: this.formGroup.controls['description'].value,
        //     isoCode: this.formGroup.controls['isoCode'].value,
        //     active: this.formGroup.controls['active'].value,
        // };
        // return this.baseStringService.update(this.language);
    }
}
