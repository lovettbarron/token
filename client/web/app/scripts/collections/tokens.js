/*global Web, Backbone*/

Web.Collections = Web.Collections || {};

(function () {
    'use strict';

    Web.Collections.Tokens = Backbone.Collection.extend({

        model: Web.Models.Tokens

    });

})();
