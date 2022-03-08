import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NgZorroModule } from './modules/ng-zorro.module';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
    declarations: [],
    imports: [CommonModule, NgZorroModule, ReactiveFormsModule],
    exports: [NgZorroModule, ReactiveFormsModule],
})
export class SharedModule {}
