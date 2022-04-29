import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { createMockGroup, Group } from '../../../types/group';
import { GroupService } from '../../../core/services/group.service';

@Component({
    selector: 'app-group-finder',
    templateUrl: './group-finder.component.html',
    styleUrls: ['./group-finder.component.scss'],
})
export class GroupFinderComponent extends BaseComponent implements OnInit {
    isLoading = false;
    options: string[] = [];
    selectedText: string;
    groups: readonly Group[] = [];
    @Input() selectGroup: Group = createMockGroup();
    @Output() emitter: EventEmitter<Group> = new EventEmitter<Group>();

    constructor(private groupService: GroupService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.isLoading = true;
        if (this.selectGroup) {
            this.selectedText = this.selectGroup.name;
        }
        const subscription$ = this.groupService.findAll().subscribe({
            next: (groups) => (this.groups = groups.filter((group) => group.active)),
            error: () => {
                this.groups = [];
                this.isLoading = false;
            },
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    searchGroupByName(value: string): void {
        const names = this.groups.map((group) => group.name);
        this.options = value ? names.filter((name) => name.includes(value)) : names;
    }

    add(): void {
        if (this.selectedText) {
            const group = this.groups.filter((value) => value.name.includes(this.selectedText));
            this.selectGroup = group[0];
            this.emitter.emit(this.selectGroup);
        } else {
            this.emitter.emit(undefined);
        }
    }
}
