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
        reverseIndex: function() {
            axios.get("http://localhost:8080/v1/todo/reverse").then(response => {
                this.results = response.data
            });
        },
        completed: function() {
            axios.get("http://localhost:8080/v1/todo/completed").then(response => {
                this.results = response.data
            });
        },
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
            document.getElementById('search-results').style.visibility = "hidden";
            var date = new Date()
            var start = window.performance.now()
            axios.get("http://localhost:8080/v1/todo/search/"+query).then(response => {
                if (response.data != null) {
                    document.getElementById('uuid').innerHTML = 'ID: ' + response.data.id;
                    document.getElementById('title').innerHTML = 'Title: ' + response.data.title;
                    document.getElementById('description').innerHTML = 'Description: ' + response.data.description;
                    document.getElementById('started').innerHTML = 'Started: ' + response.data.started;
                    document.getElementById('completed').innerHTML = 'Completed: ' + response.data.completed;
                    document.getElementById('iscompleted').innerHTML = 'Is Completed: ' + response.data.iscompleted;
                    document.getElementById('search-results').style.visibility = "visible";
                } else {
                    document.getElementById('uuid').innerHTML = "No search Results Found";
                    document.getElementById('title').innerHTML = '';
                    document.getElementById('description').innerHTML = '';
                    document.getElementById('started').innerHTML = '';
                    document.getElementById('completed').innerHTML = '';
                    document.getElementById('iscompleted').innerHTML = '';
                }
            });
            var finish = window.performance.now()
            var factor = Math.pow(10, 4);
            document.getElementById('search-time').innerHTML = 'Search Time: ' + (finish - start) + ' ms';
        },
        searchV2: function(query) {
            document.getElementById('search-results').style.visibility = "hidden";
            var date = new Date()
            var start = window.performance.now()
            axios.get("http://localhost:8080/v2/todo/search/"+query).then(response => {
                if (response.data != null) {
                    document.getElementById('uuid').innerHTML = 'ID: ' + response.data.id;
                    document.getElementById('title').innerHTML = 'Title: ' + response.data.title;
                    document.getElementById('description').innerHTML = 'Description: ' + response.data.description;
                    document.getElementById('started').innerHTML = 'Started: ' + response.data.started;
                    document.getElementById('completed').innerHTML = 'Completed: ' + response.data.completed;
                    document.getElementById('iscompleted').innerHTML = 'Is Completed: ' + response.data.iscompleted;
                    document.getElementById('search-results').style.visibility = "visible";
                } else {
                    document.getElementById('uuid').innerHTML = "No search Results Found";
                    document.getElementById('title').innerHTML = '';
                    document.getElementById('description').innerHTML = '';
                    document.getElementById('started').innerHTML = '';
                    document.getElementById('completed').innerHTML = '';
                    document.getElementById('iscompleted').innerHTML = '';
                }
            });
            var finish = window.performance.now()
            document.getElementById('search-time').innerHTML = 'Search Time: ' + (finish - start) + ' ms';
        },
    },
    mounted() {
        axios.get(url).then(response => {
            this.results = response.data
        })
    }
});
