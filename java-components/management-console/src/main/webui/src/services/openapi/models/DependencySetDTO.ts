/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IdentifiedDependencyDTO } from './IdentifiedDependencyDTO';
export type DependencySetDTO = {
    id?: number;
    name?: string;
    dependencies?: Array<IdentifiedDependencyDTO>;
    totalDependencies: number;
    untrustedDependencies: number;
    trustedDependencies: number;
    availableBuilds: number;
};

