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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Uisvc);
  }
}(this, function(expect, Uisvc) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Uisvc.RunSettings();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('RunSettings', function() {
    it('should create an instance of RunSettings', function() {
      // uncomment below and update the code to test RunSettings
      //var instane = new Uisvc.RunSettings();
      //expect(instance).to.be.a(Uisvc.RunSettings);
    });

    it('should have the property command (base name: "command")', function() {
      // uncomment below and update the code to test the property command
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property image (base name: "image")', function() {
      // uncomment below and update the code to test the property image
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property queue (base name: "queue")', function() {
      // uncomment below and update the code to test the property queue
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property metadata (base name: "metadata")', function() {
      // uncomment below and update the code to test the property metadata
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property timeout (base name: "timeout")', function() {
      // uncomment below and update the code to test the property timeout
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

    it('should have the property resources (base name: "resources")', function() {
      // uncomment below and update the code to test the property resources
      //var instance = new Uisvc.RunSettings();
      //expect(instance).to.be();
    });

  });

}));
