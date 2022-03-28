import { Component, Inject, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseComponent } from '../../../core/base/base.component';
import { UserService } from '../../../core/services/user.service';
import { IsSameValidator } from '../../../core/validators/is-same-validator';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { User } from '../../../types/user';
import { Observable } from 'rxjs';

export interface UpdateUserData {
    isAdmin: boolean;
    user: User;
}

@Component({
    selector: 'app-update-user',
    templateUrl: './update-user.component.html',
    styleUrls: ['./update-user.component.scss'],
})
export class UpdateUserComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<UpdateUserComponent>,
        @Inject(MAT_DIALOG_DATA) public data: UpdateUserData,
        private userService: UserService,
        private message: NzMessageService
    ) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.initFormGroup();
    }

    wantToChangePassword(): boolean {
        return this.formGroup.controls['changePassword'].value;
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const password = this.formGroup.controls['password'].value;
            const user: User = {
                ...this.data.user,
                Email: this.formGroup.controls['email'].value,
                Password: password ? password : '',
                Active: this.formGroup.controls['isActive'].value,
                Admin: this.formGroup.controls['isAdmin'].value,
            };
            const subscription = this.getObservable(user).subscribe({
                next: () => {
                    this.isLoading = false;
                    this.close(user);
                    this.createMessage('success', 'Successfully updated.');
                },
                error: () => {
                    this.isLoading = false;
                    this.createMessage('error', 'Update not complete. Check the fields.');
                },
            });
            this.subscriptions$.push(subscription);
        }
    }

    private getObservable(user: User): Observable<User> {
        return this.data.isAdmin ? this.userService.update(user) : this.userService.updateMe(user);
    }

    createMessage(type: string, message: string): void {
        this.message.create(type, message);
    }

    close(user: User): void {
        this.matDialogRef.close(user);
    }

    private initFormGroup(): void {
        const password = 'password';
        const checkPassword = 'checkPassword';
        this.formGroup = new FormGroup(
            {
                email: new FormControl(this.data.user.Email, [Validators.required, Validators.email]),
                changePassword: new FormControl(false, Validators.required),
                password: new FormControl(''),
                checkPassword: new FormControl(''),
                isActive: new FormControl(this.data.user.Active),
                isAdmin: new FormControl(this.data.user.Admin),
            },
            { validators: IsSameValidator.isValid(password, checkPassword) }
        );

        const subscription = this.formGroup.controls['changePassword'].valueChanges.subscribe((value) => {
            if (value) {
                FormGroupUtil.changeValidator(this.formGroup, password, [Validators.required], '');
                FormGroupUtil.changeValidator(this.formGroup, checkPassword, [Validators.required], '');
            } else {
                FormGroupUtil.changeValidator(this.formGroup, password, []);
                FormGroupUtil.changeValidator(this.formGroup, checkPassword, []);
            }
        });
        this.subscriptions$.push(subscription);
    }
}
