import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserInfoComponent } from './user-info/user-info.component';
import { UserListComponent } from './user-list/user-list.component';
import { UserRouting } from './user-routing';
import { UpdateUserComponent } from './update-user/update-user.component';
import { UserFinderComponent } from './user-finder/user-finder.component';

@NgModule({
    declarations: [
        LoginComponent,
        RegisterComponent,
        UserInfoComponent,
        UserListComponent,
        UpdateUserComponent,
        UserFinderComponent,
    ],
    exports: [LoginComponent, RegisterComponent, UserInfoComponent, UserFinderComponent],
    imports: [CommonModule, CoreModule, SharedModule, UserRouting],
})
export class UsersModule {}
