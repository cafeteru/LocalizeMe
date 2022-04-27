import { ChangeDetectorRef, Component, EventEmitter, OnInit, Output } from '@angular/core';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { UserService } from '../../../core/services/user.service';
import { BaseComponent } from '../../../core/base/base.component';
import { checkNotNullParams, sortStrings } from '../../../shared/sorts/sort-columns';
import { Permission } from '../../../types/permission';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { map } from 'rxjs';

@Component({
    selector: 'app-user-finder',
    templateUrl: './user-finder.component.html',
    styleUrls: ['./user-finder.component.scss'],
})
export class UserFinderComponent extends BaseComponent implements OnInit {
    currentPageUsers: readonly Permission[] = [];
    isLoading = false;
    options: string[] = [];
    selectedUsers: Permission[] = [];
    users: readonly Permission[] = [];
    inputValue: string;

    @Output() emitter: EventEmitter<Permission[]> = new EventEmitter<Permission[]>();

    listOfColumns: ColumnHeader<Permission>[] = [
        {
            name: 'Email',
            sortOrder: null,
            sortFn: (a, b) => {
                const validParams = checkNotNullParams(a.user.email, b.user.email);
                return validParams === 0 ? sortStrings(a.user.email, b.user.email) : validParams;
            },
            sortDirections,
        },
    ];
    private email: string;

    constructor(
        private store: Store<AppState>,
        private userService: UserService,
        private changeDetector: ChangeDetectorRef
    ) {
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
                (this.users = users
                    .filter((user) => user.active)
                    .filter((user) => user.email !== this.email)
                    .map((user) => {
                        return {
                            user: user,
                            canWriteGroup: false,
                            id: undefined,
                        };
                    })),
        });
        this.subscriptions$.push(subscription$);
    }

    searchUserByEmail(event: Event): void {
        const emails = this.users.map((userElement) => userElement.user.email);
        this.options = emails;
        if (event) {
            const value = (event.target as HTMLInputElement).value;
            if (value) {
                this.options = emails.filter((email) => email.includes(value));
            }
        }
    }

    onCurrentPageDataChange($event: Permission[]): void {
        this.currentPageUsers = $event;
    }

    add(email: string): void {
        const searchedUser = this.users.filter((userElement) => userElement.user.email.includes(email));
        const deleteRepeatUsers = new Set([...this.selectedUsers, ...searchedUser]);
        this.selectedUsers = [...Array.from(deleteRepeatUsers)];
        this.options = [];
        this.inputValue = null;
        this.changeDetector.detectChanges();
        this.emitter.emit(this.selectedUsers);
    }

    delete(permission: Permission): void {
        this.selectedUsers = this.selectedUsers.filter((userElement) => userElement != permission);
        this.emitter.emit(this.selectedUsers);
    }
}
