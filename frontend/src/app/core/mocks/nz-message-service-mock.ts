import { NzMessageDataOptions, NzMessageRef } from 'ng-zorro-antd/message/typings';

export const nzMessageServiceMock = {
    create(content: string, options?: NzMessageDataOptions): NzMessageRef {
        return undefined;
    },
};
