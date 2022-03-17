import { Component, Inject, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { BaseComponent } from '../../../core/base/base.component';
import { UserService } from '../../../core/services/user.service';
import { IsSameValidator } from '../../../core/validators/is-same-validator';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { User } from '../../../types/user';

@Component({
    selector: 'app-update-user',
    templateUrl: './update-user.component.html',
    styleUrls: ['./update-user.component.scss'],
})
export class UpdateUserComponent extends BaseComponent implements OnInit {
    @Input() isVisible = false;
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<UpdateUserComponent>,
        @Inject(MAT_DIALOG_DATA) private data: { isAdmin: false },
        private userService: UserService,
        private message: NzMessageService
    ) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.initFormGroup();
        if (!this.data.isAdmin) {
            const subscription = this.userService.findMe().subscribe((user) => {
                this.formGroup.controls['email'].setValue(user.Email);
            });
            this.subscriptions.push(subscription);
        }
    }

    wantToChangePassword(): boolean {
        return this.formGroup.controls['changePassword'].value;
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const password = this.formGroup.controls['password'].value;
            const user: User = {
                ID: '',
                Email: this.formGroup.controls['email'].value,
                Password: password ? password : '',
                IsActive: this.data.isAdmin ? this.formGroup.controls['IsActive'].value : true,
                IsAdmin: this.data.isAdmin ? this.formGroup.controls['IsAdmin'].value : false,
            };
            const subscription = this.userService.updateMe(user).subscribe({
                next: () => {
                    this.isLoading = false;
                    this.close();
                    this.createMessage('success', 'Successfully updated.');
                },
                error: () => {
                    this.isLoading = false;
                    this.createMessage('error', 'Update not complete. Check the fields.');
                },
            });
            this.subscriptions.push(subscription);
        }
    }

    createMessage(type: string, message: string): void {
        this.message.create(type, message);
    }

    close(): void {
        this.matDialogRef.close();
    }

    private initFormGroup(): void {
        const password = 'password';
        const checkPassword = 'checkPassword';
        this.formGroup = new FormGroup(
            {
                email: new FormControl('', [Validators.required, Validators.email]),
                changePassword: new FormControl(false, Validators.required),
                password: new FormControl(''),
                checkPassword: new FormControl(''),
            },
            { validators: IsSameValidator.isValid(password, checkPassword) }
        );

        const subscription = this.formGroup.controls['changePassword'].valueChanges.subscribe((res) => {
            if (res) {
                FormGroupUtil.changeValidator(this.formGroup, password, [Validators.required], '');
                FormGroupUtil.changeValidator(this.formGroup, checkPassword, [Validators.required], '');
            } else {
                FormGroupUtil.changeValidator(this.formGroup, password, []);
                FormGroupUtil.changeValidator(this.formGroup, checkPassword, []);
            }
        });
        this.subscriptions.push(subscription);
    }
}
