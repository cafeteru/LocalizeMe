import { createMockLanguage, Language } from './language';
import { createMockUser, User } from './user';
import { createMockGroup, Group } from './group';
import { Translation } from './translation';

export class BaseString {
    id: string;
    sourceLanguage: Language;
    identifier: string;
    page: string;
    group: Group;
    active: boolean;
    author: User;
    translations: Translation[];
}

export function createMockBaseString(): BaseString {
    return {
        id: '1',
        sourceLanguage: createMockLanguage(),
        identifier: 'identifier',
        group: createMockGroup(),
        page: 'page',
        active: true,
        author: createMockUser(),
        translations: [],
    };
}
