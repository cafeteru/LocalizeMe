import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GroupListComponent } from './group-list/group-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { GroupRouting } from './group-routing';
import { ModalGroupComponent } from './modal-group/modal-group.component';
import { UsersModule } from '../users/users.module';
import { GroupFinderComponent } from './group-finder/group-finder.component';
import { GroupPermissionTableComponent } from './group-permission-table/group-permission-table.component';

@NgModule({
    declarations: [GroupListComponent, ModalGroupComponent, GroupFinderComponent, GroupPermissionTableComponent],
    imports: [CommonModule, CoreModule, SharedModule, GroupRouting, UsersModule],
    exports: [GroupFinderComponent],
})
export class GroupsModule {}
