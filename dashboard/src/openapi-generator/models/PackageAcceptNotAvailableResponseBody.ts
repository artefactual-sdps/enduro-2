/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * accept_not_available_response_body result type (default view)
 * @export
 * @interface PackageAcceptNotAvailableResponseBody
 */
export interface PackageAcceptNotAvailableResponseBody {
    /**
     * Is the error a server-side fault?
     * @type {boolean}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    fault: boolean;
    /**
     * ID is a unique identifier for this particular occurrence of the problem.
     * @type {string}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    id: string;
    /**
     * Message is a human-readable explanation specific to this occurrence of the problem.
     * @type {string}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    message: string;
    /**
     * Name is the name of this class of errors.
     * @type {string}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    name: string;
    /**
     * Is the error temporary?
     * @type {boolean}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    temporary: boolean;
    /**
     * Is the error a timeout?
     * @type {boolean}
     * @memberof PackageAcceptNotAvailableResponseBody
     */
    timeout: boolean;
}

export function PackageAcceptNotAvailableResponseBodyFromJSON(json: any): PackageAcceptNotAvailableResponseBody {
    return PackageAcceptNotAvailableResponseBodyFromJSONTyped(json, false);
}

export function PackageAcceptNotAvailableResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackageAcceptNotAvailableResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fault': json['fault'],
        'id': json['id'],
        'message': json['message'],
        'name': json['name'],
        'temporary': json['temporary'],
        'timeout': json['timeout'],
    };
}

export function PackageAcceptNotAvailableResponseBodyToJSON(value?: PackageAcceptNotAvailableResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fault': value.fault,
        'id': value.id,
        'message': value.message,
        'name': value.name,
        'temporary': value.temporary,
        'timeout': value.timeout,
    };
}

