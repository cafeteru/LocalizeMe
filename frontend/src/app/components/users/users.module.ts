import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserInfoComponent } from './user-info/user-info.component';
import { UserListComponent } from './user-list/user-list.component';
import { UsersRouting } from './users-routing';

@NgModule({
    declarations: [LoginComponent, RegisterComponent, UserInfoComponent, UserListComponent],
    exports: [LoginComponent, RegisterComponent, UserInfoComponent],
    imports: [CommonModule, CoreModule, SharedModule, UsersRouting],
})
export class UsersModule {}
