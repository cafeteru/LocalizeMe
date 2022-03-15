import { NzTableSortFn, NzTableSortOrder } from 'ng-zorro-antd/table';

export interface ColumnHeader<T> {
    name: string;
    sortDirections: NzTableSortOrder[];
    sortFn: NzTableSortFn<T>;
    sortOrder: NzTableSortOrder;
}

export const sortDirections: NzTableSortOrder[] = ['ascend', 'descend', null];
