export function checkNotNullParams(a: any, b: any): number {
    if (!a) return 1;
    if (!b) return -1;
    return 0;
}

export function sortStrings(a: string, b: string): number {
    const validParams = checkNotNullParams(a, b);
    return validParams === 0 ? a.localeCompare(b) : validParams;
}
