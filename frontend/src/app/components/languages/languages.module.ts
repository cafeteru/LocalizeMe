import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LanguageListComponent } from './language-list/language-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { LanguageRouting } from './language-routing';
import { ModalLanguageComponent } from './modal-language/modal-language.component';

@NgModule({
    declarations: [LanguageListComponent, ModalLanguageComponent],
    imports: [CommonModule, CoreModule, SharedModule, LanguageRouting],
})
export class LanguagesModule {}
