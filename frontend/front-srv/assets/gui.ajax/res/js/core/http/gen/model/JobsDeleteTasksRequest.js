/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */


import ApiClient from '../ApiClient';
import JobsTaskStatus from './JobsTaskStatus';





/**
* The JobsDeleteTasksRequest model module.
* @module model/JobsDeleteTasksRequest
* @version 1.0
*/
export default class JobsDeleteTasksRequest {
    /**
    * Constructs a new <code>JobsDeleteTasksRequest</code>.
    * @alias module:model/JobsDeleteTasksRequest
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>JobsDeleteTasksRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/JobsDeleteTasksRequest} obj Optional instance to populate.
    * @return {module:model/JobsDeleteTasksRequest} The populated <code>JobsDeleteTasksRequest</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobsDeleteTasksRequest();

            
            
            

            if (data.hasOwnProperty('JobId')) {
                obj['JobId'] = ApiClient.convertToType(data['JobId'], 'String');
            }
            if (data.hasOwnProperty('TaskID')) {
                obj['TaskID'] = ApiClient.convertToType(data['TaskID'], ['String']);
            }
            if (data.hasOwnProperty('Status')) {
                obj['Status'] = ApiClient.convertToType(data['Status'], [JobsTaskStatus]);
            }
            if (data.hasOwnProperty('PruneLimit')) {
                obj['PruneLimit'] = ApiClient.convertToType(data['PruneLimit'], 'Number');
            }
        }
        return obj;
    }

    /**
    * @member {String} JobId
    */
    JobId = undefined;
    /**
    * @member {Array.<String>} TaskID
    */
    TaskID = undefined;
    /**
    * @member {Array.<module:model/JobsTaskStatus>} Status
    */
    Status = undefined;
    /**
    * @member {Number} PruneLimit
    */
    PruneLimit = undefined;








}


