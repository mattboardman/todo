const url = "http://localhost:8080/v1/todo";

const vm = new Vue({
    el: '#app',
    data: {
        results: [],
            form: {
            title: '',
            description: '',
            started: '',
            completed: '',
            iscompleted: ''
        }
    },
    methods: {
        remove: function (index) {
            axios.delete("http://localhost:8080/v1/todo/"+index);
        },
        create: function(title, description, started, completed, iscompleted) {
            axios.post("http://localhost:8080/v1/todo?title="+title+"?description="+description
                            //    title:title, 
                              //  description:description
                                //startedon:started,
                                //completedon:completed, 
                                //iscompleted:iscompleted
                          //  }
                        );
        }
    },
    mounted() {
        axios.get(url).then(response => {
            this.results = response.data
        })
    }
});