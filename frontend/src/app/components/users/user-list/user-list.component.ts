import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { User } from '../../../types/user';
import { UserService } from '../../../core/services/user.service';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { sortEmail, sortIsActive, sortIsAdmin } from '../../../shared/sorts/users-sorts';
import { UpdateUserComponent, UpdateUserData } from '../update-user/update-user.component';
import { MatDialog } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { tap } from 'rxjs';

@Component({
    selector: 'app-user-list',
    templateUrl: './user-list.component.html',
    styleUrls: ['./user-list.component.scss'],
})
export class UserListComponent extends BaseComponent implements OnInit {
    currentPageUsers: readonly User[] = [];
    users: readonly User[] = [];
    isLoading = false;

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

    constructor(
        private userService: UserService,
        public dialog: MatDialog,
        private messageService: NzMessageService,
        private modal: NzModalService
    ) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        this.loadUsers();
    }

    loadUsers(): void {
        const subscription$ = this.userService
            .findAll()
            .pipe(tap(() => (this.isLoading = true)))
            .subscribe({
                next: (users) => {
                    this.users = users;
                    this.isLoading = false;
                },
                error: () => {
                    this.users = [];
                    this.isLoading = false;
                },
            });
        this.subscriptions$.push(subscription$);
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
        const subscription$ = dialogRef.afterClosed().subscribe((result) => {
            if (result) {
                this.loadUsers();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    showDeleteModal(user: User): void {
        this.modal.confirm({
            nzTitle: 'Are you sure delete this user?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(user),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    private delete(user: User): void {
        const subscription$ = this.userService.delete(user).subscribe((result) => {
            if (result) {
                this.loadUsers();
                this.messageService.create('success', `${user.email} has been deleted`);
            } else {
                this.messageService.create('error', 'Error deleting');
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(user: User): void {
        const subscription$ = this.userService.disable(user).subscribe(() => this.loadUsers());
        this.subscriptions$.push(subscription$);
    }
}
