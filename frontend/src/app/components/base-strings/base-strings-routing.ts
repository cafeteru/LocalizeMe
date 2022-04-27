import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BaseStringListComponent } from './base-string-list/base-string-list.component';

export const stringRoutes: Routes = [
    {
        path: '',
        component: BaseStringListComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(stringRoutes)],
    exports: [RouterModule],
})
export class BaseStringsRouting {}
