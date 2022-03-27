import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StageListComponent } from './stage-list/stage-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { StageRouting } from './stage-routing.module';
import { ModalStageComponent } from './modal-stage/modal-stage.component';

@NgModule({
    declarations: [StageListComponent, ModalStageComponent],
    imports: [CommonModule, CoreModule, SharedModule, StageRouting],
})
export class StagesModule {}
