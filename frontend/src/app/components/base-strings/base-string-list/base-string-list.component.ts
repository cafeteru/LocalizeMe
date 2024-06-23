import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Store } from '@ngrx/store';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { map } from 'rxjs';
import { BaseComponent } from '../../../core/base/base.component';
import { BaseStringService } from '../../../core/services/base-string.service';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import {
    sortBaseStringByActive,
    sortBaseStringByAuthor,
    sortBaseStringByGroup,
    sortBaseStringByIdentifier,
    sortBaseStringByPage,
    sortBaseStringBySourceLanguage,
} from '../../../shared/sorts/base-string-sorts';
import { AppState } from '../../../store/app.reducer';
import { BaseString } from '../../../types/base-string';
import { User, createMockUser } from '../../../types/user';
import { ModalBaseStringComponent } from '../modal-base-string/modal-base-string.component';
import { CreateXliffComponent } from '../xliff/create-xliff/create-xliff.component';
import { ReadXliffComponent } from '../xliff/read-xliff/read-xliff.component';

interface BaseStringData {
    baseString: BaseString;
    expanded: boolean;
}

@Component({
    selector: 'app-base-string-list',
    templateUrl: './base-string-list.component.html',
    styleUrls: ['./base-string-list.component.scss'],
})
export class BaseStringListComponent extends BaseComponent implements OnInit {
    currentPageBaseStrings: readonly BaseStringData[] = [];
    isLoading = false;
    originalBaseStrings: BaseStringData[];
    baseStrings: BaseStringData[];
    filterText = '';
    user: User = createMockUser();

    listOfColumns: ColumnHeader<BaseStringData>[] = [
        {
            name: 'Identifier',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringByIdentifier(a.baseString, b.baseString),
            sortDirections,
        },
        {
            name: 'Language',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringBySourceLanguage(a.baseString, b.baseString),
            sortDirections,
        },
        {
            name: 'Group',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringByGroup(a.baseString, b.baseString),
            sortDirections,
        },
        {
            name: 'Page',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringByPage(a.baseString, b.baseString),
            sortDirections,
        },
        {
            name: 'Author',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringByAuthor(a.baseString, b.baseString),
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: (a, b) => sortBaseStringByActive(a.baseString, b.baseString),
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
            ({ baseString }) =>
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
        const subscription$ = this.baseStringService
            .findAll()
            .pipe(
                map((baseStrings) =>
                    baseStrings.map((baseString) => {
                        const baseStringData: BaseStringData = {
                            baseString,
                            expanded: false,
                        };
                        return baseStringData;
                    })
                )
            )
            .subscribe({
                next: (baseStrings) => {
                    this.originalBaseStrings = baseStrings;
                    this.baseStrings = baseStrings;
                },
                error: () => (this.isLoading = false),
                complete: () => (this.isLoading = false),
            });
        this.subscriptions$.push(subscription$);
    }

    openModal(baseStringData?: BaseStringData): void {
        const newBaseString: BaseString = {
            id: undefined,
            active: true,
            author: undefined,
            group: undefined,
            page: undefined,
            identifier: '',
            sourceLanguage: undefined,
            translations: [],
        };
        const dialogRef = this.matDialog.open(ModalBaseStringComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: baseStringData ? baseStringData.baseString : newBaseString,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result: BaseString) => {
            if (result) {
                this.loadBaseStrings();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly BaseStringData[]): void {
        this.currentPageBaseStrings = $event;
    }

    canEdit(baseStringData: BaseStringData): boolean {
        const { baseString } = baseStringData;
        if (this.user.admin || !baseString || baseString.group.public || baseString.group.owner.id === this.user.id) {
            return true;
        }
        return baseString.group.permissions.some(
            (permission) => permission.user.id === this.user.id && permission.canWrite
        );
    }

    disable(baseStringData: BaseStringData): void {
        const { baseString } = baseStringData;
        const subscription$ = this.baseStringService.disable(baseString).subscribe({
            next: () => this.loadBaseStrings(),
            error: () => this.nzMessageService.create('error', 'Error disabling'),
        });
        this.subscriptions$.push(subscription$);
    }

    canDelete(baseStringData: BaseStringData): boolean {
        const { baseString } = baseStringData;
        return this.user.admin || baseString.group.owner.id === this.user.id;
    }

    showDeleteModal(baseStringData: BaseStringData): void {
        this.nzModalService.confirm({
            nzTitle: 'Are you sure delete this string?',
            nzOkText: 'Yes',
            nzOkType: 'primary',
            nzOkDanger: true,
            nzOnOk: () => this.delete(baseStringData),
            nzCancelText: 'No',
            nzAutofocus: 'cancel',
        });
    }

    private delete(baseStringData: BaseStringData): void {
        const { baseString } = baseStringData;
        const subscription$ = this.baseStringService.delete(baseString).subscribe((result) => {
            if (result) {
                this.loadBaseStrings();
                this.nzMessageService.create('success', `${baseString.identifier} has been deleted`);
            } else {
                this.nzMessageService.create('error', 'Error deleting');
            }
        });
        this.subscriptions$.push(subscription$);
    }

    onExpandChange(baseStringData: BaseStringData, expanded: boolean): void {
        baseStringData.expanded = expanded;
    }

    openReadModal(): void {
        const dialogRef = this.matDialog.open(ReadXliffComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            maxHeight: '700px',
        });
        const subscription$ = dialogRef.afterClosed().subscribe((baseStrings?: BaseString[]) => {
            if (baseStrings) {
                this.loadBaseStrings();
            }
        });
        this.subscriptions$.push(subscription$);
    }

    openCreateModal(): void {
        this.matDialog.open(CreateXliffComponent, {
            minWidth: '550px',
            maxWidth: '75%',
        });
    }
}
