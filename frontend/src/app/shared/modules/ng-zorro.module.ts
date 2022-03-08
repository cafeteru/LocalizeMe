import { NgModule } from '@angular/core';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageServiceModule } from 'ng-zorro-antd/message';

@NgModule({
    imports: [
        NzButtonModule,
        NzFormModule,
        NzIconModule,
        NzInputModule,
        NzLayoutModule,
        NzMenuModule,
        NzMessageServiceModule,
        NzModalModule,
    ],
    exports: [
        NzButtonModule,
        NzFormModule,
        NzIconModule,
        NzInputModule,
        NzLayoutModule,
        NzMenuModule,
        NzMessageServiceModule,
        NzModalModule,
    ],
})
export class NgZorroModule {}
