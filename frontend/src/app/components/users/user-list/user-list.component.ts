import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { User } from '../../../types/user';
import { UserService } from '../../../core/services/user.service';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { sortEmail, sortIsActive, sortIsAdmin } from '../../../shared/sorts/users-sorts';

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

    constructor(private userService: UserService) {
        super();
    }

    override ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.userService.findAll().subscribe({
            next: (users) => (this.users = users),
            error: () => {
                this.users = [];
            },
        });
        this.subscriptions.push(subscription);
    }

    onItemChecked(id: number, checked: boolean): void {
        // this.updateCheckedSet(id, checked);
        this.refreshCheckedStatus();
    }

    onAllChecked(value: boolean): void {
        // this.listOfCurrentPageData.forEach((item) => this.updateCheckedSet(item.ID, value));
        this.refreshCheckedStatus();
    }

    onCurrentPageDataChange($event: readonly User[]): void {
        this.currentPageUsers = $event;
        this.refreshCheckedStatus();
    }

    refreshCheckedStatus(): void {
        // this.checked = this.listOfCurrentPageData.every((item) => this.setOfCheckedId.has(item.ID));
        // this.indeterminate =
        //     this.listOfCurrentPageData.some((item) => this.setOfCheckedId.has(item.ID)) && !this.checked;
    }
}
