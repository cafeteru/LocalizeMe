import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserListComponent } from './user-list/user-list.component';

export const userRoutes: Routes = [
    {
        path: '',
        component: UserListComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(userRoutes)],
    exports: [RouterModule],
})
export class UserRouting {}
