import { Component, Inject, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BaseComponent } from '../../../core/base/base.component';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { GroupService } from '../../../core/services/group.service';
import { User } from '../../../types/user';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { createMockGroup, Group, GroupDto } from '../../../types/group';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { Observable, of } from 'rxjs';

@Component({
    selector: 'app-modal-group',
    templateUrl: './modal-group.component.html',
    styleUrls: ['./modal-group.component.scss'],
})
export class ModalGroupComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;
    owner: User;

    constructor(
        @Inject(MAT_DIALOG_DATA) public group: Group,
        private matDialogRef: MatDialogRef<ModalGroupComponent>,
        private nzMessageService: NzMessageService,
        private groupService: GroupService,
        private store: Store<AppState>
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            name: new FormControl(this.group.name, Validators.required),
            active: new FormControl(this.group.active),
        });
        const subscription$ = this.store.select('userInfo').subscribe((userReducer) => (this.owner = userReducer.user));
        this.subscriptions$.push(subscription$);
    }

    get titleModal(): string {
        return this.group.id ? 'Update group' : 'Create group';
    }

    get btnModal(): string {
        return this.group.id ? 'Update' : 'Create';
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    close(group?: Group): void {
        this.matDialogRef.close(group);
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const observable = this.group.id ? this.update() : this.create();
            const subscription$ = observable.subscribe({
                next: (data) => {
                    this.close(data);
                    const message = this.group.id ? 'Successfully updated group' : 'Successfully created group';
                    this.createMessage('success', message);
                },
                error: () => {
                    this.isLoading = false;
                    const message = this.group.id
                        ? 'Update not complete. Check the fields.'
                        : 'Create not complete. Check the fields.';
                    this.createMessage('error', message);
                },
                complete: () => (this.isLoading = false),
            });
            this.subscriptions$.push(subscription$);
        }
    }

    private create(): Observable<Group> {
        const groupDto: GroupDto = {
            name: this.formGroup.controls['name'].value,
            owner: this.owner,
            permissions: [],
        };
        return this.groupService.create(groupDto);
    }

    private update(): Observable<Group> {
        return of(createMockGroup());
    }
}
