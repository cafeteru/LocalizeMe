import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { ModalLanguageComponent } from '../../languages/modal-language/modal-language.component';
import { ModalGroupComponent } from '../modal-group/modal-group.component';
import { Language } from '../../../types/language';
import { Group } from '../../../types/group';

@Component({
    selector: 'app-group-list',
    templateUrl: './group-list.component.html',
    styleUrls: ['./group-list.component.scss'],
})
export class GroupListComponent extends BaseComponent implements OnInit {
    isLoading = false;

    constructor(public matDialog: MatDialog) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
    }

    openModal(): void {
        const newGroup: Group = {
            id: undefined,
            active: true,
            name: undefined,
            permissions: [],
            owner: undefined,
        };
        const dialogRef = this.matDialog.open(ModalGroupComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: newGroup,
        });
        const subscription$ = dialogRef.afterClosed().subscribe((result) => {
            if (result) {
            }
        });
        this.subscriptions$.push(subscription$);
    }
}
