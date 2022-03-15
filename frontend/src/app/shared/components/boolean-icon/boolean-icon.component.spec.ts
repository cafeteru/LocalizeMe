import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BooleanIconComponent } from './boolean-icon.component';
import { SharedModule } from '../../shared.module';

describe('BooleanIconComponent', () => {
    let component: BooleanIconComponent;
    let fixture: ComponentFixture<BooleanIconComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [BooleanIconComponent],
            imports: [SharedModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(BooleanIconComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
