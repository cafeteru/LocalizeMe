import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { Group } from '../../../types/group';
import { GroupService } from '../../../core/services/group.service';

@Component({
    selector: 'app-group-finder',
    templateUrl: './group-finder.component.html',
    styleUrls: ['./group-finder.component.scss'],
})
export class GroupFinderComponent extends BaseComponent implements OnInit {
    isLoading = false;
    options: string[] = [];
    selectedGroupName: string;
    groups: readonly Group[] = [];
    @Output() emitter: EventEmitter<Group> = new EventEmitter<Group>();

    constructor(private groupService: GroupService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.isLoading = true;
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
        const group = this.groups.filter((value) => value.name.includes(this.selectedGroupName));
        this.emitter.emit(group[0]);
    }
}
