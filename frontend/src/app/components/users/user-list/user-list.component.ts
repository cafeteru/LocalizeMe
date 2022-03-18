import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { User } from '../../../types/user';
import { UserService } from '../../../core/services/user.service';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { sortEmail, sortIsActive, sortIsAdmin } from '../../../shared/sorts/users-sorts';
import { UpdateUserComponent, UpdateUserData } from '../update-user/update-user.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
    selector: 'app-user-list',
    templateUrl: './user-list.component.html',
    styleUrls: ['./user-list.component.scss'],
})
export class UserListComponent extends BaseComponent implements OnInit {
    currentPageUsers: readonly User[] = [];
    users: readonly User[] = [];

    listOfColumns: ColumnHeader<User>[] = [
        {
            name: 'Email',
            sortOrder: null,
            sortFn: sortEmail,
            sortDirections,
        },
        {
            name: 'Admin',
            sortOrder: null,
            sortFn: sortIsAdmin,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortIsActive,
            sortDirections,
        },
    ];

    constructor(private userService: UserService, public dialog: MatDialog) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.userService.findAll().subscribe({
            next: (users) => (this.users = users),
            error: () => (this.users = []),
        });
        this.subscriptions.push(subscription);
    }

    onCurrentPageDataChange($event: readonly User[]): void {
        this.currentPageUsers = $event;
    }

    openUpdate(user: User): void {
        const data: UpdateUserData = {
            isAdmin: true,
            user,
        };
        const dialogRef = this.dialog.open(UpdateUserComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data,
        });
        const subscription = dialogRef.afterClosed().subscribe((result) => this.updateUsers(result));
        this.subscriptions.push(subscription);
    }

    disable(user: User): void {
        const subscription = this.userService.disable(user).subscribe((result) => this.updateUsers(result));
        this.subscriptions.push(subscription);
    }

    private updateUsers(result: User): void {
        this.users = this.users.map((value) => (value.ID === result.ID ? { ...result } : { ...value }));
    }
}
