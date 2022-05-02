import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TranslationListComponent } from './translation-list/translation-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { ModalTranslationComponent } from './modal-translation/modal-translation.component';
import { LanguagesModule } from '../languages/languages.module';
import { StagesModule } from '../stages/stages.module';
import { BaseStringTranslationTableComponent } from './base-string-translation-table/base-string-translation-table.component';

@NgModule({
    declarations: [TranslationListComponent, ModalTranslationComponent, BaseStringTranslationTableComponent],
    imports: [CommonModule, CoreModule, SharedModule, LanguagesModule, StagesModule],
    exports: [TranslationListComponent, BaseStringTranslationTableComponent],
})
export class TranslationsModule {}
