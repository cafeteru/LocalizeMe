import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from './store/app.reducer';
import { BaseComponent } from './core/base/base.component';
import { LoginService } from './core/services/login.service';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent extends BaseComponent implements OnInit {
    isCollapsed = false;
    isVisible = false;
    isLogged = false;

    constructor(private store: Store<AppState>, private loginService: LoginService) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.store.select('user').subscribe((user) => (this.isLogged = Boolean(user.Email)));
        this.subscriptions.push(subscription);
    }

    showModal(): void {
        this.setIsVisible(true);
    }

    setIsVisible(value: boolean) {
        this.isVisible = value;
    }

    logout(): void {
        this.loginService.logout();
    }
}
