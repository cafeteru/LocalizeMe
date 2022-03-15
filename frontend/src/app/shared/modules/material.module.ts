import { NgModule } from '@angular/core';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
    imports: [MatDialogModule, MatIconModule],
    exports: [MatDialogModule, MatIconModule],
})
export class MaterialModule {}
