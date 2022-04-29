import { createMockLanguage, Language } from './language';
import { createMockUser, User } from './user';
import { createMockStage, Stage } from './stage';

export class Translation {
    active: boolean;
    author: User;
    content: string;
    date: any;
    language: Language;
    stage: Stage;
    version: number;
}

export function createMockBaseTranslation(): Translation {
    return {
        active: true,
        author: createMockUser(),
        content: '',
        date: new Date(),
        language: createMockLanguage(),
        stage: createMockStage(),
        version: 1,
    };
}
