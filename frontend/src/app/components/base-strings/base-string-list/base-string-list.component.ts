import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { MatDialog } from '@angular/material/dialog';
import { BaseStringService } from '../../../core/services/base-string.service';
import { ModalBaseStringComponent } from '../modal-base-string/modal-base-string.component';
import { BaseString } from '../../../types/base-string';
import { Group } from '../../../types/group';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortBaseStringByActive,
    sortBaseStringByAuthor,
    sortBaseStringByGroup,
    sortBaseStringByIdentifier,
    sortBaseStringBySourceLanguage,
} from '../../../shared/sorts/base-string-sorts';
import { Store } from '@ngrx/store';
import { AppState } from '../../../store/app.reducer';
import { map } from 'rxjs';
import { createMockUser, User } from '../../../types/user';

@Component({
    selector: 'app-base-string-list',
    templateUrl: './base-string-list.component.html',
    styleUrls: ['./base-string-list.component.scss'],
})
export class BaseStringListComponent extends BaseComponent implements OnInit {
    currentPageBaseStrings: readonly BaseString[] = [];
    isLoading = false;
    originalBaseStrings: BaseString[];
    baseStrings: BaseString[];
    filterText = '';
    user: User = createMockUser();

    listOfColumns: ColumnHeader<BaseString>[] = [
        {
            name: 'Identifier',
            sortOrder: null,
            sortFn: sortBaseStringByIdentifier,
            sortDirections,
        },
        {
            name: 'Language',
            sortOrder: null,
            sortFn: sortBaseStringBySourceLanguage,
            sortDirections,
        },
        {
            name: 'Group',
            sortOrder: null,
            sortFn: sortBaseStringByGroup,
            sortDirections,
        },
        {
            name: 'Author',
            sortOrder: null,
            sortFn: sortBaseStringByAuthor,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortBaseStringByActive,
            sortDirections,
        },
    ];

    constructor(
        private nzMessageService: NzMessageService,
        private nzModalService: NzModalService,
        private baseStringService: BaseStringService,
        private store: Store<AppState>,
        public matDialog: MatDialog
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        const subscription$ = this.store
            .select('userInfo')
            .pipe(map((userReducer) => userReducer.user))
            .subscribe((user) => (this.user = user));
        this.subscriptions$.push(subscription$);
        this.loadBaseStrings();
    }

    filterStrings($event: string) {
        this.baseStrings = this.originalBaseStrings.filter(
            (baseString) =>
                baseString.identifier.includes($event) ||
                baseString.sourceLanguage.isoCode.includes($event) ||
                baseString.group.name.includes($event) ||
                baseString.author.email.includes($event)
        );
    }

    loadBaseStrings(): void {
        this.originalBaseStrings = [];
        this.baseStrings = [];
        this.isLoading = true;
        const subscription$ = this.baseStringService.findAll().subscribe({
            next: (baseStrings) => {
                this.originalBaseStrings = baseStrings;
                this.baseStrings = baseStrings;
            },
            error: () => (this.isLoading = false),
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
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
                this.loadBaseStrings();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly BaseString[]): void {
        this.currentPageBaseStrings = $event;
    }

    canEdit(baseString: BaseString): boolean {
        if (this.user.admin || !baseString || baseString.group.public || baseString.group.owner.id === this.user.id) {
            return true;
        }
        baseString.group.permissions.forEach((permission) => {
            if (permission.user.id === this.user.id && permission.canWriteGroup) {
                return true;
            }
        });
        return false;
    }
}
