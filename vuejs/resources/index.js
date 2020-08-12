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
            fetch('api/findLowerPrimeNumber/' + this.user_input, {
                "method": "GET"
            })
            .then(response => { 
                if(response.ok){
                    return response.json()    
                } else{
                    alert("Server returned " + response.status + " : " + response.statusText);
                }                
            })
            .then(response => {
                this.output = response['result']
                this.clicked = true
            })
            .catch(err => {
                console.log(err);
            });

            // axios
            //   .get('api/findLowerPrimeNumber/' + this.user_input)
            //   .then(response => (this.output = response['result']))
            // this.clicked = true
        }
    },
});