import { ConfirmType, ModalOptions } from 'ng-zorro-antd/modal/modal-types';
import { NzModalRef } from 'ng-zorro-antd/modal/modal-ref';

export const mockNzModalService = {
    confirm<T>(options?: ModalOptions<T>, confirmType?: ConfirmType): NzModalRef<T> {
        return undefined;
    },
};
