var app = new Vue({
    el: '#app',
    data: {
        title: 'Finding Prime Number',
        user_input: '',
        output: '',
        clicked: false,
        error_message: '',
        is_show_error_message: false
    },
    methods: {
        findLowerPrimeNumber: function() {
            fetch('api/findLowerPrimeNumber/' + this.user_input, {
                "method": "GET"
            })
            .then(response => { 
                if(response.ok){
                    response.json().then(json => { 
                        this.output = json['result']
                        this.clicked = true
                        this.is_show_error_message = false
                    })
                    .catch(err => {
                        console.log(err);
                    }) 
                } else {
                    this.error_message = response.statusText
                    this.is_show_error_message = true
                }
            })

            // axios
            //   .get('api/findLowerPrimeNumber/' + this.user_input)
            //   .then(response => (this.output = response['result']))
            // this.clicked = true
        }
    },
});