import { LocalizeMePipe } from './localize-me.pipe';

describe('LocalizeMePipe', () => {
    it('create an instance', () => {
        const pipe = new LocalizeMePipe();
        expect(pipe).toBeTruthy();
    });
});
