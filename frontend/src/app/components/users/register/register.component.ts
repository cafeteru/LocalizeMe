import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { IsSameValidator } from '../../../core/validators/is-same-validator';
import { FormGroupUtil } from '../../../shared/utils/form-group-util';
import { UserService } from '../../../core/services/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { switchMap } from 'rxjs';
import { LoginComponent } from '../login/login.component';
import { UserDto } from '../../../types/user';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
})
export class RegisterComponent extends BaseComponent implements OnInit {
    @Input() isVisible = false;
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<RegisterComponent>,
        private userService: UserService,
        private nzMessageService: NzMessageService,
        public matDialog: MatDialog
    ) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.formGroup = new FormGroup(
            {
                email: new FormControl('', [Validators.required, Validators.email]),
                password: new FormControl('', Validators.required),
                checkPassword: new FormControl('', Validators.required),
            },
            { validators: IsSameValidator.isValid('password', 'checkPassword') }
        );
    }

    send(): void {
        if (FormGroupUtil.valid(this.formGroup)) {
            this.isLoading = true;
            const user: UserDto = {
                email: this.formGroup.controls['email'].value,
                password: this.formGroup.controls['password'].value,
            };
            const subscription$ = this.userService
                .register(user)
                .pipe(switchMap(() => this.userService.login(user)))
                .subscribe({
                    next: () => {
                        this.isLoading = false;
                        this.close();
                        this.createMessage('success', 'Successfully registered.');
                    },
                    error: () => {
                        this.isLoading = false;
                        this.createMessage('error', 'Register not complete. Check the fields.');
                    },
                });
            this.subscriptions$.push(subscription$);
        }
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    close(): void {
        this.matDialogRef.close();
    }

    openRegister(): void {
        this.close();
        this.matDialog.open(LoginComponent, {
            maxWidth: '75%',
        });
    }
}
