import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Urls } from './shared/constants/urls';
import { MenuComponent } from './components/menu/menu.component';
import { IsActiveGuard } from './core/guards/is-active.guard';
import { IsAdminGuard } from './core/guards/is-admin.guard';
import { CheckTokenGuard } from './core/guards/check-token.guard';

const routes: Routes = [
    { path: Urls.menu, component: MenuComponent },
    {
        path: Urls.users,
        loadChildren: () => import('./components/users/users.module').then((u) => u.UsersModule),
        canActivate: [CheckTokenGuard, IsActiveGuard, IsAdminGuard],
    },
    {
        path: '**',
        redirectTo: Urls.menu,
    },
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule],
})
export class AppRouting {}
