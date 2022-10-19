import { UntypedFormGroup, ValidatorFn } from '@angular/forms';
import { FormGroupUtil } from '../../shared/utils/form-group-util';

export class IsSameValidator {
    static isValid(name1: string, name2: string): ValidatorFn {
        return (formGroup: UntypedFormGroup) => {
            const control1 = formGroup.get(name1);
            const control2 = formGroup.get(name2);
            if (control1 && control2) {
                FormGroupUtil.deleteError(formGroup, name2);
                const value1 = control1.value;
                const value2 = control2.value;
                if (value1 !== value2) {
                    control2.setErrors({ confirm: true });
                    control2.markAsTouched({ onlySelf: true });
                }
            }
            return null;
        };
    }
}
