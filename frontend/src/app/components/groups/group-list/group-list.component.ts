import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { ModalGroupComponent } from '../modal-group/modal-group.component';
import { Group } from '../../../types/group';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { sortGroupByActive, sortGroupByName, sortGroupByOwnerEmail } from '../../../shared/sorts/groups-sorts';
import { GroupService } from '../../../core/services/group.service';

@Component({
    selector: 'app-group-list',
    templateUrl: './group-list.component.html',
    styleUrls: ['./group-list.component.scss'],
})
export class GroupListComponent extends BaseComponent implements OnInit {
    currentPageGroup: readonly Group[] = [];
    groups: readonly Group[] = [];
    isLoading = false;

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
            name: 'Active',
            sortOrder: null,
            sortFn: sortGroupByActive,
            sortDirections,
        },
    ];

    constructor(private groupService: GroupService, public matDialog: MatDialog) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
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
        };
        const dialogRef = this.matDialog.open(ModalGroupComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: newGroup,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result) => {
            if (result) {
                this.loadGroups();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(group: Group): void {}

    showDeleteModal(group: Group): void {}
}
