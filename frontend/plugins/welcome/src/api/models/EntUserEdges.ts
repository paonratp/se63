/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDoctorinfo,
    EntDoctorinfoFromJSON,
    EntDoctorinfoFromJSONTyped,
    EntDoctorinfoToJSON,
    EntFinancier,
    EntFinancierFromJSON,
    EntFinancierFromJSONTyped,
    EntFinancierToJSON,
    EntMedicalrecordstaff,
    EntMedicalrecordstaffFromJSON,
    EntMedicalrecordstaffFromJSONTyped,
    EntMedicalrecordstaffToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntPatientrights,
    EntPatientrightsFromJSON,
    EntPatientrightsFromJSONTyped,
    EntPatientrightsToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * 
     * @type {EntFinancier}
     * @memberof EntUserEdges
     */
    financier?: EntFinancier;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntUserEdges
     */
    historytaking?: EntNurse;
    /**
     * 
     * @type {EntMedicalrecordstaff}
     * @memberof EntUserEdges
     */
    medicalrecordstaff?: EntMedicalrecordstaff;
    /**
     * 
     * @type {EntDoctorinfo}
     * @memberof EntUserEdges
     */
    user2doctorinfo?: EntDoctorinfo;
    /**
     * 
     * @type {EntPatientrights}
     * @memberof EntUserEdges
     */
    userPatientrights?: EntPatientrights;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'financier': !exists(json, 'financier') ? undefined : EntFinancierFromJSON(json['financier']),
        'historytaking': !exists(json, 'historytaking') ? undefined : EntNurseFromJSON(json['historytaking']),
        'medicalrecordstaff': !exists(json, 'medicalrecordstaff') ? undefined : EntMedicalrecordstaffFromJSON(json['medicalrecordstaff']),
        'user2doctorinfo': !exists(json, 'user2doctorinfo') ? undefined : EntDoctorinfoFromJSON(json['user2doctorinfo']),
        'userPatientrights': !exists(json, 'userPatientrights') ? undefined : EntPatientrightsFromJSON(json['userPatientrights']),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'financier': EntFinancierToJSON(value.financier),
        'historytaking': EntNurseToJSON(value.historytaking),
        'medicalrecordstaff': EntMedicalrecordstaffToJSON(value.medicalrecordstaff),
        'user2doctorinfo': EntDoctorinfoToJSON(value.user2doctorinfo),
        'userPatientrights': EntPatientrightsToJSON(value.userPatientrights),
    };
}


