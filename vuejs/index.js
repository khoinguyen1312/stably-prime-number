var app = new Vue({
    el: '#app',
    data: {
        title: 'Finding Prime Number',
        user_input: '',
        output: '',
        clicked: false
    },
    methods: {
        findLowerPrimeNumber: function() {
            this.output = this.user_input
            this.clicked = true
        }
    },
});