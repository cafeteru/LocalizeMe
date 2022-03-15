import { HttpContext, HttpHeaders, HttpParams } from '@angular/common/http';

interface DefaultHttpOptions {
    headers?:
        | HttpHeaders
        | {
              [header: string]: string | string[];
          };
    context?: HttpContext;
    observe?: 'body';
    params?:
        | HttpParams
        | {
              [param: string]: string | number | boolean | ReadonlyArray<string | number | boolean>;
          };
    reportProgress?: boolean;
    responseType?: 'json';
    withCredentials?: boolean;
}

export const getDefaultHttpOptions = (): DefaultHttpOptions => {
    return {
        responseType: 'json',
        headers: new HttpHeaders({
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.Authorization}`,
        }),
    };
};
