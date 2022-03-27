export interface Stage {
    ID: string;
    Name: string;
    Active: boolean;
}

export function createMockStage(): Stage {
    return {
        Active: true,
        ID: '1',
        Name: 'name',
    };
}
