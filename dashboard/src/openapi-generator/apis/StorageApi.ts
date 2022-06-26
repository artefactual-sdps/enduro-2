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


import * as runtime from '../runtime';
import {
    StorageDownloadNotFoundResponseBody,
    StorageDownloadNotFoundResponseBodyFromJSON,
    StorageDownloadNotFoundResponseBodyToJSON,
    StorageMoveNotAvailableResponseBody,
    StorageMoveNotAvailableResponseBodyFromJSON,
    StorageMoveNotAvailableResponseBodyToJSON,
    StorageMoveNotFoundResponseBody,
    StorageMoveNotFoundResponseBodyFromJSON,
    StorageMoveNotFoundResponseBodyToJSON,
    StorageMoveNotValidResponseBody,
    StorageMoveNotValidResponseBodyFromJSON,
    StorageMoveNotValidResponseBodyToJSON,
    StorageMoveRequestBody,
    StorageMoveRequestBodyFromJSON,
    StorageMoveRequestBodyToJSON,
    StorageMoveStatusNotFoundResponseBody,
    StorageMoveStatusNotFoundResponseBodyFromJSON,
    StorageMoveStatusNotFoundResponseBodyToJSON,
    StorageMoveStatusResponseBody,
    StorageMoveStatusResponseBodyFromJSON,
    StorageMoveStatusResponseBodyToJSON,
    StorageRejectNotAvailableResponseBody,
    StorageRejectNotAvailableResponseBodyFromJSON,
    StorageRejectNotAvailableResponseBodyToJSON,
    StorageRejectNotFoundResponseBody,
    StorageRejectNotFoundResponseBodyFromJSON,
    StorageRejectNotFoundResponseBodyToJSON,
    StorageRejectNotValidResponseBody,
    StorageRejectNotValidResponseBodyFromJSON,
    StorageRejectNotValidResponseBodyToJSON,
    StorageSubmitNotAvailableResponseBody,
    StorageSubmitNotAvailableResponseBodyFromJSON,
    StorageSubmitNotAvailableResponseBodyToJSON,
    StorageSubmitNotValidResponseBody,
    StorageSubmitNotValidResponseBodyFromJSON,
    StorageSubmitNotValidResponseBodyToJSON,
    StorageSubmitRequestBody,
    StorageSubmitRequestBodyFromJSON,
    StorageSubmitRequestBodyToJSON,
    StorageSubmitResponseBody,
    StorageSubmitResponseBodyFromJSON,
    StorageSubmitResponseBodyToJSON,
    StorageUpdateNotAvailableResponseBody,
    StorageUpdateNotAvailableResponseBodyFromJSON,
    StorageUpdateNotAvailableResponseBodyToJSON,
    StorageUpdateNotValidResponseBody,
    StorageUpdateNotValidResponseBodyFromJSON,
    StorageUpdateNotValidResponseBodyToJSON,
    StoredLocationResponse,
    StoredLocationResponseFromJSON,
    StoredLocationResponseToJSON,
} from '../models';

export interface StorageDownloadRequest {
    aipId: string;
}

export interface StorageMoveRequest {
    aipId: string;
    moveRequestBody: StorageMoveRequestBody;
}

export interface StorageMoveStatusRequest {
    aipId: string;
}

export interface StorageRejectRequest {
    aipId: string;
}

export interface StorageSubmitRequest {
    aipId: string;
    submitRequestBody: StorageSubmitRequestBody;
}

export interface StorageUpdateRequest {
    aipId: string;
}

/**
 * StorageApi - interface
 * 
 * @export
 * @interface StorageApiInterface
 */
