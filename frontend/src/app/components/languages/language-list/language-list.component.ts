import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MatDialog } from '@angular/material/dialog';
import { Language } from '../../../types/language';
import { ModalLanguageComponent } from '../modal-language/modal-language.component';
import { LanguageService } from '../../../core/services/language.service';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortLanguagesByActive,
    sortLanguagesByDescription,
    sortLanguagesByIsoCode,
} from '../../../shared/sorts/languages-sorts';

@Component({
    selector: 'app-language-list',
    templateUrl: './language-list.component.html',
    styleUrls: ['./language-list.component.scss'],
})
export class LanguageListComponent extends BaseComponent implements OnInit {
    currentPageLanguage: readonly Language[] = [];
    languages: readonly Language[] = [];
    isLoading = false;

    listOfColumns: ColumnHeader<Language>[] = [
        {
            name: 'Description',
            sortOrder: null,
            sortFn: sortLanguagesByDescription,
            sortDirections,
        },
        {
            name: 'IsoCode',
            sortOrder: null,
            sortFn: sortLanguagesByIsoCode,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortLanguagesByActive,
            sortDirections,
        },
    ];

    constructor(
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private languageService: LanguageService,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loadLanguages();
    }

    loadLanguages(): void {
        this.languages = [];
        this.isLoading = true;
        const subscription$ = this.languageService.findAll().subscribe({
            next: (languages) => (this.languages = languages),
            error: () => (this.isLoading = false),
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly Language[]): void {
        this.currentPageLanguage = $event;
    }

    openModal(language?: Language): void {
        const newLanguage: Language = {
            id: undefined,
            active: true,
            isoCode: undefined,
            description: undefined,
        };
        const dialogRef = this.matDialog.open(ModalLanguageComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: language ? language : newLanguage,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: Language) => {
            if (result) {
                this.loadLanguages();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    disable(language: Language): void {
        const subscription$ = this.languageService.disable(language).subscribe({
            next: () => this.loadLanguages(),
            error: () => this.nzMessageService.create('error', 'Error disabling'),
        });
        this.subscriptions$.push(subscription$);
    }

    showDeleteModal(language: Language): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this language?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(language),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    private delete(language: Language): void {
        const subscription$ = this.languageService.delete(language).subscribe((result) => {
            if (result) {
                this.loadLanguages();
                this.nzMessageService.create('success', `${language.isoCode} has been deleted`);
            } else {
                this.nzMessageService.create('error', 'Error deleting');
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
