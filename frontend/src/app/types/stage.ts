export interface StageDto {
    name: string;
}

export interface Stage {
    id: string;
    name: string;
    active: boolean;
}

export function createMockStage(): Stage {
    return {
        active: true,
        id: '1',
        name: 'name',
    };
}
