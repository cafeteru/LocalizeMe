import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LanguageListComponent } from './language-list/language-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { LanguageRouting } from './language-routing';
import { ModalLanguageComponent } from './modal-language/modal-language.component';
import { LanguageFinderComponent } from './language-finder/language-finder.component';

@NgModule({
    declarations: [LanguageListComponent, ModalLanguageComponent, LanguageFinderComponent],
    imports: [CommonModule, CoreModule, SharedModule, LanguageRouting],
    exports: [LanguageFinderComponent],
})
export class LanguagesModule {}
