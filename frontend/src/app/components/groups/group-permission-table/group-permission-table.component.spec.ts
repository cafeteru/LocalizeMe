import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupPermissionTableComponent } from './group-permission-table.component';
import { SharedModule } from '../../../shared/shared.module';
import { CoreModule } from '../../../core/core.module';

describe('GroupPermissionTableComponent', () => {
    let component: GroupPermissionTableComponent;
    let fixture: ComponentFixture<GroupPermissionTableComponent>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [GroupPermissionTableComponent],
            imports: [SharedModule, CoreModule],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(GroupPermissionTableComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
