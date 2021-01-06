/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
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
    EntPlaylistVideo,
    EntPlaylistVideoFromJSON,
    EntPlaylistVideoFromJSONTyped,
    EntPlaylistVideoToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntVideoEdges
 */
export interface EntVideoEdges {
    /**
     * 
     * @type {EntUser}
     * @memberof EntVideoEdges
     */
    owner?: EntUser;
    /**
     * PlaylistVideos holds the value of the playlist_videos edge.
     * @type {Array<EntPlaylistVideo>}
     * @memberof EntVideoEdges
     */
    playlistVideos?: Array<EntPlaylistVideo>;
}

export function EntVideoEdgesFromJSON(json: any): EntVideoEdges {
    return EntVideoEdgesFromJSONTyped(json, false);
}

export function EntVideoEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntVideoEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'owner': !exists(json, 'owner') ? undefined : EntUserFromJSON(json['owner']),
        'playlistVideos': !exists(json, 'playlistVideos') ? undefined : ((json['playlistVideos'] as Array<any>).map(EntPlaylistVideoFromJSON)),
    };
}

export function EntVideoEdgesToJSON(value?: EntVideoEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'owner': EntUserToJSON(value.owner),
        'playlistVideos': value.playlistVideos === undefined ? undefined : ((value.playlistVideos as Array<any>).map(EntPlaylistVideoToJSON)),
    };
}

