/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface EnduroPreprocessing
 */
export interface EnduroPreprocessing {
    /**
     * 
     * @type {boolean}
     * @memberof EnduroPreprocessing
     */
    enabled: boolean;
    /**
     * 
     * @type {string}
     * @memberof EnduroPreprocessing
     */
    taskQueue: string;
    /**
     * 
     * @type {string}
     * @memberof EnduroPreprocessing
     */
    workflowName: string;
}

/**
 * Check if a given object implements the EnduroPreprocessing interface.
 */
export function instanceOfEnduroPreprocessing(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "enabled" in value;
    isInstance = isInstance && "taskQueue" in value;
    isInstance = isInstance && "workflowName" in value;

    return isInstance;
}

export function EnduroPreprocessingFromJSON(json: any): EnduroPreprocessing {
    return EnduroPreprocessingFromJSONTyped(json, false);
}

export function EnduroPreprocessingFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroPreprocessing {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'enabled': json['enabled'],
        'taskQueue': json['task_queue'],
        'workflowName': json['workflow_name'],
    };
}

export function EnduroPreprocessingToJSON(value?: EnduroPreprocessing | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'enabled': value.enabled,
        'task_queue': value.taskQueue,
        'workflow_name': value.workflowName,
    };
}

