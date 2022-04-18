export interface LanguageDto {
    description: string;
    isoCode: string;
}

export interface Language {
    active: boolean;
    description: string;
    id: string;
    isoCode: string;
}

export function createMockLanguage(): Language {
    return {
        active: true,
        id: '1',
        description: 'description',
        isoCode: 'isoCode',
    };
}
