/*global Web, Backbone*/

Web.Models = Web.Models || {};

(function () {
    'use strict';

    Web.Models.Token = Backbone.Model.extend({

        url: '',

        initialize: function() {
        },

        defaults: {
            id: '',
            userid: '',
            title: '',
            quantity: '',
            interval: '',
            start: '',
            tokens: {
                id: '',
                tokenid: '',
                timestamp: '',
                value: ''
            }
        },

        validate: function(attrs, options) {
        },

        parse: function(response, options)  {
            return response;
        }
    });

})();
