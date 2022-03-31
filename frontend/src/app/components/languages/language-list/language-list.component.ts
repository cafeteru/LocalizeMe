import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { StageService } from '../../../core/services/stage.service';
import { MatDialog } from '@angular/material/dialog';
import { Language } from '../../../types/language';
import { ModalLanguageComponent } from '../modal-language/modal-language.component';

@Component({
    selector: 'app-language-list',
    templateUrl: './language-list.component.html',
    styleUrls: ['./language-list.component.scss'],
})
export class LanguageListComponent extends BaseComponent implements OnInit {
    isLoading = false;

    constructor(
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private stageService: StageService,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loadLanguages();
    }

    loadLanguages(): void {}

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
}
