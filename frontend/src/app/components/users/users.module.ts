import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserInfoComponent } from './user-info/user-info.component';

@NgModule({
    declarations: [LoginComponent, RegisterComponent, UserInfoComponent],
    exports: [LoginComponent, RegisterComponent, UserInfoComponent],
    imports: [CommonModule, CoreModule, SharedModule],
})
export class UsersModule {}
