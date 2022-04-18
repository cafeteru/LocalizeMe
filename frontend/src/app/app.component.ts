import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from './store/app.reducer';
import { BaseComponent } from './core/base/base.component';
import { UserService } from './core/services/user.service';
import { MatDialog } from '@angular/material/dialog';
import { LoginComponent } from './components/users/login/login.component';
import { Urls } from './shared/constants/urls';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { map } from 'rxjs';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent extends BaseComponent implements OnInit {
    Urls = Urls;
    isCollapsed = false;
    isLogged = false;
    isAdmin = false;

    constructor(
        private store: Store<AppState>,
        private loginService: UserService,
        private router: Router,
        private message: NzMessageService,
        public dialog: MatDialog
    ) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => {
                this.isLogged = Boolean(user.email);
                this.isAdmin = user.admin;
                this.isCollapsed = !this.isLogged;
            });
        this.subscriptions$.push(subscription$);
    }

    showLogin(): void {
        this.dialog.open(LoginComponent, {
            maxWidth: '75%',
        });
    }

    logout(): void {
        this.router.navigateByUrl(Urls.menu).then(() => {
            const type = 'success';
            const message = 'Successfully logout.';
            this.message.create(type, message);
        });
        this.loginService.logout();
        this.isLogged = false;
        this.isAdmin = false;
    }
}
