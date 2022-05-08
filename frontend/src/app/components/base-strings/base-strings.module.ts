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
import { ReadXliffComponent } from './xliff/read-xliff/read-xliff.component';
import { StagesModule } from '../stages/stages.module';
import { CreateXliffComponent } from './xliff/create-xliff/create-xliff.component';
import { CreateXliffListBaseStringsComponent } from './xliff/create-xliff-list-base-strings/create-xliff-list-base-strings.component';

@NgModule({
    declarations: [
        BaseStringListComponent,
        ModalBaseStringComponent,
        ReadXliffComponent,
        CreateXliffComponent,
        CreateXliffListBaseStringsComponent,
    ],
    imports: [
        CommonModule,
        CoreModule,
        SharedModule,
        BaseStringsRouting,
        GroupsModule,
        LanguagesModule,
        TranslationsModule,
        StagesModule,
    ],
})
export class BaseStringsModule {}
