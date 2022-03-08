import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Urls } from './shared/constants/urls';
import { MenuComponent } from './components/menu/menu.component';

const routes: Routes = [
    { path: Urls.menu, component: MenuComponent },
    {
        path: '**',
        redirectTo: Urls.menu,
    },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRoutingModule {}
