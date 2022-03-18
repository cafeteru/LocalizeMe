import { IsSameValidator } from './is-same-validator';
import { FormControl, FormGroup } from '@angular/forms';

describe('IsSameValidator', () => {
    let formGroup: FormGroup;

    beforeEach(() => {
        formGroup = new FormGroup(
            {
                'a': new FormControl(''),
                'b': new FormControl(''),
            },
            { validators: IsSameValidator.isValid('a', 'b') }
        );
    });

    it('check isValid', () => {
        expect(formGroup.valid).toBeTrue();
        formGroup.get('a').setValue('a');
        expect(formGroup.valid).toBeFalse();
    });
});
