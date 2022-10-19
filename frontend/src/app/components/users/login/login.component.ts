import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';
import { UserService } from '../../../core/services/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { RegisterComponent } from '../register/register.component';
import { UserDto } from '../../../types/user';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent extends BaseComponent implements OnInit {
    formGroup = new UntypedFormGroup({});
    isLoading = false;

    constructor(
        private matDialogRef: MatDialogRef<LoginComponent>,
        private nzMessageService: NzMessageService,
        private userService: UserService,
        public matDialog: MatDialog
    ) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.formGroup = new UntypedFormGroup({
            email: new UntypedFormControl('', [Validators.required, Validators.email]),
            password: new UntypedFormControl('', Validators.required),
        });
    }

    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }

    login(): void {
        this.isLoading = true;
        const user: UserDto = {
            email: this.formGroup.controls['email'].value,
            password: this.formGroup.controls['password'].value,
        };
        const subscription$ = this.userService.login(user).subscribe({
            next: (result) => {
                this.isLoading = false;
                this.close();
                if (result.active) {
                    this.createMessage('success', 'Successfully logged.');
                } else {
                    this.userService.logout();
                    this.createMessage('error', 'Session not started. User is not active.');
                }
            },
            error: () => {
                this.isLoading = false;
                this.createMessage('error', 'Session not started. Check the fields.');
            },
        });
        this.subscriptions$.push(subscription$);
    }

    close(): void {
        this.matDialogRef.close();
    }

    openRegister(): void {
        this.close();
        this.matDialog.open(RegisterComponent, {
            minWidth: '550px',
            maxWidth: '75%',
        });
    }
}
