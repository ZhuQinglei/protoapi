/**
* This file is generated by 'protoapi'
* The file contains helper functions that would be used in generated api file, usually in './api.ts' or './xxxService.ts'
* The generated code is written in TypeScript
* -------------------------------------------
* 该文件生成于protoapi
* 文件包含一些函数协助生成的前端调用API
* 文件内代码使用TypeScript
*/

import { CommonError, GenericError, AuthError, ValidateError, BindError } from './AppServiceObjs'


export type CommonErrorType = GenericError | AuthError | BindError | ValidateError
/**
 * Defined Http Code for response handling
 */
export enum httpCode {
    DEFAULT = 0,
    NORMAL = 200,
    BIZ_ERROR = 400,
    COMMON_ERROR = 420,
    INTERNAL_ERROR = 500,
}
/**
 *
 * @param {CommonError} commonErr the error object
 */
export function mapCommonErrorType(commonErr: CommonError): CommonErrorType | null {
    console.log(commonErr)
    for (let key in commonErr) {
        if (commonErr.hasOwnProperty(key) && commonErr[key]) {
            switch (key) {
                case 'genericError':
                    return commonErr[key] as GenericError
                case 'authError':
                    return commonErr[key] as AuthError
                case 'validateError':
                    return commonErr[key] as ValidateError
                case 'bindError':
                    return commonErr[key] as BindError
                default:
                    return null
            }
  
        }
    }
    return null
}
/**
 *
 * @param {response} response the error response
 */
export function errorHandling(err): Promise<never> {
    if(err.response === undefined) {
        throw err;
    }
    let data;
    try {
        data = JSON.parse(err.response.data);
    } catch (err) {
        data = err.response.data;
    }
    switch (err.response.status) {
        case httpCode.BIZ_ERROR:
            return Promise.reject({...err, message: data.message});

        case httpCode.COMMON_ERROR:
            let returnErr = mapCommonErrorType(data);
            if(!returnErr){
                throw data
            }
            var result
            switch (returnErr.kind) {
                case "validateError": 
                    result = {...err.response, message: 'validate error', errors: returnErr.errors}
                    break
                default:
                    result = {...err.response, message: returnErr.message}
                    break
            }
            return Promise.reject(result);

    }
    throw data;
}

/**
 *
 * @param val a string
 * @returns an encoded string that can be append to api url
 */
export function encode(val: string): string {
    return encodeURIComponent(val).
        replace(/%40/gi, '@').
        replace(/%3A/gi, ':').
        replace(/%24/g, '$').
        replace(/%2C/gi, ',').
        replace(/%20/g, '+').
        replace(/%5B/gi, '[').
        replace(/%5D/gi, ']');
}

/**
 * Build a URL by appending params to the end
 * @param url : the base url for the service
 * @param params : the request object. e.g. for HelloRequest would be the object of type HelloRequest
 * @returns: returns a full Url string - for GET by key/value pairs
 * @example:
 * baseUrl = "http://localhost:8080"
 * arg = {name: "wengwei", nick: "wentian"}
 * returns => http://localhost:8080?name="wengwei"&nick="wentian"
 */
export function generateQueryUrl<T>(url: string, params: T): string {
    if (!params) {
        return url;
    }

    let parts: string[] = [];


    for (let key in params) {
        let val;
        if (Object.prototype.hasOwnProperty(key)) {
            val = params[key];
        }

        if (val === null || typeof val === 'undefined') {
            return '';
        }

        let k, vals;
        // if is array
        if (val.toString() === '[object Array]') {
            k = key + '[]';
        } else {
            k = key
            vals = [val];
        }

        vals.forEach(v => {
            // if is date
            if (v.toString() === '[object File]') {
                v = v.toISOString();
                // if is object
            } else if (typeof v === 'object') {
                v = JSON.stringify(v);
            }
            parts.push(encode(k) + '=' + encode(v))
        });
    }
    let serializedParams = parts.join('&');

    if (serializedParams) {
        url += (url.indexOf('?') === -1 ? '?' : '&') + serializedParams;
    }
    return url
}

/**
 *
 * @param url the base url for the service
 * @param serviceName the service name
 * @param functionName the function name
 * @example
 * baseUrl = "http://localhost:8080"
 * serviceName = "HelloService"
 * functionName = "SayHello"
 * returns => http://localhost:8080/HelloService.SayHello
 */
export function generateUrl<T>(url: string, serviceName: string, functionName: string): string {
    return url + "/" + serviceName + "." + functionName;
}
