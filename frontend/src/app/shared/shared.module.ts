import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NgZorroModule } from './modules/ng-zorro.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MaterialModule } from './modules/material.module';
import { BooleanIconComponent } from './components/boolean-icon/boolean-icon.component';

@NgModule({
    declarations: [BooleanIconComponent],
    imports: [CommonModule, NgZorroModule, ReactiveFormsModule, MaterialModule, FormsModule],
    exports: [NgZorroModule, ReactiveFormsModule, MaterialModule, BooleanIconComponent, FormsModule],
})
export class SharedModule {}
