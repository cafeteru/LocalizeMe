import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoreModule } from '../core/core.module';
import { SharedModule } from '../shared/shared.module';
import { MenuComponent } from './menu/menu.component';
import { UsersModule } from './users/users.module';
import { StagesModule } from './stages/stages.module';

@NgModule({
    declarations: [MenuComponent],
    exports: [UsersModule, StagesModule],
    imports: [CommonModule, CoreModule, SharedModule, UsersModule, StagesModule],
})
export class ComponentsModule {}
