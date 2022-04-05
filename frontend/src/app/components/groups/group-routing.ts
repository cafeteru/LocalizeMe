import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { GroupListComponent } from './group-list/group-list.component';

export const groupRoutes: Routes = [
    {
        path: '',
        component: GroupListComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(groupRoutes)],
    exports: [RouterModule],
})
export class GroupRouting {}
