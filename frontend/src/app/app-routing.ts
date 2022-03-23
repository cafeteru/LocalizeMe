import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Urls } from './shared/constants/urls';
import { MenuComponent } from './components/menu/menu.component';
import { IsActiveGuard } from './core/guards/is-active.guard';
import { IsAdminGuard } from './core/guards/is-admin.guard';
import { CheckTokenGuard } from './core/guards/check-token.guard';
import { LoadTokenGuard } from './core/guards/load-token.guard';

export const routes: Routes = [
    { path: Urls.menu, component: MenuComponent, canActivate: [LoadTokenGuard] },
    {
        path: Urls.users,
        loadChildren: () => import('./components/users/users.module').then((u) => u.UsersModule),
        canActivate: [LoadTokenGuard, CheckTokenGuard, IsActiveGuard, IsAdminGuard],
    },
    {
        path: Urls.stages,
        loadChildren: () => import('./components/stages/stages.module').then((s) => s.StagesModule),
        canActivate: [LoadTokenGuard, CheckTokenGuard, IsActiveGuard, IsAdminGuard],
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
