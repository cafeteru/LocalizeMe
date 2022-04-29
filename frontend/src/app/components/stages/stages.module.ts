import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StageListComponent } from './stage-list/stage-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { StageRouting } from './stage-routing';
import { ModalStageComponent } from './modal-stage/modal-stage.component';
import { StageFinderComponent } from './stage-finder/stage-finder.component';

@NgModule({
    declarations: [StageListComponent, ModalStageComponent, StageFinderComponent],
    imports: [CommonModule, CoreModule, SharedModule, StageRouting],
    exports: [StageFinderComponent],
})
export class StagesModule {}
