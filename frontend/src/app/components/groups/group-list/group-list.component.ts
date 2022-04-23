import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { ModalGroupComponent } from '../modal-group/modal-group.component';
import { Group } from '../../../types/group';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortGroupByActive,
    sortGroupByName,
    sortGroupByOwnerEmail,
    sortGroupByPublic,
} from '../../../shared/sorts/groups-sorts';
import { GroupService } from '../../../core/services/group.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { createMockUser, User } from '../../../types/user';

@Component({
    selector: 'app-group-list',
    templateUrl: './group-list.component.html',
    styleUrls: ['./group-list.component.scss'],
})
export class GroupListComponent extends BaseComponent implements OnInit {
    currentPageGroup: readonly Group[] = [];
    groups: readonly Group[] = [];
    isLoading = false;
    user: User = createMockUser();

    listOfColumns: ColumnHeader<Group>[] = [
        {
            name: 'Name',
            sortOrder: null,
            sortFn: sortGroupByName,
            sortDirections,
        },
        {
            name: 'Owner',
            sortOrder: null,
            sortFn: sortGroupByOwnerEmail,
            sortDirections,
        },
        {
            name: 'Public',
            sortOrder: null,
            sortFn: sortGroupByPublic,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortGroupByActive,
            sortDirections,
        },
    ];

    constructor(
        private store: Store<AppState>,
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private groupService: GroupService,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => (this.user = user));
        this.subscriptions$.push(subscription$);
        this.loadGroups();
    }

    loadGroups(): void {
        this.groups = [];
        this.isLoading = true;
        const subscription$ = this.groupService.findAll().subscribe({
            next: (groups) => (this.groups = groups),
            error: () => (this.isLoading = false),
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly Group[]): void {
        this.currentPageGroup = $event;
    }

    openModal(group?: Group): void {
        const newGroup: Group = {
            id: undefined,
            active: true,
            name: undefined,
            permissions: [],
            owner: undefined,
            public: true,
        };
        const dialogRef = this.matDialog.open(ModalGroupComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: group ? group : newGroup,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result) => {
            if (result) {
                this.loadGroups();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(group: Group): void {
        const subscription$ = this.groupService.disable(group).subscribe({
            next: () => this.loadGroups(),
            error: () => this.nzMessageService.create('error', 'Error disabling'),
        });
        this.subscriptions$.push(subscription$);
    }

    showDeleteModal(group: Group): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this group?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(group),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    canEdit(group: Group): boolean {
        if (this.user.admin || group.public || group.owner.id === this.user.id) {
            return true;
        }
        group.permissions.forEach((permission) => {
            if (permission.user.id === this.user.id && permission.canWriteGroup) {
                return true;
            }
        });
        return false;
    }

    canDelete(group: Group): boolean {
        return this.user.admin || group.owner.id === this.user.id;
    }

    private delete(group: Group): void {
        const subscription$ = this.groupService.delete(group).subscribe((result) => {
            if (result) {
                this.loadGroups();
                this.nzMessageService.create('success', `${group.name} has been deleted`);
            } else {
                this.nzMessageService.create('error', 'Error deleting');
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
