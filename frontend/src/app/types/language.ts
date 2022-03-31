export interface LanguageDto {
    description: 'description';
    isoCode: 'isoCode';
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
