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
    EntPatientrecordEdges,
    EntPatientrecordEdgesFromJSON,
    EntPatientrecordEdgesFromJSONTyped,
    EntPatientrecordEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientrecord
 */
export interface EntPatientrecord {
    /**
     * Age holds the value of the "Age" field.
     * @type {number}
     * @memberof EntPatientrecord
     */
    age?: number;
    /**
     * Allergic holds the value of the "Allergic" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    allergic?: string;
    /**
     * Birthday holds the value of the "Birthday" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    birthday?: string;
    /**
     * Bloodtype holds the value of the "Bloodtype" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    bloodtype?: string;
    /**
     * Date holds the value of the "Date" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    date?: string;
    /**
     * Disease holds the value of the "Disease" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    disease?: string;
    /**
     * Email holds the value of the "Email" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    email?: string;
    /**
     * Home holds the value of the "Home" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    home?: string;
    /**
     * Idcardnumber holds the value of the "Idcardnumber" field.
     * @type {number}
     * @memberof EntPatientrecord
     */
    idcardnumber?: number;
    /**
     * Name holds the value of the "Name" field.
     * @type {string}
     * @memberof EntPatientrecord
     */
    name?: string;
    /**
     * Phonenumber holds the value of the "Phonenumber" field.
     * @type {number}
     * @memberof EntPatientrecord
     */
    phonenumber?: number;
    /**
     * 
     * @type {EntPatientrecordEdges}
     * @memberof EntPatientrecord
     */
    edges?: EntPatientrecordEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatientrecord
     */
    id?: number;
}

export function EntPatientrecordFromJSON(json: any): EntPatientrecord {
    return EntPatientrecordFromJSONTyped(json, false);
}

export function EntPatientrecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientrecord {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'Age') ? undefined : json['Age'],
        'allergic': !exists(json, 'Allergic') ? undefined : json['Allergic'],
        'birthday': !exists(json, 'Birthday') ? undefined : json['Birthday'],
        'bloodtype': !exists(json, 'Bloodtype') ? undefined : json['Bloodtype'],
        'date': !exists(json, 'Date') ? undefined : json['Date'],
        'disease': !exists(json, 'Disease') ? undefined : json['Disease'],
        'email': !exists(json, 'Email') ? undefined : json['Email'],
        'home': !exists(json, 'Home') ? undefined : json['Home'],
        'idcardnumber': !exists(json, 'Idcardnumber') ? undefined : json['Idcardnumber'],
        'name': !exists(json, 'Name') ? undefined : json['Name'],
        'phonenumber': !exists(json, 'Phonenumber') ? undefined : json['Phonenumber'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientrecordEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPatientrecordToJSON(value?: EntPatientrecord | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Age': value.age,
        'Allergic': value.allergic,
        'Birthday': value.birthday,
        'Bloodtype': value.bloodtype,
        'Date': value.date,
        'Disease': value.disease,
        'Email': value.email,
        'Home': value.home,
        'Idcardnumber': value.idcardnumber,
        'Name': value.name,
        'Phonenumber': value.phonenumber,
        'edges': EntPatientrecordEdgesToJSON(value.edges),
        'id': value.id,
    };
}

