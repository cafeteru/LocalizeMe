import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { BaseComponent } from '../../../core/base/base.component';

@Component({
    selector: 'app-user-info',
    templateUrl: './user-info.component.html',
    styleUrls: ['./user-info.component.scss'],
})
export class UserInfoComponent extends BaseComponent implements OnInit {
    email = '';

    constructor(private store: Store<AppState>) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.store.select('user').subscribe((user) => (this.email = user.Email));
        this.subscriptions.push(subscription);
    }
}
