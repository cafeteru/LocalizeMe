import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { Translation } from '../../../types/translation';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortTranslationByActive,
    sortTranslationByLanguage,
    sortTranslationByStage,
    sortTranslationByVersion,
} from '../../../shared/sorts/translation-sorts';
import { MatDialog } from '@angular/material/dialog';
import { ModalTranslationComponent } from '../modal-translation/modal-translation.component';
import { NzModalService } from 'ng-zorro-antd/modal';

@Component({
    selector: 'app-translation-list',
    templateUrl: './translation-list.component.html',
    styleUrls: ['./translation-list.component.scss'],
})
export class TranslationListComponent extends BaseComponent implements OnInit {
    @Input() translations: Translation[] = [];
    currentPageTranslations: readonly Translation[] = [];
    @Output() emitter: EventEmitter<Translation[]> = new EventEmitter<Translation[]>();

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

    constructor(private nzModalService: NzModalService, public matDialog: MatDialog) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
    }

    onCurrentPageDataChange($event: Translation[]) {
        this.currentPageTranslations = $event;
    }

    openModal(translation?: Translation): void {
        const newTranslation: Translation = {
            version: 1,
            stage: undefined,
            active: true,
            author: undefined,
            content: undefined,
            date: new Date(),
            language: undefined,
        };
        const dialogRef = this.matDialog.open(ModalTranslationComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: translation ? translation : newTranslation,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: Translation) => {
            if (result) {
                this.translations = [...this.translations, result];
            }
            this.emitter.emit(this.translations);
        });
        this.subscriptions$.push(subscription$);
    }

    disable(translation: Translation): void {
        translation.active = !translation.active;
        this.emitter.emit(this.translations);
    }

    showDeleteModal(translation: Translation): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this translation?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(translation),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    private delete(translation: Translation): void {
        this.translations = this.translations.filter((element) => element.content != translation.content);
    }
}
