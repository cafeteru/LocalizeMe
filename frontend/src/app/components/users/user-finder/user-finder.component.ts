import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { UserService } from '../../../core/services/user.service';
import { BaseComponent } from '../../../core/base/base.component';
import { Permission } from '../../../types/permission';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { map } from 'rxjs';
import { sortPermissionsByCanWrite, sortPermissionsByUserEmail } from '../../../shared/sorts/permissions-sorts';

@Component({
    selector: 'app-user-finder',
    templateUrl: './user-finder.component.html',
    styleUrls: ['./user-finder.component.scss'],
})
export class UserFinderComponent extends BaseComponent implements OnInit {
    currentPagePermissions: readonly Permission[] = [];
    isLoading = false;
    options: string[] = [];
    @Input() selectedPermissions: Permission[] = [];
    permissions: readonly Permission[] = [];
    inputValue: string;

    @Output() emitter: EventEmitter<Permission[]> = new EventEmitter<Permission[]>();

    listOfColumns: ColumnHeader<Permission>[] = [
        {
            name: 'Email',
            sortOrder: null,
            sortFn: sortPermissionsByUserEmail,
            sortDirections,
        },
        {
            name: 'Can write?',
            sortOrder: null,
            sortFn: sortPermissionsByCanWrite,
            sortDirections,
        },
    ];
    private email: string;

    constructor(private store: Store<AppState>, private userService: UserService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        const userSubscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => (this.email = user.email));
        this.subscriptions$.push(userSubscription$);
        const subscription$ = this.userService.findAll().subscribe({
            next: (users) =>
                (this.permissions = users
                    .filter((user) => user.active)
                    .filter((user) => !user.admin)
                    .filter((user) => user.email !== this.email)
                    .filter((user) =>
                        this.selectedPermissions && this.selectedPermissions.length > 0
                            ? this.selectedPermissions
                                  .map((permission) => permission.user)
                                  .some((value) => user.id !== value.id)
                            : true
                    )
                    .map((user) => {
                        return {
                            user: user,
                            canWrite: false,
                        };
                    })),
        });
        this.subscriptions$.push(subscription$);
    }

    searchUserByEmail(event: Event): void {
        const emails = this.permissions.map((permission) => permission.user.email);
        this.options = emails;
        if (event) {
            const value = (event.target as HTMLInputElement).value;
            if (value) {
                this.options = emails.filter((email) => email.toLocaleLowerCase().includes(value.toLocaleLowerCase()));
            }
        }
    }

    onCurrentPageDataChange($event: Permission[]): void {
        this.currentPagePermissions = $event;
    }

    add(email: string): void {
        const searchedUser = this.permissions.filter((userElement) => userElement.user.email.includes(email));
        this.selectedPermissions = this.selectedPermissions ? this.selectedPermissions : [];
        const deleteRepeatUsers = new Set([...this.selectedPermissions, ...searchedUser]);
        this.selectedPermissions = [...Array.from(deleteRepeatUsers)];
        this.options = [];
        this.inputValue = null;
        this.emitter.emit(this.selectedPermissions);
    }

    updateCanWrite(permission: Permission) {
        permission.canWrite = !permission.canWrite;
        this.emitter.emit(this.selectedPermissions);
    }

    delete(permission: Permission): void {
        this.selectedPermissions = this.selectedPermissions.filter((userElement) => userElement != permission);
        this.emitter.emit(this.selectedPermissions);
    }
}
