import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../core/base/base.component';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.reducer';
import { map } from 'rxjs';

@Component({
    selector: 'app-menu',
    templateUrl: './menu.component.html',
    styleUrls: ['./menu.component.scss'],
})
export class MenuComponent extends BaseComponent implements OnInit {
    isLogged = false;
    isAdmin = false;

    constructor(private store: Store<AppState>) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => {
                this.isLogged = Boolean(user.email);
                this.isAdmin = user.admin;
            });
        this.subscriptions$.push(subscription$);
    }
}
