import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { UpdateUserComponent } from '../update-user/update-user.component';
import { createMockUser, User } from '../../../types/user';
import { map } from 'rxjs';
import { UserService } from '../../../core/services/user.service';

@Component({
    selector: 'app-user-info',
    templateUrl: './user-info.component.html',
    styleUrls: ['./user-info.component.scss'],
})
export class UserInfoComponent extends BaseComponent implements OnInit {
    user: User = createMockUser();

    constructor(private store: Store<AppState>, public dialog: MatDialog, private userService: UserService) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => (this.user = user));
        this.subscriptions$.push(subscription$);
    }

    openUpdate(): void {
        const originalUser = { ...this.user };
        const dialogRef = this.dialog.open(UpdateUserComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: {
                isAdmin: false,
                user: this.user,
            },
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: User) => {
            if (result) {
                this.user = result;
            }
            if (originalUser.email != result.email) {
                this.userService.logout();
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
