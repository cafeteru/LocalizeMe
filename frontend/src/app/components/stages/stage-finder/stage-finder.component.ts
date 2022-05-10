import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { Stage } from '../../../types/stage';
import { StageService } from '../../../core/services/stage.service';

@Component({
    selector: 'app-stage-finder',
    templateUrl: './stage-finder.component.html',
    styleUrls: ['./stage-finder.component.scss'],
})
export class StageFinderComponent extends BaseComponent implements OnInit {
    isLoading = false;
    options: string[] = [];
    selectedText: string;
    stages: readonly Stage[] = [];
    @Input() valid = false;
    @Input() selectedStage: Stage;
    @Output() emitter: EventEmitter<Stage> = new EventEmitter<Stage>();

    constructor(private stageService: StageService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.isLoading = true;
        if (this.selectedStage) {
            this.selectedText = this.selectedStage.name;
        }
        const subscription$ = this.stageService.findAll().subscribe({
            next: (stages) => (this.stages = stages.filter((stage) => stage.active)),
            error: () => {
                this.stages = [];
                this.isLoading = false;
            },
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    searchStageByName(value: string): void {
        const strings = this.stages.map((stage) => stage.name);
        this.options = value
            ? this.stages
                  .map((stage) => stage.name)
                  .filter((name) => name.toLocaleLowerCase().includes(value.toLocaleLowerCase()))
            : strings;
    }

    add(): void {
        if (this.selectedText) {
            const stages = this.stages.filter((value) => this.selectedText.includes(value.name));
            this.selectedStage = stages[0];
            this.emitter.emit(this.selectedStage);
        } else {
            this.emitter.emit(undefined);
        }
    }
}
