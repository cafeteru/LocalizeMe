import { IsSameValidator } from './is-same-validator';
import { UntypedFormControl, UntypedFormGroup } from '@angular/forms';

describe('IsSameValidator', () => {
    let formGroup: UntypedFormGroup;

    beforeEach(() => {
        formGroup = new UntypedFormGroup(
            {
                'a': new UntypedFormControl(''),
                'b': new UntypedFormControl(''),
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
