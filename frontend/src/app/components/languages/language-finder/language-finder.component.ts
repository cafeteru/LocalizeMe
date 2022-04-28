import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { Language } from '../../../types/language';
import { LanguageService } from '../../../core/services/language.service';

@Component({
    selector: 'app-language-finder',
    templateUrl: './language-finder.component.html',
    styleUrls: ['./language-finder.component.scss'],
})
export class LanguageFinderComponent extends BaseComponent implements OnInit {
    isLoading = false;
    options: string[] = [];
    selectedLanguageText: string;
    languages: readonly Language[] = [];
    @Input() valid: boolean;
    @Output() emitter: EventEmitter<Language> = new EventEmitter<Language>();

    constructor(private languageService: LanguageService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.isLoading = true;
        const subscription$ = this.languageService.findAll().subscribe({
            next: (languages) => (this.languages = languages.filter((language) => language.active)),
            error: () => {
                this.languages = [];
                this.isLoading = false;
            },
            complete: () => (this.isLoading = false),
        });
        this.subscriptions$.push(subscription$);
    }

    searchGroupByName(value: string): void {
        const strings = this.languages.map((language) => `${language.isoCode} - ${language.description}`);
        this.options = value ? strings.filter((name) => name.includes(value)) : strings;
    }

    add(): void {
        const language = this.languages.filter((value) => this.selectedLanguageText.includes(value.isoCode));
        this.emitter.emit(language[0]);
    }
}
