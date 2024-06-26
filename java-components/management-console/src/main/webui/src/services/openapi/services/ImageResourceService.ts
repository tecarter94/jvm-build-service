/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ImageDTO } from '../models/ImageDTO';
import type { PageParametersImageDTO } from '../models/PageParametersImageDTO';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class ImageResourceService {
    /**
     * @param requestBody
     * @returns void
     * @throws ApiError
     */
    public static putApiImage(
        requestBody?: string,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/image',
            body: requestBody,
            mediaType: 'text/plain',
        });
    }
    /**
     * @param image
     * @returns ImageDTO OK
     * @throws ApiError
     */
    public static getApiImageScan(
        image: string,
    ): CancelablePromise<ImageDTO> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/image/scan/{image}',
            path: {
                'image': image,
            },
        });
    }
    /**
     * @param repository
     * @param page
     * @param perPage
     * @returns PageParametersImageDTO OK
     * @throws ApiError
     */
    public static getApiImage(
        repository: string,
        page?: number,
        perPage?: number,
    ): CancelablePromise<PageParametersImageDTO> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/image/{repository}',
            path: {
                'repository': repository,
            },
            query: {
                'page': page,
                'perPage': perPage,
            },
        });
    }
}
