import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';

import { FormGroupUtil } from './form-group-util';

describe('FormGroupUtil', () => {
    let formGroup: UntypedFormGroup;
    const emailKey = 'email';
    const passwordKey = 'password';
    const noValidKey = 'noExist';
    const email = 'email@gmail.com';
    const errorKey = `There is no ${noValidKey} value inside the formGroup`;

    beforeEach(() => {
        formGroup = new UntypedFormGroup({
            email: new UntypedFormControl('', [Validators.required, Validators.email]),
            password: new UntypedFormControl('', [Validators.required]),
        });
    });

    it('check setValue with a key that exists', () => {
        expect(formGroup.get(emailKey).value).toBe('');
        FormGroupUtil.setValue(formGroup, emailKey, email);
        expect(formGroup.get(emailKey).value).toBe(email);
    });

    it('check setValue with a key that does not exist', () => {
        expect(formGroup.get(noValidKey)).toBeNull();
        expect(() => FormGroupUtil.setValue(formGroup, noValidKey, email)).toThrowError(errorKey);
    });

    it('Check setValue with an invalid formGroup', () => {
        expect(() => FormGroupUtil.setValue(null, noValidKey, email)).toThrowError(FormGroupUtil.errorFormGroup);
        expect(() => FormGroupUtil.setValue(undefined, noValidKey, email)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('check getValue with a key that exists', () => {
        expect(FormGroupUtil.getValue(formGroup, emailKey)).toBe('');
        FormGroupUtil.setValue(formGroup, emailKey, email);
        expect(FormGroupUtil.getValue(formGroup, emailKey)).toBe(email);
    });

    it('check getValue with a key that does not exist', () => {
        expect(() => FormGroupUtil.getValue(formGroup, noValidKey)).toThrowError(errorKey);
    });

    it('check getValue with a key that does not exist', () => {
        expect(() => FormGroupUtil.getValue(null, emailKey)).toThrowError(FormGroupUtil.errorFormGroup);
        expect(() => FormGroupUtil.getValue(undefined, emailKey)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('check getError with a key that exists', () => {
        expect(FormGroupUtil.getError(formGroup, emailKey)).toEqual(Object({ required: true }));
        FormGroupUtil.setValue(formGroup, emailKey, emailKey);
        expect(FormGroupUtil.getError(formGroup, emailKey)).not.toBeNull();
        expect(FormGroupUtil.getError(formGroup, emailKey)).toEqual(Object({ email: true }));
        FormGroupUtil.setValue(formGroup, emailKey, email);
        expect(FormGroupUtil.getError(formGroup, emailKey)).toBeNull();
    });

    it('check getError with a key that does not exist', () => {
        expect(() => FormGroupUtil.getError(formGroup, noValidKey)).toThrowError(errorKey);
    });

    it('check getError with an invalid formGroup', () => {
        expect(() => FormGroupUtil.getError(null, passwordKey)).toThrowError(FormGroupUtil.errorFormGroup);
        expect(() => FormGroupUtil.getError(undefined, emailKey)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('check checkError with a key that exists', () => {
        const key = 'email';
        const abstractControl = formGroup.get(key);
        abstractControl.markAllAsTouched();
        expect(FormGroupUtil.checkError(formGroup, key)).toBeTrue();
        abstractControl.markAllAsTouched();
        FormGroupUtil.setValue(formGroup, emailKey, emailKey);
        expect(FormGroupUtil.checkError(formGroup, key)).toBeTrue();
        abstractControl.markAllAsTouched();
        FormGroupUtil.setValue(formGroup, emailKey, email);
        expect(FormGroupUtil.checkError(formGroup, key)).toBeFalse();
    });

    it('check checkError with a key that does not exist', () => {
        expect(() => FormGroupUtil.checkError(formGroup, noValidKey)).toThrowError(errorKey);
    });

    it('check checkError with a null formGroup\n', () => {
        expect(() => FormGroupUtil.checkError(null, passwordKey)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('check checkError with a null formGroup', () => {
        expect(FormGroupUtil.valid(formGroup)).toBeFalse();
        FormGroupUtil.setValue(formGroup, emailKey, emailKey);
        expect(FormGroupUtil.valid(formGroup)).toBeFalse();
        FormGroupUtil.setValue(formGroup, passwordKey, emailKey);
        expect(FormGroupUtil.valid(formGroup)).toBeFalse();
        FormGroupUtil.setValue(formGroup, emailKey, email);
    });

    it('Check validFormGroup with invalid formGroup', () => {
        expect(() => FormGroupUtil.valid(null)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('check changeValidator with a key that exists', () => {
        expect(FormGroupUtil.getError(formGroup, passwordKey)).toEqual(Object({ required: true }));
        FormGroupUtil.setValue(formGroup, passwordKey, '1');
        expect(FormGroupUtil.getError(formGroup, passwordKey)).toBeNull();
        FormGroupUtil.changeValidator(formGroup, passwordKey, [Validators.minLength(2)], '1');
        expect(FormGroupUtil.getError(formGroup, passwordKey)).toEqual(
            Object({ minlength: Object({ requiredLength: 2, actualLength: 1 }) })
        );
        FormGroupUtil.setValue(formGroup, passwordKey, '');
        expect(FormGroupUtil.getError(formGroup, passwordKey)).toBeNull();
        FormGroupUtil.changeValidator(formGroup, passwordKey, [Validators.required, Validators.minLength(7)]);
        expect(FormGroupUtil.getError(formGroup, passwordKey)).toEqual(Object({ required: true }));
    });

    it('check changeValidator with a key that does not exist', () => {
        expect(() => FormGroupUtil.changeValidator(formGroup, noValidKey, [Validators.minLength(2)], '1')).toThrowError(
            errorKey
        );
    });

    it('Check changeValidator with a null formGroup', () => {
        expect(() => FormGroupUtil.changeValidator(null, emailKey, [Validators.minLength(2)], '1')).toThrowError(
            FormGroupUtil.errorFormGroup
        );
        expect(() => FormGroupUtil.changeValidator(undefined, passwordKey, [Validators.minLength(2)])).toThrowError(
            FormGroupUtil.errorFormGroup
        );
    });

    it('Check addFormControl with valid data', () => {
        expect(Object.keys(formGroup.controls).length).toBe(2);
        FormGroupUtil.addFormControl(formGroup, 'test', new UntypedFormControl(''));
        expect(Object.keys(formGroup.controls).length).toBe(3);
    });

    it('check addFormControl with invalid data', () => {
        expect(() => FormGroupUtil.addFormControl(formGroup, 'test', null)).toThrowError(
            FormGroupUtil.errorFormControl
        );
        expect(() => FormGroupUtil.addFormControl(undefined, 'test', new UntypedFormControl(''))).toThrowError(
            FormGroupUtil.errorFormGroup
        );
    });

    it('Check clean with invalid formGroup', () => {
        expect(() => FormGroupUtil.clean(null)).toThrowError(FormGroupUtil.errorFormGroup);
    });

    it('Check clean with valid formGroup', () => {
        FormGroupUtil.clean(formGroup);
        expect(formGroup.controls.password.value).toBe('');
        expect(formGroup.controls.email.value).toBe('');
    });

    it('Check deleteError with invalid formControl', () => {
        expect(() => FormGroupUtil.deleteError(null, '')).toThrowError(FormGroupUtil.errorFormControl);
    });

    it('Check deleteError with valid formControl and no exist error', () => {
        formGroup.controls.email.setValue('test');
        FormGroupUtil.deleteError(formGroup.controls.email, 'test');
        expect(formGroup.controls.email.errors.test).toBe(undefined);
    });

    it('Check deleteError with valid formControl and exist error', () => {
        formGroup.controls.email.setValue('test');
        FormGroupUtil.deleteError(formGroup.controls.email, 'email');
        expect(formGroup.controls.email.errors.email).toBe(undefined);
    });
});
