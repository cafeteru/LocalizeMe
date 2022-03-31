import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LanguageListComponent } from './language-list/language-list.component';

export const languageRoutes: Routes = [
    {
        path: '',
        component: LanguageListComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(languageRoutes)],
    exports: [RouterModule],
})
export class LanguageRouting {}
