import { Component, OnInit } from '@angular/core';
import { LocalizeMeService } from './services/localize-me.service';
import { BaseComponent } from './components/base.component';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent extends BaseComponent implements OnInit {
    title = 'app';

    constructor(private localizeMeService: LocalizeMeService) {
        super();
    }

    ngOnInit(): void {
        super.ngOnInit();
        localStorage.setItem('isoCode', 'esp');
        const subscription$ = this.localizeMeService.login().subscribe();
        this.subscriptions$.push(subscription$);
    }
}