export interface StorageApiInterface {
    /**
     * Download package by AIPID
     * @summary download storage
     * @param {string} aipId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageDownloadRaw(requestParameters: StorageDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<string>>;

    /**
     * Download package by AIPID
     * download storage
     */
    storageDownload(requestParameters: StorageDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<string>;

    /**
     * List locations
     * @summary list storage
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageListRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Array<StoredLocationResponse>>>;

    /**
     * List locations
     * list storage
     */
    storageList(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Array<StoredLocationResponse>>;

    /**
     * Move a package to a permanent storage location
     * @summary move storage
     * @param {string} aipId 
     * @param {StorageMoveRequestBody} moveRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageMoveRaw(requestParameters: StorageMoveRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Move a package to a permanent storage location
     * move storage
     */
    storageMove(requestParameters: StorageMoveRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Retrieve the status of a permanent storage location move of the package
     * @summary move_status storage
     * @param {string} aipId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageMoveStatusRaw(requestParameters: StorageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<StorageMoveStatusResponseBody>>;

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status storage
     */
    storageMoveStatus(requestParameters: StorageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<StorageMoveStatusResponseBody>;

    /**
     * Reject a package
     * @summary reject storage
     * @param {string} aipId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageRejectRaw(requestParameters: StorageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Reject a package
     * reject storage
     */
    storageReject(requestParameters: StorageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

    /**
     * Start the submission of a package
     * @summary submit storage
     * @param {string} aipId 
     * @param {StorageSubmitRequestBody} submitRequestBody 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageSubmitRaw(requestParameters: StorageSubmitRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<StorageSubmitResponseBody>>;

    /**
     * Start the submission of a package
     * submit storage
     */
    storageSubmit(requestParameters: StorageSubmitRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<StorageSubmitResponseBody>;

    /**
     * Signal the storage service that an upload is complete
     * @summary update storage
     * @param {string} aipId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof StorageApiInterface
     */
    storageUpdateRaw(requestParameters: StorageUpdateRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>>;

    /**
     * Signal the storage service that an upload is complete
     * update storage
     */
    storageUpdate(requestParameters: StorageUpdateRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void>;

}

/**
 * 
 */
export class StorageApi extends runtime.BaseAPI implements StorageApiInterface {

    /**
     * Download package by AIPID
     * download storage
     */
    async storageDownloadRaw(requestParameters: StorageDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<string>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageDownload.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/storage/{aip_id}/download`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.TextApiResponse(response) as any;
    }

    /**
     * Download package by AIPID
     * download storage
     */
    async storageDownload(requestParameters: StorageDownloadRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<string> {
        const response = await this.storageDownloadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List locations
     * list storage
     */
    async storageListRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Array<StoredLocationResponse>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/storage/location`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(StoredLocationResponseFromJSON));
    }

    /**
     * List locations
     * list storage
     */
    async storageList(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Array<StoredLocationResponse>> {
        const response = await this.storageListRaw(initOverrides);
        return await response.value();
    }

    /**
     * Move a package to a permanent storage location
     * move storage
     */
    async storageMoveRaw(requestParameters: StorageMoveRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageMove.');
        }

        if (requestParameters.moveRequestBody === null || requestParameters.moveRequestBody === undefined) {
            throw new runtime.RequiredError('moveRequestBody','Required parameter requestParameters.moveRequestBody was null or undefined when calling storageMove.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/storage/{aip_id}/store`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: StorageMoveRequestBodyToJSON(requestParameters.moveRequestBody),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Move a package to a permanent storage location
     * move storage
     */
    async storageMove(requestParameters: StorageMoveRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.storageMoveRaw(requestParameters, initOverrides);
    }

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status storage
     */
    async storageMoveStatusRaw(requestParameters: StorageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<StorageMoveStatusResponseBody>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageMoveStatus.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/storage/{aip_id}/store`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StorageMoveStatusResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve the status of a permanent storage location move of the package
     * move_status storage
     */
    async storageMoveStatus(requestParameters: StorageMoveStatusRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<StorageMoveStatusResponseBody> {
        const response = await this.storageMoveStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Reject a package
     * reject storage
     */
    async storageRejectRaw(requestParameters: StorageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageReject.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/storage/{aip_id}/reject`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Reject a package
     * reject storage
     */
    async storageReject(requestParameters: StorageRejectRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.storageRejectRaw(requestParameters, initOverrides);
    }

    /**
     * Start the submission of a package
     * submit storage
     */
    async storageSubmitRaw(requestParameters: StorageSubmitRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<StorageSubmitResponseBody>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageSubmit.');
        }

        if (requestParameters.submitRequestBody === null || requestParameters.submitRequestBody === undefined) {
            throw new runtime.RequiredError('submitRequestBody','Required parameter requestParameters.submitRequestBody was null or undefined when calling storageSubmit.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/storage/{aip_id}/submit`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: StorageSubmitRequestBodyToJSON(requestParameters.submitRequestBody),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StorageSubmitResponseBodyFromJSON(jsonValue));
    }

    /**
     * Start the submission of a package
     * submit storage
     */
    async storageSubmit(requestParameters: StorageSubmitRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<StorageSubmitResponseBody> {
        const response = await this.storageSubmitRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Signal the storage service that an upload is complete
     * update storage
     */
    async storageUpdateRaw(requestParameters: StorageUpdateRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.aipId === null || requestParameters.aipId === undefined) {
            throw new runtime.RequiredError('aipId','Required parameter requestParameters.aipId was null or undefined when calling storageUpdate.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/storage/{aip_id}/update`.replace(`{${"aip_id"}}`, encodeURIComponent(String(requestParameters.aipId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Signal the storage service that an upload is complete
     * update storage
     */
    async storageUpdate(requestParameters: StorageUpdateRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<void> {
        await this.storageUpdateRaw(requestParameters, initOverrides);
    }

}
