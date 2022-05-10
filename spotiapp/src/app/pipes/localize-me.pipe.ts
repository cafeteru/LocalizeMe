import { Pipe, PipeTransform } from '@angular/core';
import { BaseComponent } from '../components/base.component';
import { LocalizeMeService } from '../services/localize-me.service';
import { Observable } from 'rxjs';

@Pipe({
    name: 'localizeMe',
})
export class LocalizeMePipe extends BaseComponent implements PipeTransform {
    constructor(private localizeMeService: LocalizeMeService) {
        super();
    }

    transform(identifier: string, isoCode: string): Observable<string> {
        return this.localizeMeService.findByIdentifierAndLanguage(identifier, isoCode);
    }
}
