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
            iscompleted: '',
        }
    },
    methods: {
        remove: function (index) {
            axios.delete("http://localhost:8080/v1/todo/"+index);
        },
        create: function(title, description) {
            axios.post("http://localhost:8080/v1/todo?title="+title+"?description="+description                        );
        },
        update: function(index ,title, description, iscompleted) {
            var flag = false;
            if (iscompleted === "true") {
                flag = true;
            }
            axios.put("http://localhost:8080/v1/todo", {id: index, title: title, description: description, iscompleted: flag});
        }
    },
    mounted() {
        axios.get(url).then(response => {
            this.results = response.data
        })
    }
});