import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../../core/base/base.component';
import { BaseString } from '../../../../types/base-string';
import { ColumnHeader, sortDirections } from '../../../../shared/components/utils/nz-table-utils';
import {
    sortBaseStringByActive,
    sortBaseStringByAuthor,
    sortBaseStringByGroup,
    sortBaseStringByIdentifier,
    sortBaseStringBySourceLanguage,
} from '../../../../shared/sorts/base-string-sorts';
import { BaseStringService } from '../../../../core/services/base-string.service';
import { map } from 'rxjs';
import { Language } from '../../../../types/language';

interface BaseStringData {
    baseString: BaseString;
    expanded: boolean;
    selected: boolean;
}

@Component({
    selector: 'app-create-xliff-list-base-strings',
    templateUrl: './create-xliff-list-base-strings.component.html',
    styleUrls: ['./create-xliff-list-base-strings.component.scss'],
})
export class CreateXliffListBaseStringsComponent extends BaseComponent implements OnInit {
    currentPageBaseStrings: readonly BaseStringData[] = [];
    isLoading = false;
    baseStrings: BaseStringData[];
    allSelected: boolean = false;
    selectedLength: number = 0;
    @Output() emitter: EventEmitter<BaseString[]> = new EventEmitter<BaseString[]>();

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

    constructor(private baseStringService: BaseStringService) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.loadBaseStrings();
    }

    loadBaseStrings(): void {
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
                            selected: false,
                        };
                        return baseStringData;
                    })
                )
            )
            .subscribe({
                next: (baseStrings) => {
                    this.baseStrings = baseStrings;
                },
                error: () => (this.isLoading = false),
                complete: () => (this.isLoading = false),
            });
        this.subscriptions$.push(subscription$);
    }

    onCurrentPageDataChange($event: readonly BaseStringData[]): void {
        this.currentPageBaseStrings = $event;
    }

    markAll(selected: boolean): void {
        this.allSelected = selected;
        this.baseStrings.forEach((baseString) => (baseString.selected = selected));
        this.setSelected();
    }

    markPage(selected: boolean): void {
        this.currentPageBaseStrings.forEach((bookingRow) => (bookingRow.selected = selected));
        this.updateAllSelected();
    }

    setSelected(): void {
        const filter = this.baseStrings.filter((baseString) => baseString.selected === true);
        this.selectedLength = filter.length;
        this.emitter.emit(filter.map((baseString) => baseString.baseString));
    }

    onChangeSelected(baseStringData: BaseStringData, selected: boolean): void {
        baseStringData.selected = selected;
        this.updateAllSelected();
    }

    private updateAllSelected(): void {
        this.allSelected = false;
        this.allSelected = this.baseStrings.every((bookingRow) => bookingRow.selected);
        this.setSelected();
    }
}
