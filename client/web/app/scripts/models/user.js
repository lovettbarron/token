/*global Web, Backbone*/

Web.Models = Web.Models || {};

(function () {
    'use strict';

    Web.Models.User = Backbone.Model.extend({

        url: '',

        initialize: function() {

        },

        defaults: {

        },

        validate: function(attrs, options) {

        },

        parse: function(response, options)  {
            return response;
        },

        Authenticated: function() {
            // Placeholder
            return true
        },

        save: function(hash) {
            $.cookie('user_id', auth_hash.id)
            $.cookie('session_id', session.id)
        }


    });

})();
