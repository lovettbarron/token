/*global Web, $*/


window.Web = {
    Models: {},
    Collections: {},
    Views: {},
    Routers: {},
    Global: {
        path: 'localhost:8000'
    },
    init: function () {
        'use strict';
        console.log('Hello from Backbone!');
    }
};

$(document).ready(function () {
    'use strict';
    Web.init();
});
