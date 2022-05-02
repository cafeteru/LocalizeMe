import { Component, Input } from '@angular/core';
import { Translation } from '../../../types/translation';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortTranslationByActive,
    sortTranslationByLanguage,
    sortTranslationByStage,
    sortTranslationByVersion,
} from '../../../shared/sorts/translation-sorts';

@Component({
    selector: 'app-base-string-translation-table',
    templateUrl: './base-string-translation-table.component.html',
    styleUrls: ['./base-string-translation-table.component.scss'],
})
export class BaseStringTranslationTableComponent {
    @Input() translations: Translation[];
    currentPageTranslations: readonly Translation[] = [];

    listOfColumns: ColumnHeader<Translation>[] = [
        {
            name: 'Version',
            sortOrder: null,
            sortFn: sortTranslationByVersion,
            sortDirections,
        },
        {
            name: 'Stage',
            sortOrder: null,
            sortFn: sortTranslationByStage,
            sortDirections,
        },
        {
            name: 'Language',
            sortOrder: null,
            sortFn: sortTranslationByLanguage,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortTranslationByActive,
            sortDirections,
        },
    ];

    onCurrentPageDataChange($event: Translation[]) {
        this.currentPageTranslations = $event;
    }
}
