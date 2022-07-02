import { Stage } from './stage';

export class XliffDto {
    baseStringIds: string[];
    sourceLanguageId: string;
    stage: Stage;
    targetLanguageId: string;
}
