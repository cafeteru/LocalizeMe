import { NgModule } from '@angular/core';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

@NgModule({
    imports: [MatDialogModule, MatIconModule, MatInputModule],
    exports: [MatDialogModule, MatIconModule, MatInputModule],
})
export class MaterialModule {}
