/*global Web, Backbone, JST*/

Web.Views = Web.Views || {};

(function () {
    'use strict';

    Web.Views.User = Backbone.View.extend({

        template: JST['app/scripts/templates/user.ejs'],

        tagName: 'div',

        id: '',

        className: '',

        events: {},

        initialize: function () {
            this.listenTo(this.model, 'change', this.render);
        },

        render: function () {
            this.$el.html(this.template(this.model.toJSON()));
        }

    });

})();
