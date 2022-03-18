import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../../../core/services/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { RegisterComponent } from '../register/register.component';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent extends BaseComponent implements OnInit {
    formGroup = new FormGroup({});
    isLoading = false;

    constructor(
        private userService: UserService,
        private message: NzMessageService,
        private matDialogRef: MatDialogRef<LoginComponent>,
        public dialog: MatDialog
    ) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            email: new FormControl('', [Validators.required, Validators.email]),
            password: new FormControl('', Validators.required),
        });
    }

    createMessage(type: string, message: string): void {
        this.message.create(type, message);
    }

    login(): void {
        this.isLoading = true;
        const subscription = this.userService
            .login({
                email: this.formGroup.controls['email'].value,
                password: this.formGroup.controls['password'].value,
            })
            .subscribe({
                next: () => {
                    this.isLoading = false;
                    this.close();
                    this.createMessage('success', 'Successfully logged.');
                },
                error: () => {
                    this.isLoading = false;
                    this.createMessage('error', 'Session not started. Check the fields.');
                },
            });
        this.subscriptions.push(subscription);
    }

    close(): void {
        this.matDialogRef.close();
    }

    openRegister(): void {
        this.close();
        this.dialog.open(RegisterComponent, {
            minWidth: '550px',
            maxWidth: '75%',
        });
    }
}
