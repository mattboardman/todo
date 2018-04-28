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
            window.location.reload()
        },
        create: function(title, description) {
            axios.post("http://localhost:8080/v1/todo?title="+title+"?description="+description);
            window.location.reload()
        },
        update: function(index ,title, description, iscompleted) {
            var flag = false;
            if (iscompleted === "true") {
                flag = true;
            }
            axios.put("http://localhost:8080/v1/todo", {id: index, title: title, description: description, iscompleted: flag});
            window.location.reload()
        },        
        search: function(query) {
            axios.get("http://localhost:8080/v1/todo/search/"+query).then(response => {
                document.getElementById('title').innerHTML = response.data.title;
                document.getElementById('description').innerHTML = response.data.description;
                document.getElementById('started').innerHTML = response.data.started;
                document.getElementById('completed').innerHTML = response.data.completed;
                document.getElementById('iscompleted').innerHTML = response.data.iscompleted;
                document.getElementById('search-results').style.visibility = "visible";
            });
        },
    },
    mounted() {
        axios.get(url).then(response => {
            this.results = response.data
        })
    }
});
