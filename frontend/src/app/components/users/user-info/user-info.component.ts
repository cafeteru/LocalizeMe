import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { UpdateUserComponent } from '../update-user/update-user.component';
import { User } from '../../../types/user';
import { map } from 'rxjs';

@Component({
    selector: 'app-user-info',
    templateUrl: './user-info.component.html',
    styleUrls: ['./user-info.component.scss'],
})
export class UserInfoComponent extends BaseComponent implements OnInit {
    email = '';

    constructor(private store: Store<AppState>, public dialog: MatDialog) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => (this.email = user.email));
        this.subscriptions$.push(subscription$);
    }

    openUpdate(): void {
        const user: User = {
            id: '',
            email: this.email,
            active: true,
            admin: false,
            password: '',
        };
        const dialogRef = this.dialog.open(UpdateUserComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: {
                isAdmin: false,
                user,
            },
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: User) => {
            if (result) {
                this.email = result.email;
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
