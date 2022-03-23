import { Component, OnInit } from '@angular/core';
import { User } from '../../../types/user';
import { UpdateUserComponent, UpdateUserData } from '../../users/update-user/update-user.component';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CreateStageComponent } from '../create-stage/create-stage.component';

@Component({
    selector: 'app-stage-list',
    templateUrl: './stage-list.component.html',
    styleUrls: ['./stage-list.component.scss'],
})
export class StageListComponent extends BaseComponent implements OnInit {
    constructor(public dialog: MatDialog, private messageService: NzMessageService) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
    }

    openModal(): void {
        const dialogRef = this.dialog.open(CreateStageComponent, {
            minWidth: '550px',
            maxWidth: '75%',
        });
        const subscription = dialogRef.afterClosed().subscribe();
        this.subscriptions.push(subscription);
    }
}
