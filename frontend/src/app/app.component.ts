import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from './store/app.reducer';
import { BaseComponent } from './core/base/base.component';
import { UserService } from './core/services/user.service';
import { MatDialog } from '@angular/material/dialog';
import { LoginComponent } from './components/users/login/login.component';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent extends BaseComponent implements OnInit {
    isCollapsed = false;
    isLogged = false;

    constructor(private store: Store<AppState>, private loginService: UserService, public dialog: MatDialog) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.store.select('user').subscribe((user) => (this.isLogged = Boolean(user.Email)));
        this.subscriptions.push(subscription);
    }

    showLogin(): void {
        this.dialog.open(LoginComponent, {
            maxWidth: '75%',
        });
    }

    logout(): void {
        this.loginService.logout();
    }
}
