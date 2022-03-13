import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NgZorroModule } from './modules/ng-zorro.module';
import { ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './modules/material.module';

@NgModule({
    declarations: [],
    imports: [CommonModule, NgZorroModule, ReactiveFormsModule, MaterialModule],
    exports: [NgZorroModule, ReactiveFormsModule, MaterialModule],
})
export class SharedModule {}
