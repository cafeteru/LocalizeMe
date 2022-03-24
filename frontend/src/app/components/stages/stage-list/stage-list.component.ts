import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CreateStageComponent } from '../create-stage/create-stage.component';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { Stage } from '../../../types/stage';
import { sortActive, sortName } from '../../../shared/sorts/stages-sorts';
import { StageService } from '../../../core/services/stage.service';

@Component({
    selector: 'app-stage-list',
    templateUrl: './stage-list.component.html',
    styleUrls: ['./stage-list.component.scss'],
})
export class StageListComponent extends BaseComponent implements OnInit {
    currentPageStages: readonly Stage[] = [];
    stages: readonly Stage[] = [];

    listOfColumns: ColumnHeader<Stage>[] = [
        {
            name: 'Email',
            sortOrder: null,
            sortFn: sortName,
            sortDirections,
        },
        {
            name: 'Active',
            sortOrder: null,
            sortFn: sortActive,
            sortDirections,
        },
    ];

    constructor(
        private stageService: StageService,
        public dialog: MatDialog,
        private messageService: NzMessageService
    ) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        const subscription = this.stageService.findAll().subscribe({
            next: (stages) => (this.stages = stages),
            error: () => (this.stages = []),
        });
        this.subscriptions.push(subscription);
    }

    onCurrentPageDataChange($event: readonly Stage[]): void {
        this.currentPageStages = $event;
    }

    openModal(): void {
        const dialogRef = this.dialog.open(CreateStageComponent, {
            minWidth: '550px',
            maxWidth: '75%',
        });
        const subscription = dialogRef.afterClosed().subscribe();
        this.subscriptions.push(subscription);
    }

    showDeleteModal(user: any) {}

    disable(user: any) {}

    openUpdate(user: any) {}
}
