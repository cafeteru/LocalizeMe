import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GroupListComponent } from './group-list/group-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { GroupRouting } from './group-routing';

@NgModule({
    declarations: [GroupListComponent],
    imports: [CommonModule, CoreModule, SharedModule, GroupRouting],
})
export class GroupsModule {}
