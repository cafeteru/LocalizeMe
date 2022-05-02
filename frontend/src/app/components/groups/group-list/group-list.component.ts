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

interface GroupData {
    group: Group;
    expanded: boolean;
}

@Component({
    selector: 'app-group-list',
    templateUrl: './group-list.component.html',
    styleUrls: ['./group-list.component.scss'],
})
export class GroupListComponent extends BaseComponent implements OnInit {
    currentPageGroupData: readonly GroupData[] = [];
    groupData: readonly GroupData[] = [];
    isLoading = false;
    user: User = createMockUser();

    listOfColumns: ColumnHeader<GroupData>[] = [
        {
            name: 'Name',
            sortOrder: null,
            sortFn: (a, b) => sortGroupByName(a.group, b.group),
            sortDirections,
        },
        {
            name: 'Owner',
            sortOrder: null,
            sortFn: (a, b) => sortGroupByOwnerEmail(a.group, b.group),
            sortDirections,
        },
        {
            name: 'Public',
            sortOrder: null,
            sortFn: (a, b) => sortGroupByPublic(a.group, b.group),
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: (a, b) => sortGroupByActive(a.group, b.group),
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
        this.groupData = [];
        this.isLoading = true;
        const subscription$ = this.groupService
            .findAll()
            .pipe(
                map((groups) =>
                    groups.map((group) => {
                        const groupData: GroupData = {
                            group,
                            expanded: false,
                        };
                        return groupData;
                    })
                )
            )
            .subscribe({
                next: (groupData) => (this.groupData = groupData),
                error: () => (this.isLoading = false),
                complete: () => (this.isLoading = false),
            });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly GroupData[]): void {
        this.currentPageGroupData = $event;
    }

    openModal(groupData?: GroupData): void {
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
            data: groupData ? groupData.group : newGroup,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result) => {
            if (result) {
                this.loadGroups();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(groupData?: GroupData): void {
        const { group } = groupData;
        const subscription$ = this.groupService.disable(group).subscribe({
            next: () => this.loadGroups(),
            error: () => this.nzMessageService.create('error', 'Error disabling'),
        });
        this.subscriptions$.push(subscription$);
    }

    showDeleteModal(groupData?: GroupData): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this group?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(groupData),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    canEdit(groupData?: GroupData): boolean {
        const { group } = groupData;
        if (this.user.admin || group.public || group.owner.id === this.user.id) {
            return true;
        }
        return group.permissions.some((permission) => permission.user.id === this.user.id && permission.canWrite);
    }

    canDelete(groupData?: GroupData): boolean {
        const { group } = groupData;
        return this.user.admin || group.owner.id === this.user.id;
    }

    private delete(groupData?: GroupData): void {
        const { group } = groupData;
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

    onExpandChange(groupData: GroupData, expanded: boolean): void {
        groupData.expanded = expanded;
    }
}
