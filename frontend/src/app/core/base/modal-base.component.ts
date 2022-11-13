import { Component } from '@angular/core';
import { BaseComponent } from './base.component';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
    template: '',
})
/**
 * Basic component that has the common properties
 */
export abstract class ModalBaseComponent extends BaseComponent {

    protected constructor(
      protected nzMessageService: NzMessageService,
    ) {
        super();
    }


    createMessage(type: string, message: string): void {
        this.nzMessageService.create(type, message);
    }
}
