import { Language } from './language';
import { User } from './user';
import { Stage } from './stage';

export class Translation {
    content: string;
    language: Language;
    version: number;
    active: boolean;
    author: User;
    date: any;
    stage: Stage;
}
