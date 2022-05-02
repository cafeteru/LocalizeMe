import { Component, Input } from '@angular/core';
import { Permission } from '../../../types/permission';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { sortPermissionsByCanWrite, sortPermissionsByUserEmail } from '../../../shared/sorts/permissions-sorts';

@Component({
    selector: 'app-group-permission-table',
    templateUrl: './group-permission-table.component.html',
    styleUrls: ['./group-permission-table.component.scss'],
})
export class GroupPermissionTableComponent {
    @Input() permissions: Permission[];
    currentPagePermissions: readonly Permission[] = [];

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

    onCurrentPageDataChange($event: Permission[]): void {
        this.currentPagePermissions = $event;
    }
}
