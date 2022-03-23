import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { StageListComponent } from './stage-list/stage-list.component';

export const stageRoutes: Routes = [
    {
        path: '',
        component: StageListComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(stageRoutes)],
    exports: [RouterModule],
})
export class StageRouting {
}
