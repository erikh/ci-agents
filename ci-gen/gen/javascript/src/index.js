/**
 * uisvc
 * API for the user interface service; the service that is directly responsible for presenting data to users. This service typically runs at the border, and leverages session cookies or authentication tokens that we generate for users. It also is responsible for handling the act of oauth and user creation through its login hooks. uisvc typically talks to the datasvc and other services to accomplish its goal, it does not save anything locally or carry state. 
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from './ApiClient';
import Error from './model/Error';
import ModelSubmission from './model/ModelSubmission';
import Ref from './model/Ref';
import RepoConfig from './model/RepoConfig';
import Repository from './model/Repository';
import Run from './model/Run';
import RunSettings from './model/RunSettings';
import Task from './model/Task';
import TaskSettings from './model/TaskSettings';
import User from './model/User';
import UserError from './model/UserError';
import DefaultApi from './api/DefaultApi';


/**
* API_for_the_user_interface_service_the_service_that_is_directly_responsible_for_presenting_data_to_users_This_service_typically_runs_at_the_border_and_leverages_session_cookies_or_authentication_tokens_that_we_generate_for_users__It_also_is_responsible_for_handling_the_act_of_oauth_and_user_creation_through_its_login_hooks_uisvc_typically_talks_to_the_datasvc_and_other_services_to_accomplish_its_goal_it_does_not_save_anything_locally_or_carry_state_.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var Uisvc = require('index'); // See note below*.
* var xxxSvc = new Uisvc.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new Uisvc.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new Uisvc.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new Uisvc.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 1.0.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The Error model constructor.
     * @property {module:model/Error}
     */
    Error,

    /**
     * The ModelSubmission model constructor.
     * @property {module:model/ModelSubmission}
     */
    ModelSubmission,

    /**
     * The Ref model constructor.
     * @property {module:model/Ref}
     */
    Ref,

    /**
     * The RepoConfig model constructor.
     * @property {module:model/RepoConfig}
     */
    RepoConfig,

    /**
     * The Repository model constructor.
     * @property {module:model/Repository}
     */
    Repository,

    /**
     * The Run model constructor.
     * @property {module:model/Run}
     */
    Run,

    /**
     * The RunSettings model constructor.
     * @property {module:model/RunSettings}
     */
    RunSettings,

    /**
     * The Task model constructor.
     * @property {module:model/Task}
     */
    Task,

    /**
     * The TaskSettings model constructor.
     * @property {module:model/TaskSettings}
     */
    TaskSettings,

    /**
     * The User model constructor.
     * @property {module:model/User}
     */
    User,

    /**
     * The UserError model constructor.
     * @property {module:model/UserError}
     */
    UserError,

    /**
    * The DefaultApi service constructor.
    * @property {module:api/DefaultApi}
    */
    DefaultApi
};
