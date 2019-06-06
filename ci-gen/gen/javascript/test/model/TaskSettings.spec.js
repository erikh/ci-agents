/**
 * uisvc
 * API for the user interface service; the service that is directly responsible for presenting data to users. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.0-SNAPSHOT
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Uisvc);
  }
}(this, function(expect, Uisvc) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Uisvc.TaskSettings();
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

  describe('TaskSettings', function() {
    it('should create an instance of TaskSettings', function() {
      // uncomment below and update the code to test TaskSettings
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be.a(Uisvc.TaskSettings);
    });

    it('should have the property mountpoint (base name: "mountpoint")', function() {
      // uncomment below and update the code to test the property mountpoint
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property workdir (base name: "workdir")', function() {
      // uncomment below and update the code to test the property workdir
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property runs (base name: "runs")', function() {
      // uncomment below and update the code to test the property runs
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property defaultTimeout (base name: "default_timeout")', function() {
      // uncomment below and update the code to test the property defaultTimeout
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property defaultQueue (base name: "default_queue")', function() {
      // uncomment below and update the code to test the property defaultQueue
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property defaultImage (base name: "default_image")', function() {
      // uncomment below and update the code to test the property defaultImage
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property metadata (base name: "metadata")', function() {
      // uncomment below and update the code to test the property metadata
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

    it('should have the property env (base name: "env")', function() {
      // uncomment below and update the code to test the property env
      //var instance = new Uisvc.TaskSettings();
      //expect(instance).to.be();
    });

  });

}));
