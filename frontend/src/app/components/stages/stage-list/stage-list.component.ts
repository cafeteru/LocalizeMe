import { Component, OnInit } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { MatDialog } from '@angular/material/dialog';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ModalStageComponent } from '../modal-stage/modal-stage.component';
import { ColumnHeader, sortDirections } from '../../../shared/components/utils/nz-table-utils';
import { Stage } from '../../../types/stage';
import { sortActive, sortName } from '../../../shared/sorts/stages-sorts';
import { StageService } from '../../../core/services/stage.service';
import { tap } from 'rxjs';
import { User } from '../../../types/user';

@Component({
    selector: 'app-stage-list',
    templateUrl: './stage-list.component.html',
    styleUrls: ['./stage-list.component.scss'],
})
export class StageListComponent extends BaseComponent implements OnInit {
    currentPageStages: readonly Stage[] = [];
    stages: readonly Stage[] = [];
    isLoading = false;

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
        this.loadStages();
    }

    loadStages(): void {
        const subscription = this.stageService
            .findAll()
            .pipe(tap(() => (this.isLoading = true)))
            .subscribe({
                next: (stages) => {
                    this.stages = stages;
                    this.isLoading = false;
                },
                error: () => {
                    this.stages = [];
                    this.isLoading = false;
                },
            });
        this.subscriptions.push(subscription);
    }

    onCurrentPageDataChange($event: readonly Stage[]): void {
        this.currentPageStages = $event;
    }

    openModal(stage?: Stage): void {
        const newStage: Stage = {
            Name: undefined,
            Active: true,
            ID: undefined,
        };
        const dialogRef = this.dialog.open(ModalStageComponent, {
            minWidth: '550px',
            maxWidth: '75%',
            data: stage ? stage : newStage,
        });
        const subscription = dialogRef.afterClosed().subscribe((result: Stage) => {
            if (result) {
                this.loadStages();
            }
        });
        this.subscriptions.push(subscription);
    }

    disable(stage: Stage): void {
        const subscription = this.stageService.disable(stage).subscribe((result) => this.loadStages());
        this.subscriptions.push(subscription);
    }

    showDeleteModal(user: any) {}
}
