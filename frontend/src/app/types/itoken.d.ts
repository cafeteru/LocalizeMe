export interface IToken {
    id: string;
    email: string;
    exp: number;
    active: boolean;
    admin: boolean;
}
