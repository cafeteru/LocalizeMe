import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../core/base/base.component';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { LoginService } from '../../core/services/login.service';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent extends BaseComponent implements OnInit {
    @Input() isVisible = false;
    @Output() emitter = new EventEmitter<boolean>();

    formGroup: FormGroup;
    isOkLoading = false;

    constructor(private loginService: LoginService, private message: NzMessageService) {
        super();
    }

    override ngOnInit() {
        super.ngOnInit();
        this.formGroup = new FormGroup({
            email: new FormControl('', Validators.required),
            password: new FormControl('', Validators.required),
        });
    }

    createMessage(type: string, message: string): void {
        this.message.create(type, message);
    }

    login(): void {
        this.isOkLoading = true;
        const subscription = this.loginService
            .login({
                email: this.formGroup.controls['email'].value,
                password: this.formGroup.controls['password'].value,
            })
            .subscribe({
                next: () => {
                    this.isOkLoading = false;
                    this.emitIsVisible(false);
                    this.createMessage('success', 'Successfully logged.');
                },
                error: () => {
                    this.isOkLoading = false;
                    this.createMessage('error', 'Session not started. Check the fields.');
                },
            });
        this.subscriptions.push(subscription);
    }

    handleCancel(): void {
        this.emitIsVisible(false);
    }

    private emitIsVisible(value: boolean): void {
        this.isVisible = value;
        this.emitter.emit(this.isVisible);
    }
}
