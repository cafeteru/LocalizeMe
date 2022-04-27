import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MatDialog } from '@angular/material/dialog';
import { BaseStringService } from '../../../core/services/base-string.service';
import { ModalBaseStringComponent } from '../modal-base-string/modal-base-string.component';
import { BaseString } from '../../../types/base-string';

@Component({
    selector: 'app-base-string-list',
    templateUrl: './base-string-list.component.html',
    styleUrls: ['./base-string-list.component.scss'],
})
export class BaseStringListComponent extends BaseComponent implements OnInit {
    isLoading = false;

    constructor(
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private baseStringService: BaseStringService,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
    }

    loadBaseStrings(): void {
        // this.languages = [];
        // this.isLoading = true;
        // const subscription$ = this.languageService.findAll().subscribe({
        //     next: (languages) => (this.languages = languages),
        //     error: () => (this.isLoading = false),
        //     complete: () => (this.isLoading = false),
        // });
        // this.subscriptions$.push(subscription$);
    }

    openModal(baseString?: BaseString): void {
        const newBaseString: BaseString = {
            id: undefined,
            active: true,
            author: undefined,
            group: undefined,
            identifier: '',
            sourceLanguage: undefined,
            translations: [],
        };
        const dialogRef = this.matDialog.open(ModalBaseStringComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: baseString ? baseString : newBaseString,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: BaseString) => {
            if (result) {
                // this.loadLanguages();
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
