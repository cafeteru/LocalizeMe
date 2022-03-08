import { Subscription } from 'rxjs';
import { Component, OnDestroy, OnInit } from '@angular/core';

@Component({
    template: '',
})
/**
 * Basic component that has the common properties
 */
export abstract class BaseComponent implements OnInit, OnDestroy {
    protected subscriptions: Subscription[] = [];

    ngOnInit(): void {
        this.subscriptions = [];
    }

    ngOnDestroy(): void {
        this.subscriptions.forEach((subscription) => subscription.unsubscribe());
    }
}
