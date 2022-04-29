import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BaseStringListComponent } from './base-string-list/base-string-list.component';
import { CoreModule } from '../../core/core.module';
import { SharedModule } from '../../shared/shared.module';
import { BaseStringsRouting } from './base-strings-routing';
import { ModalBaseStringComponent } from './modal-base-string/modal-base-string.component';
import { GroupsModule } from '../groups/groups.module';
import { LanguagesModule } from '../languages/languages.module';
import { TranslationsModule } from '../translations/translations.module';

@NgModule({
    declarations: [BaseStringListComponent, ModalBaseStringComponent],
    imports: [
        CommonModule,
        CoreModule,
        SharedModule,
        BaseStringsRouting,
        GroupsModule,
        LanguagesModule,
        TranslationsModule,
    ],
})
export class BaseStringsModule {}
