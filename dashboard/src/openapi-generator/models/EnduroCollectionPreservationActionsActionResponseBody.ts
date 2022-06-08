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
 * PreservationAction describes a preservation action. (default view)
 * @export
 * @interface EnduroCollectionPreservationActionsActionResponseBody
 */
export interface EnduroCollectionPreservationActionsActionResponseBody {
    /**
     * 
     * @type {string}
     * @memberof EnduroCollectionPreservationActionsActionResponseBody
     */
    actionId: string;
    /**
     * 
     * @type {number}
     * @memberof EnduroCollectionPreservationActionsActionResponseBody
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof EnduroCollectionPreservationActionsActionResponseBody
     */
    name: string;
    /**
     * 
     * @type {Date}
     * @memberof EnduroCollectionPreservationActionsActionResponseBody
     */
    startedAt: Date;
    /**
     * 
     * @type {string}
     * @memberof EnduroCollectionPreservationActionsActionResponseBody
     */
    status: EnduroCollectionPreservationActionsActionResponseBodyStatusEnum;
}


/**
 * @export
 */
export const EnduroCollectionPreservationActionsActionResponseBodyStatusEnum = {
    Unspecified: 'unspecified',
    Complete: 'complete',
    Processing: 'processing',
    Failed: 'failed'
} as const;
export type EnduroCollectionPreservationActionsActionResponseBodyStatusEnum = typeof EnduroCollectionPreservationActionsActionResponseBodyStatusEnum[keyof typeof EnduroCollectionPreservationActionsActionResponseBodyStatusEnum];


export function EnduroCollectionPreservationActionsActionResponseBodyFromJSON(json: any): EnduroCollectionPreservationActionsActionResponseBody {
    return EnduroCollectionPreservationActionsActionResponseBodyFromJSONTyped(json, false);
}

export function EnduroCollectionPreservationActionsActionResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroCollectionPreservationActionsActionResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'actionId': json['action_id'],
        'id': json['id'],
        'name': json['name'],
        'startedAt': (new Date(json['started_at'])),
        'status': json['status'],
    };
}

export function EnduroCollectionPreservationActionsActionResponseBodyToJSON(value?: EnduroCollectionPreservationActionsActionResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'action_id': value.actionId,
        'id': value.id,
        'name': value.name,
        'started_at': (value.startedAt.toISOString()),
        'status': value.status,
    };
}

