import { Component, EventEmitter, Input, OnChanges, OnInit, Output } from '@angular/core';
import { BaseComponent } from '../../../core/base/base.component';
import { Language } from '../../../types/language';
import { LanguageService } from '../../../core/services/language.service';

@Component({
    selector: 'app-language-finder',
    templateUrl: './language-finder.component.html',
    styleUrls: ['./language-finder.component.scss'],
})
export class LanguageFinderComponent extends BaseComponent implements OnInit, OnChanges {
    isLoading = false;
    options: string[] = [];
    selectedText: string;
    originalLanguages: readonly Language[] = [];
    languages: readonly Language[] = [];
    @Input() valid = false;
    @Input() title = 'Languages';
    @Input() selectLanguage: Language;
    @Input() filterLanguage: Language;
    @Output() emitter: EventEmitter<Language> = new EventEmitter<Language>();

    constructor(private languageService: LanguageService) {
        super();
    }

    ngOnInit() {
        super.ngOnInit();
        this.isLoading = true;
        if (this.selectLanguage) {
            this.selectedText = this.getName(this.selectLanguage);
        }
        const subscription$ = this.languageService
            .findAll()
            .pipe()
            .subscribe({
                next: (languages) => {
                    this.originalLanguages = languages.filter((language) => language.active);
                    this.languages = [...this.originalLanguages];
                },
                error: () => {
                    this.languages = [];
                    this.isLoading = false;
                },
                complete: () => (this.isLoading = false),
            });
        this.subscriptions$.push(subscription$);
    }

    ngOnChanges(): void {
        this.languages = [...this.originalLanguages];
        this.searchGroupByName('');
        if (this.filterLanguage) {
            this.languages = [...this.originalLanguages].filter((language) => language.id != this.filterLanguage.id);
            this.options = this.languages.map((language) => this.getName(language));
        }
    }

    getName(language: Language): string {
        if (!language) {
            return '';
        }
        return `${language.isoCode} - ${language.description}`;
    }

    searchGroupByName(value: string): void {
        const strings = this.languages.map((language) => this.getName(language));
        this.options = value ? strings.filter((name) => name.includes(value)) : strings;
    }

    add(): void {
        if (this.selectedText) {
            const languages = this.languages.filter((value) => this.selectedText.includes(value.isoCode));
            this.selectLanguage = languages[0];
            this.emitter.emit(this.selectLanguage);
        } else {
            this.emitter.emit(undefined);
            this.searchGroupByName('');
        }
    }
}
